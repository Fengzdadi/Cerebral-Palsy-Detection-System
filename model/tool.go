package model

import (
	"Cerebral-Palsy-Detection-System/WS/Pkg/e"
	"Cerebral-Palsy-Detection-System/WS/Serializer"
	logging "github.com/sirupsen/logrus"
)

type UserRegisterService struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=10"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16"`
}

func (service UserRegisterService) Register() Serializer.Response {
	var user User
	var count int
	code := e.SUCCESS
	DB.Model(&User{}).Where("user_name=?", service.UserName).Count(&count)
	if count == 1 {
		code = e.ERROR
		return Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: "已经存在该用户了",
		}
	}
	user = User{
		UserName: service.UserName,
		Status:   Active,
	}
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ERROR
		return Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	user.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	//创建用户
	if err := DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
	return Serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
	}

}
