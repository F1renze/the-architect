package model

import (
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
		return nil, EmptyUserNameErr
	}

	r, err := m.db.Exec(
		"INSERT INTO `user` (name, avatar) VALUES (?, ?)",
		name, utils.CheckNullString(avatar),
	)
	if err != nil {
		return nil, err
	}

	id, _ := r.LastInsertId()

	return m.QueryUser(uint32(id))
}

func (m *model) QueryUser(id uint32) (user *pb.User, err error) {
	query := "SELECT id, name FROM `user` WHERE id = ?"

	user = new(pb.User)

	err = m.db.Get(user, query, id)
	return
}
