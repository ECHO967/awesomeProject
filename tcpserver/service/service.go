package service

import (
	global2 "awesomeProject/config/global"
	"awesomeProject/pkg/errcode"
	"awesomeProject/pkg/jwt"
	"awesomeProject/tcpserver/dao"
	"awesomeProject/tcpserver/model"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global2.DBEngin, global2.ReEngin)
	return svc
}

// 登录方法
func (svc *Service) Login(username string, password string) (*model.User, string, error) {
	user, err := svc.dao.GetUser(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global2.TCPLogger.Errorf("TCPServer: svc.dao.GetUser failed: %v", err)
		return nil, "", errcode.ToRPCError(errcode.ErrorOther)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) || password != user.Password {
		global2.TCPLogger.Error("用户名或密码错误")
		return nil, "", errcode.ToRPCError(errcode.ErrorPassword)
	}
	token, err := jwt.GenerateToken(username, password)
	if err != nil {
		global2.TCPLogger.Errorf("TCPServer: app.GenerateToken err: %v", err)
		return nil, "", errcode.ToRPCError(errcode.ErrorOther)
	}

	return &user, token, nil
}

// 更新nickname方法
func (svc *Service) UpdateNick(nickname string, username string, token string) error {
	_, err := jwt.ParseToken(token)
	if err != nil {
		global2.TCPLogger.Errorf("TCPServer:app.ParseToken err: %v", err)
		return errcode.ToRPCError(errcode.ErrorAuth)
	}
	return svc.dao.UpdateNick(nickname, username)
}

// 更新头像方法
func (svc *Service) UpdateProf(profile string, username string, token string) error {
	_, err := jwt.ParseToken(token)
	if err != nil {
		global2.TCPLogger.Errorf("TCPServer:app.ParseToken err: %v", err)
		return errcode.ToRPCError(errcode.ErrorAuth)
	}
	return svc.dao.UpdateProf(profile, username)
}
