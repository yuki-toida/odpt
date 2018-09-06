package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/refodpt/backend/config"
)

// Connect func
func Connect() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Config.DB.User, config.Config.DB.Pass, config.Config.DB.Host, config.Config.DB.Port, config.Config.DB.Name)
	db, err := gorm.Open("mysql", conn+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}

	db.DB().SetMaxIdleConns(config.Config.DB.Pool)
	return db
}
