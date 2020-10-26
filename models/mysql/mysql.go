package mysql

import (
	"fiber/helpers"
	"github.com/davecgh/go-spew/spew"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	mysqlDb0 := helpers.LoadConfiguration("mysql_db_0")
	spew.Dump(mysqlDb0)
	db, err := gorm.Open(mysql.Open(mysqlDb0.(string)), &gorm.Config{})
	if err != nil {
		panic(err)
		//logs.Error(err)
	}
	DB = db
}