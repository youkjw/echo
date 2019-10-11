package orm

import (
	. "echo/conf"
	"echo/module/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

		//newDb.SetLogger(gorm.Logger{})
		newDb.LogMode(true)
		db = newDb
	}

	return db
}

func connect() (*gorm.DB, error){
	var err error
	connection := Conf.DB.UserName + ":" + Conf.DB.Pwd + "@(" + Conf.DB.Host + ":" + Conf.DB.Port + ")/" + Conf.DB.DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection)

	if err != nil {
		return nil, err
	}
	return db, nil
}