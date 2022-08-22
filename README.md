# User management system
## 简介
`eUser management system`实现一个用户管理系统，用户可以登录、拉取和编辑他们的profiles。

## 文档
你可以从 [EntryTask Report](https://docs.google.com/document/d/1gH2gl-HlCo6TcvrccGWOzt1Md9NvVZVu370W8cd5SgE/edit?usp=sharing#)中查阅以下内容:
- 项目内容与设计要求
- 系统设计，包括API接口，数据库设计，grpc设计
- 配置部署
- 压测结果

## 配置运行
在启动之前, 你需要在本地安装 **MySQL** and **Redis**。 
 `go.mod` 中查看本项目使用到的第三方库。 
 按照文档中的部署来初始化数据库。


```bash
# 修改配置文件
vim ./config/config.yaml

# 处理第三方库依赖
go mod tidy

# 启动 RPC server，进入tcpserver目录下
go build
./tcpserver

# 启动 HTTP server,进入httpserver目录下
go build 
./httpserver
```

最后测试一下是否正确运行:
```bash
curl --location --request GET 'http://127.0.0.1:8080/api/home'
```
可以在浏览器上打开进行用户信息登录修改操作
