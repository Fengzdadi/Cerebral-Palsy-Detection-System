package model

import (
	"Cerebral-Palsy-Detection-System/Pkg/e"
	"Cerebral-Palsy-Detection-System/Serializer"
	"fmt"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Email          string //`gorm:"unique"`
	Avatar         string `gorm:"size:1000"`
	Phone          string
	Status         string
}

const (
	PassWordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarUrl 封面地址
func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	// NickName string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=10"`
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
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: "已经存在该用户了",
		}
	}
	user = User{
		UserName: service.UserName,
		Status:   Active,
	}
	// 加密密码
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
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: "创建失败",
		}
	}
	return Serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
	}
}

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16"`
}

func (service UserLoginService) Login() Serializer.Response {
	var user User
	var count int
	code := e.SUCCESS
	DB.Model(&User{}).Where("user_name=?", service.UserName).Count(&count)
	if count == 0 {
		code = e.ERROR
		return Serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: "用户不存在",
		}
	}
	DB.Where("user_name=?", service.UserName).First(&user)
	if !user.CheckPassword(service.Password) {
		code = e.ERROR
		return Serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: "密码错误",
		}
	}

	return Serializer.Response{
		Code: code,
		Msg:  e.GetMsg(code),
	}
}

type UserUpdatePwdService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16"`
	NewPwd   string `form:"new_pwd" json:"new_pwd" binding:"required,min=8,max=16"`
}

func (service UserUpdatePwdService) Update() Serializer.Response {
	var user User
	user.SetPassword(service.NewPwd)
	DB.Model(&User{}).Where("user_name=?", service.UserName).Update("password_digest", user.PasswordDigest)
	return Serializer.Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
	}
}

func GetUserid(username string) uint {
	var user User
	fmt.Println("username:", username)
	DB.Model(&User{}).Where("user_name=?", username).First(&user)
	fmt.Println("userid:", user.ID)
	return user.ID
}
