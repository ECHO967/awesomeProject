package grpc

import (
	global2 "awesomeProject/config/global"
	"awesomeProject/httpserver/upload"
	"awesomeProject/pkg/common"
	"awesomeProject/pkg/proto"
	"context"
	"errors"
	"google.golang.org/grpc"
	"mime/multipart"
	"os"
)

type Service struct {
	ctx        context.Context
	clientConn *grpc.ClientConn
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.clientConn = global2.ClientConn
	return svc
}

// 定义各个结构体
type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UpdateNicknameRequest struct {
	Username string `form:"username"`
	Nickname string `form:"nickname"`
	Token    string `form:"token"`
}

type UpdateProfRequest struct {
	Username string `form:"username"`
	Profile  string `form:"Profile"`
	Token    string `form:"token"`
}

// RPC客户端调用登录服务
func (svc *Service) Login(para *LoginRequest) (*common.UserInfo, string, error) {
	userServiceClient := user.NewSearchServiceClient(svc.clientConn)

	resp, err := userServiceClient.Login(svc.ctx, &user.LoginRequest{
		Username: para.Username,
		Password: para.Password,
	})
	if err != nil {

		return nil, "", err
	}
	return &common.UserInfo{Username: para.Username, Nickname: resp.Nickname, Profile: resp.Profile}, resp.Token, nil
}
func (svc *Service) UpdateNic(para *common.UpdateInfo) error {
	userServiceClient := user.NewSearchServiceClient(svc.clientConn)

	_, err := userServiceClient.UpdateNickname(svc.ctx, &user.UpdateNicknameRequest{
		Username: para.Username,
		Nickname: para.Nickname,
		Token:    para.Token,
	})
	if err != nil {
		return err
	}
	return nil
}

// RPC客户端调用上传文件服务
func (svc *Service) UploadFile(username string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	name := username + upload.GetFileExt(fileHeader.Filename)
	fileName := upload.GetFileName(name)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	fileInfo := global2.UploadSetting.UploadServerUrl + "/" + fileName
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return "", errors.New("failed to creat save directory.")
		}
	}
	if upload.CheckMaxSize(file) {
		return "", errors.New("exceeded maximum file limit.")
	}
	if upload.CheckPermission(uploadSavePath) {
		return "", errors.New("insuffient file permission.")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return "", err
	}
	return fileInfo, nil
}

// RPC客户端调用更新用户信息服务
func (svc *Service) UpdateProf(para *common.UpdateInfo) error {
	userServiceClient := user.NewSearchServiceClient(svc.clientConn)

	_, err := userServiceClient.UpdateProfile(svc.ctx, &user.UpdateProfileRequest{
		Username: para.Username,
		Profile:  para.Profile,
		Token:    para.Token,
	})

	if err != nil {
		return err
	}

	return nil
}
