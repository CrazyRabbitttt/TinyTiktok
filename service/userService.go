package service

import (
	"Web-Go/Common"
	"Web-Go/ConnSql"
	"Web-Go/Model"
	"gorm.io/gorm"
)

//限定的一些特定的字段值
const (
	MaxUserNameLength = 32
	MaxPasswdLength   = 32
	MinPasswdLength   = 6
)

func CreateNewUser(userName string, passWord string) (Model.User, error) {
	db := ConnSql.ThemodelOfSql()

	//进行user的创建
	newUser := Model.User{
		Name:     userName,
		Password: passWord,
	}

	if IfUserNotExistByName(userName) {
		result := db.Table("tik_user").Create(&newUser)
		return newUser, result.Error
	}
	return newUser, Common.ErrorCreateUser

}

/*pass the name and passwd, check if it is legal*/
func IsUserLegal(userName string, passWord string) error {

	//check userName
	if userName == "" {
		return Common.ErrorUserNameNull
	}
	if len(userName) > MaxUserNameLength {
		return Common.ErrorUserNameNotValid
	}

	//check passwd
	if passWord == "" {
		return Common.ErrorPasswordNull
	}
	if len(passWord) > MaxPasswdLength || len(passWord) < MinPasswdLength {
		return Common.ErrorPasswordNotValid
	}

	return nil
}

func IfUserNotExistByName(userName string) bool {

	db := ConnSql.ThemodelOfSql() //获得数据库的连接
	var dbUser Model.User
	result := db.Table("tik_user").Where("name = ?", userName).First(&dbUser)

	//检查表中是否是存在记录的
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return true
		}
		return false
	}
	return true
}
