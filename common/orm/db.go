package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func InitDB() {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s",
		"developer", "developer", "127.0.0.1", 3306, "arch_dev", true, "Local",
	)

	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Fatal("failed to connect db", err)
	}

	db.SingularTable(true)
	db.LogMode(true)

	db.DropTableIfExists(&User{}, &UserAuth{}, &Identification{})

	db.AutoMigrate(&User{}, &UserAuth{}, &Identification{})

	identifications := []Identification{
		{Name: "Mobile"},
		{Name: "Email"},
		{Name: "Github"},
	}

	for i := range identifications {
		db.Create(&identifications[i])
	}
}
