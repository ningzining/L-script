package transform

import (
	"L-script/utils/dbutils"
	"fmt"
	"github.com/ningzining/L-tool/structutil"
	"gorm.io/gorm"
)

type OldUser struct {
	Id       int64
	Username string
	Email    string
	Age      int64
}

func QueryOldUser(tx *gorm.DB) (list []OldUser) {
	tx.Model(OldUser{}).Limit(10).Find(&list)
	return
}

type NewUser struct {
	Id       int64
	Username string
	Age      int64
	OpenId   string
	Alias    string // 别名
}

func InsertNewUser(tx *gorm.DB, list []NewUser) {
	tx.Model(NewUser{}).Create(&list)
	return
}

func Transform() {
	var (
		oldDsn = ""
		newDsn = ""
	)
	oldMysql := dbutils.Mysql(oldDsn)
	newMysql := dbutils.Mysql(newDsn)

	oldUserList := QueryOldUser(oldMysql)
	var newUserList []NewUser
	for _, oldUser := range oldUserList {
		newUser := structutil.Convert[NewUser](oldUser)
		newUser.Alias = fmt.Sprintf("%s%d", oldUser.Username, oldUser.Id)
		newUserList = append(newUserList, newUser)
	}
	InsertNewUser(newMysql, newUserList)

}
