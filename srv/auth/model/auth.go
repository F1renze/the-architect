package model

import (
	"database/sql"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/common/infra/db"
	"github.com/f1renze/the-architect/common/utils"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"time"

	pb "github.com/f1renze/the-architect/srv/auth/proto"
)

type AuthModel interface {
	QueryCredential(authType pb.AuthType, authId, credential string) (*pb.AuthInfo, error)
	CreateCredential(authInfo *pb.AuthInfo) (int64, error)
	RefreshLoginTime(id uint32, ip string) error
}

func NewModel() AuthModel {
	return &model{
		db: db.GetDB(),
	}
}

type model struct {
	db *sqlx.DB
}

func (m *model) RefreshLoginTime(id uint32, ip string) error {
	query := "UPDATE `user_auth` SET latest_login_at = ?, ip_addr = INET_ATON(?)"
	_, err := m.db.Exec(query, time.Now(), ip)
	if err != nil {
		log.ErrorF("auth model: refresh login time failed, %s", err)
		return errno.DBErr.With(err)
	}
	return nil
}

func (m *model) QueryCredential(authType pb.AuthType, authId, credential string) (*pb.AuthInfo, error) {
	authInfo, err := m.GetAuthInfoByAuthId(authType, authId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errno.RecordNotExists
		}
		log.ErrorF("auth model: query credential failed, %s", err)
		return nil, errno.DBErr.With(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(authInfo.Credential), []byte(credential))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errno.PwdNotCorrect.With(err)
		}
		log.ErrorF("auth model: compare password error, %s", err)
		return nil, errno.SystemErr.With(err)
	}

	return authInfo, nil
}

func (m *model) CreateCredential(authInfo *pb.AuthInfo) (int64, error) {
	uid := authInfo.GetUid()
	authId := authInfo.GetAuthId()
	credential := authInfo.GetCredential()
	authType := authInfo.GetAuthType()

	if uid == 0 || authId == "" || credential == "" || !ValidateAuthId(authType, authId) {
		return -1, errno.InvalidAuthInfo
	}

	query := "INSERT INTO `user_auth` (uid, auth_type, auth_id, credential) VALUES (?, ?, ?, ?)"

	encrypted, err := bcrypt.GenerateFromPassword([]byte(credential), 8)
	if err != nil {
		return -1, errno.GeneratePwdErr.With(err)
	}
	r, err := m.db.Exec(query, uid, authType, authId, encrypted)
	if err != nil {
		if utils.IskMySQLError(err, errno.MySQLDupEntryErrNo) {
			ok, err := m.CheckAuthTypeAlreadyBind(uid, authType)
			if ok && err == nil {
				return -1, errno.AuthTypeAlreadyBind
			} else if err == nil {
				return -1, errno.AuthIdAlreadyUsed
			}
		}
		log.ErrorF("auth model: create credential failed, %s", err)
		return -1, errno.DBErr.With(err)
	}

	id, _ := r.LastInsertId()
	return id, nil
}

func (m *model) GetAuthInfoByAuthId(authType pb.AuthType, authId string) (*pb.AuthInfo, error){
	if !ValidateAuthId(authType, authId) {
		return nil, errno.InvalidAuthInfo
	}
	// get ip addr sql: select id, INET_NTOA(ip_addr) as ip_addr from `user_auth` where id = 1;
	query := "SELECT id, uid, verified, credential FROM `user_auth` WHERE auth_type = ? AND auth_id = ?"

	authInfo := &pb.AuthInfo{}
	err := m.db.Get(authInfo, query, authType, authId)
	return authInfo, err
}

func (m *model) CheckAuthTypeAlreadyBind(uid uint32, authType pb.AuthType) (bool, error) {
	query := "SELECT auth_id, verified FROM `user_auth` WHERE auth_type = ? AND uid = ?"

	_, err := m.db.Query(query, authType, uid)
	return err != sql.ErrNoRows, err
}

func ValidateAuthId(authType pb.AuthType, authId string) bool {
	switch authType {
	case pb.AuthType_Mobile:
		return utils.ValidateMobile(authId)
	case pb.AuthType_Email:
		return utils.ValidateEmailFormat(authId)
	case pb.AuthType_Github:
		return true
	default:
	}
	return false
}
