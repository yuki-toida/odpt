package db

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/refodpt/backend/config"
)

// Connect func
func Connect() *gorm.DB {
	conn := config.Config.DB.User + ":" + config.Config.DB.Pass + "@tcp(" + config.Config.DB.Host + ":" + config.Config.DB.Port + ")/" + config.Config.DB.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}

	db.DB().SetMaxIdleConns(config.Config.DB.Pool)
	return db
}
