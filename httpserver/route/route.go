package route

import (
	"awesomeProject/config/global"
	"awesomeProject/httpserver/route/userhandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建路由规则
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())   //Logger将日志写入gin.DefaultWriter
	r.Use(gin.Recovery()) //Recovery会recover任何panic，返回500

	user := userhandler.NewUser()
	r.StaticFS("/login", http.Dir("../html/"))
	r.StaticFS("/static", http.Dir(global.UploadSetting.UploadSavePath))
	api := r.Group("/api") //将api包裹起来看起来更清晰
	//GET跟POST请求
	api.GET("/home", func(context *gin.Context) {
		context.HTML(http.StatusOK, "home.html", nil)
	})

	api.POST("/user/login", user.LoginHandler)
	api.POST("/user/nick", user.UpdateNickname)
	api.POST("/user/prof", user.UpdatePic)
	return r
}
