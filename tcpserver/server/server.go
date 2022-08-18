package server

import (
	pb "awesomeProject/pkg/proto"
	"awesomeProject/tcpserver/service"
	"context"
	//user "github.com/ECHO967/entry-task/common/proto"
)

// rpc服务端提供服务
type UserServer struct {
	pb.UnimplementedSearchServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

// RPC服务端 用户登录
func (t *UserServer) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	svc := service.New(ctx)
	user, token, err := svc.Login(r.GetUsername(), r.GetPassword())
	if err != nil {
		return &pb.LoginResponse{Err: 1}, err
	}
	userInfo := pb.LoginResponse{
		Nickname: user.Nickname,
		Profile:  user.Profile,
		Token:    token,
		Err:      0,
	}

	return &userInfo, err
}

// RPC服务端 更新nickname
func (t *UserServer) UpdateNickname(ctx context.Context, r *pb.UpdateNicknameRequest) (*pb.UpdateNicknameResponse, error) {
	svc := service.New(ctx)
	err := svc.UpdateNick(r.GetNickname(), r.GetUsername(), r.GetToken())
	rep := pb.UpdateNicknameResponse{Err: 0}
	if err != nil {
		return &rep, err
	}
	return &rep, err
}

// RPC服务端 更新头像
func (t *UserServer) UpdateProfile(ctx context.Context, r *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	////TODO implement me
	//panic("implement me")
	svc := service.New(ctx)
	err := svc.UpdateProf(r.GetProfile(), r.GetUsername(), r.GetToken())
	if err != nil {
		return &pb.UpdateProfileResponse{Err: 1}, err
	}
	return &pb.UpdateProfileResponse{Err: 0}, err
}
