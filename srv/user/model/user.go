package model

import (
	"database/sql"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/jmoiron/sqlx"

	"github.com/f1renze/the-architect/common/infra/db"
	"github.com/f1renze/the-architect/common/utils"
	pb "github.com/f1renze/the-architect/srv/user/proto"
)

type UserModel interface {
	QueryUser(id uint32) (*pb.User, error)
	CreateUser(name, avatar string) (*pb.User, error)
}

func NewModel() UserModel {
	return &model{
		db: db.GetDB(),
	}
}

type model struct {
	db *sqlx.DB
}

func (m *model) CreateUser(name, avatar string) (*pb.User, error) {

	if len(name) < 1 && len(avatar) < 1 {
		return nil, errno.InvalidUserName
	}

	r, err := m.db.Exec(
		"INSERT INTO `user` (name, avatar) VALUES (?, ?)",
		name, utils.CheckNullString(avatar),
	)
	if err != nil {
		if utils.IskMySQLError(err, errno.MySQLDupEntryErrNo) {
			return nil, errno.UserNameAlreadyUsed
		}
		log.ErrorF("user model: create user failed, %s", err)
		return nil, errno.DBErr.With(err)
	}

	id, _ := r.LastInsertId()

	return m.QueryUser(uint32(id))
}

func (m *model) QueryUser(id uint32) (user *pb.User, err error) {
	query := "SELECT id, name FROM `user` WHERE id = ?"

	user = new(pb.User)

	err = m.db.Get(user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errno.RecordNotExists
		}
		log.ErrorF("user model: query user error, %s", err)
		return nil, errno.DBErr.With(err)
	}
	return
}
