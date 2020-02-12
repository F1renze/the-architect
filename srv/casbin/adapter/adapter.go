package adapter

import (
	"github.com/casbin/casbin/v2/persist"
	"github.com/casbin/gorm-adapter/v2"
	//_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/f1renze/the-architect/common/errno"
	pb "github.com/f1renze/the-architect/srv/casbin/proto"
)

var supportDriverNames = [...]string{"file", "mysql", "postgres", "mssql"}

// todo append sqlx adapter
func NewAdapter(in *pb.NewAdapterRequest) (persist.Adapter, error) {

	var support = false
	for _, driverName := range supportDriverNames {
		if driverName == in.DriverName {
			support = true
			break
		}
	}
	if !support {
		return nil, errno.UnSupportAdapterDriver
	}

	a, err := gormadapter.NewAdapter(in.DriverName, in.ConnectString, in.DbSpecified)
	if err != nil {
		return nil, errno.CreateAdapterErr.With(err)
	}

	return a, nil
}
