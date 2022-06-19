package service

import (
	"Web-Go/Common"
	"Web-Go/ConnSql"
	"Web-Go/Model"
	"errors"
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

	result := db.Table("tik_user").Where("name = ?", userName).First(&newUser)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) { //如果没有的话，正好进行插入
			db.Table("tik_user").Create(&newUser)
		}
	} else {
		//没错就是用户已经是存在的了
		return newUser, Common.ErrorUserExits
	}

	db.Table("tik_user").Create(&newUser)
	db.Table("tik_user").Where("name = ?", userName).First(&newUser)

	return newUser, nil

}

/*pass the name and passwd, check if it is legal*/
func IsUserLegal(userName string, passWord string) error {

	//check userName
	println("传入的用户名参数:", userName, passWord)

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
