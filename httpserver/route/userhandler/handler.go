package userhandler

import (
	"awesomeProject/config/global"
	"awesomeProject/httpserver/grpc"
	"awesomeProject/pkg/common"
	"awesomeProject/pkg/errcode"
	pbr "awesomeProject/pkg/proto/err"
	"awesomeProject/pkg/response"
	"context"
	"net/http"
	//"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// 登录 API层
func (u User) LoginHandler(c *gin.Context) {
	//get para(username&password)
	para := grpc.LoginRequest{}
	//返回结果跟参数
	resp := response.NewResponse(c)
	err := c.ShouldBind(&para)
	if err != nil {
		global.HTTPLogger.Errorf("shoudbind errs: %v", err)
		errsp := errcode.InvalidParams.WithDetail(err.Error())
		resp.ToErrorResponse(errsp)
		return
	}
	//连接grpc
	ctx := context.Background()
	svc := grpc.New(ctx)
	user, token, err := svc.Login(&para)
	if err != nil {
		sts := errcode.FromError(err)
		if len(sts.Details()) == 0 {
			global.HTTPLogger.Errorf("svc.Login err: %v", err)
			resp.ToErrorResponse(errcode.ErrorGetUserFail.WithDetail(err.Error()))
			return
		}
		detail := sts.Details()[0].(*pbr.Error)
		if int(detail.Code) == errcode.ErrorPassword.Code() {
			//根据判断code内容，进入不同的error处理界面
			global.HTTPLogger.Errorf("svc.Login err: %v", err)
			c.HTML(http.StatusOK, "errorpass.html", gin.H{}) //用户名跟密码不正确重新登录
			return
		} else {
			global.HTTPLogger.Errorf("svc.Login err: %v", err)
			resp.ToErrorResponse(errcode.ErrorGetUserFail.WithDetail(err.Error()))
			return
		}
	}
	//设置cookie
	c.SetCookie("token", token, 7200, "", "127.0.0.1", false, true)
	if user.Profile == "" {
		user.Profile = "http://127.0.0.1:8080/static/Initial.webp"
	}

	c.HTML(http.StatusOK, "show.html", gin.H{"username": user.Username, "nickname": user.Nickname, "profile": user.Profile})
	resp.ToResponse("Login Success")

	return
}

// 修改nickname API层
func (u User) UpdateNickname(c *gin.Context) {
	para := common.UpdateInfo{}
	resp := response.NewResponse(c)
	//登录后具有sessionID信息，请求中带有session_id，通过sessionID查询用户信息
	para.Token, _ = c.Cookie("token")
	para.Username, _ = c.GetQuery("username")
	para.Profile, _ = c.GetQuery("profile")
	para.Nickname = c.PostForm("nickname")
	ctx := context.Background()
	svc := grpc.New(ctx)
	err := svc.UpdateNic(&para)
	if err != nil {
		sts := errcode.FromError(err)
		detail := sts.Details()[0].(*pbr.Error)
		if int(detail.Code) == errcode.ErrorAuth.Code() {
			global.HTTPLogger.Errorf("svc.ChangeNick err: %v", err)
			c.HTML(http.StatusOK, "errorpass.html", gin.H{})
			return
		} else {
			global.HTTPLogger.Errorf("svc.ChangeNick err: %v", err)
			resp.ToErrorResponse(errcode.ErrorChangeNickFail.WithDetail(err.Error()))
			return
		}
	}
	if para.Profile == "" {
		para.Profile = "http://127.0.0.1:8080/static/Initial.webp"
	}
	c.HTML(http.StatusOK, "show.html", gin.H{"username": para.Username, "nickname": para.Nickname, "profile": para.Profile})
	return

}

// 修改头像 API层
func (u User) UpdatePic(c *gin.Context) {
	para := common.UpdateInfo{}
	resp := response.NewResponse(c)
	//登录后具有sessionID信息，请求中带有session_id，通过sessionID查询用户信息
	para.Username = c.Query("username")
	para.Nickname = c.Query("nickname")
	para.Token, _ = c.Cookie("token")
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil || fileHeader == nil {
		global.HTTPLogger.Errorf("FormFile errs: %v", err)
		esp := errcode.InvalidParams.WithDetail(err.Error())
		resp.ToErrorResponse(esp)
		return
	}
	ctx := context.Background()
	svc := grpc.New(ctx)
	fileInfo, err := svc.UploadFile(para.Username, file, fileHeader)
	if err != nil {
		global.HTTPLogger.Errorf("svc.UploadFile err: %v", err)
		errRsp := errcode.ErrorChangeProfFail.WithDetail(err.Error())
		resp.ToErrorResponse(errRsp)
	}
	para.Profile = fileInfo
	err = svc.UpdateProf(&para)
	if err != nil {
		sts := errcode.FromError(err)
		if len(sts.Details()) > 0 {
			detail := sts.Details()[0].(*pbr.Error)

			//detail := sts.Details()[0].(*pbr.Error)
			if int(detail.Code) == errcode.ErrorAuth.Code() {
				global.HTTPLogger.Errorf("svc.ChangeProf err: %v", err)
				c.HTML(http.StatusOK, "errorpass.html", gin.H{})
				return
			} else {
				global.HTTPLogger.Errorf("svc.ChangeProf err: %v", err)
				resp.ToErrorResponse(errcode.ErrorChangeProfFail.WithDetail(err.Error()))
				return
			}
		}
	}
	c.HTML(http.StatusOK, "show.html", gin.H{"username": para.Username, "nickname": para.Nickname, "profile": para.Profile})
	return

}
