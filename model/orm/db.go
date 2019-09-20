package orm

import (
	"echo/module/log"
	"github.com/hb-go/echo-web/model/orm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "echo/conf"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {

		log.Debugf("Model NewDB")

		newDb, err := connect()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		newDb.SetLogger(orm.Logger{})
		newDb.LogMode(true)
		db = newDb
	}

	return db
}

func connect() (*gorm.DB, error){
	conntction := Conf.DB.UserName + ":" + Conf.DB.Pwd + "@/(" + Conf.DB.Host + ")" + Conf.DB.DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", conntction)

	if err != nil {
		return nil, err
	}
	return db, nil
}