package dbutils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("gorm link error: %s", err.Error())
		return nil
	}
	return db
}
