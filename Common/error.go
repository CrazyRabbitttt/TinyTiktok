package Common

import "errors"

var (
	ErrorUserNameNull     = errors.New("用户名为空")
	ErrorUserNameNotValid = errors.New("用户名不符合规范")
	ErrorPasswordNull     = errors.New("密码为空")
	ErrorPasswordNotValid = errors.New("密码不符合规范")
	ErrorCreateUser       = errors.New("创建新用户失败")
	ErrorUserExits        = errors.New("用户已经是存在的")
)
