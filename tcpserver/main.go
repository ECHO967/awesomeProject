package main

import (
	global2 "awesomeProject/config/global"
	"awesomeProject/pkg/logger"
	pb "awesomeProject/pkg/proto"
	"awesomeProject/pkg/setting"
	"awesomeProject/tcpserver/model"
	"awesomeProject/tcpserver/server"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net"
	"time"
)

var port string

// 服务端初始化
func init() {
	flag.StringVar(&port, "p", "8001", "启动端口号")
	flag.Parse()

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupReEngine()
	if err != nil {
		log.Fatalf("init.setupReENgin err: %v", err)
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
	err = setting.ReadSection("TCPLogger", &global2.TCPLogSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global2.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global2.DataBaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Redis", &global2.RedisSetting)
	if err != nil {
		return err
	}
	global2.JWTSetting.Expire *= time.Second
	return nil
}

// log配置
func setupLogger() error {
	fileName := global2.TCPLogSetting.LogSavePath + "/" + global2.TCPLogSetting.LogFileName + global2.TCPLogSetting.LogFileExt
	global2.TCPLogger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

// 载入数据库配置
func setupDBEngine() error {
	var err error
	global2.DBEngin, err = model.NewDBEngine(global2.DataBaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 载入缓存配置
func setupReEngine() error {
	var err error
	global2.ReEngin = redis.NewClient(&redis.Options{
		Addr:     global2.RedisSetting.Host,
		Password: global2.RedisSetting.Password,
		DB:       global2.RedisSetting.DB,
	})
	_, err = global2.ReEngin.Ping().Result()
	return err
}

// RPC服务端主函数
func main() {
	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, server.NewUserServer())
	fmt.Println("Rpc-Server Main Func Success")
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err : %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("s.Serve err :%v", err)
	}
}
