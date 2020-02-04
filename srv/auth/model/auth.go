package model

import (
	"github.com/f1renze/the-architect/common/infra/db"
	"github.com/f1renze/the-architect/common/utils"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"

	pb "github.com/f1renze/the-architect/srv/auth/proto"
)

type AuthModel interface {
	QueryCredential(authType pb.AuthType, authId, credential string) (*pb.AuthInfo, error)
	CreateCredential(authInfo *pb.AuthInfo) (int64, error)
}

func NewModel() AuthModel {
	return &model{
		db: db.GetDB(),
	}
}

type model struct {
	db *sqlx.DB
}

func (m *model) QueryCredential(authType pb.AuthType, authId, credential string) (authInfo *pb.AuthInfo, err error) {
	if !ValidateAuthId(authType, authId) {
		return nil, InvalidAuthInfoErr
	}

	query := "SELECT uid, verified, credential FROM `user_auth` WHERE auth_type = ? AND auth_id = ?"

	authInfo = &pb.AuthInfo{}
	err = m.db.Get(authInfo, query, authType, authId)
	err2 := bcrypt.CompareHashAndPassword([]byte(authInfo.Credential), []byte(credential))
	if err = utils.NoErrors(err, err2); err != nil {
		return nil, err
	}

	return
}

func (m *model) CreateCredential(authInfo *pb.AuthInfo) (int64, error) {
	uid := authInfo.GetUid()
	authId := authInfo.GetAuthId()
	credential := authInfo.GetCredential()
	authType := authInfo.GetAuthType()

	if uid == 0 || authId == "" || credential == "" || !ValidateAuthId(authType, authId) {
		return -1, InvalidAuthInfoErr
	}

	query := "INSERT INTO `user_auth` (uid, auth_type, auth_id, credential) VALUES (?, ?, ?, ?)"

	encrypted, err := bcrypt.GenerateFromPassword([]byte(credential), 8)
	if err != nil {
		return -1, err
	}
	r, err := m.db.Exec(query, uid, authType, authId, encrypted)
	if err != nil {
		return -1, err
	}
	id, _ := r.LastInsertId()
	return id, nil
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
