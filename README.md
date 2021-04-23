#### go_gin_web 练习
go version :1.16.3

~~~~
#本地项目绑定线上仓库
`git remote add origin git@github.com:jierui303-byte/go_gin_web`
~~~~

~~~~
#Gin包安装命令：
`go get -u github.com/gin-gonic/gin`
~~~~

~~~~
#session功能集成-安装session库[session存储redis]
`go get -u github.com/gin-contrib/sessions`
`go get -u github.com/gin-contrib/sessions/redis`
~~~~


~~~~
#安装redis
`go get -u github.com/go-redis/redis`
#安装mysql驱动
`go get "github.com/go-sql-driver/mysql"`
#安装mongodb
`go get gopkg.in/mgo.v2`
~~~~

~~~~
#接入阿里云短信平台
`go get -u github.com/aliyun/alibaba-cloud-sdk-go`
~~~~

~~~~
#安装xorm操作数据库的orm框架
`go get -u github.com/go-xorm/xorm`
`go mod download github.com/jmespath/go-jmespat`
`go get github.com/micro/go-micro/store/service`
`go get github.com/micro/go-micro/v2/logger`
`go mod download github.com/modern-go/concurrent`
`go mod download gopkg.in/ini.v1`
~~~~

~~~~
#图形验证码[支持生成的验证码自动存储redis以及验证码校验方式]
`go get -u github.com/mojocn/base64Captcha`
#安装最新版本验证码包会有部分函数不支持，需降到1.2.2版本
`go mod edit -replace=github.com/mojocn/base64Captcha@v1.3.1=github.com/mojocn/base64Captcha@v1.2.2`
~~~~

~~~~
#fastDFS安装【分布式文件存储系统】
1:fastDFS安装，配置和服务启动
2:配置nginx模块
3:go语言编程上传fastDFS
#安装fastDFS包
`go get -u github.com/tedcy/fdfs_client`
~~~~