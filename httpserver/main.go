package main

import (
	global2 "awesomeProject/config/global"
	"awesomeProject/httpserver/route"
	"awesomeProject/pkg/logger"
	"awesomeProject/pkg/setting"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// 初始化
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupConn()
	if err != nil {
		log.Fatalf("init.setupConn err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// 完成config中的配置
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("HTTPServer", &global2.HTTPServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("HTTPLogger", &global2.HTTPLogSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Upload", &global2.UploadSetting)
	if err != nil {
		return err
	}
	global2.HTTPServerSetting.ReadTimeout *= time.Second
	global2.HTTPServerSetting.WriteTimeout *= time.Second
	return nil
}

// 载入rpc客户端
func setupConn() error {
	ctx := context.Background()
	var err error
	global2.ClientConn, err = GetClientConn(ctx, "127.0.0.1:8001", nil)
	return err
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}

// log配置
func setupLogger() error {
	fileName := global2.HTTPLogSetting.LogSavePath + "/" + global2.HTTPLogSetting.LogFileName + global2.HTTPLogSetting.LogFileExt
	global2.HTTPLogger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func main() {
	gin.SetMode(global2.HTTPServerSetting.RunMode)
	router := route.NewRouter()
	//渲染html模版
	router.LoadHTMLGlob("../html/*")
	//服务器配置
	s := &http.Server{
		Addr:           ":" + global2.HTTPServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global2.HTTPServerSetting.ReadTimeout,
		WriteTimeout:   global2.HTTPServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	defer global2.ClientConn.Close()
	s.ListenAndServe()

}
