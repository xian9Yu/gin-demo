package cmd

import (
	"github.com/gin-gonic/gin"
	"jwt/middleware"
	"jwt/models"
	"jwt/router"
	"jwt/utils"
	"log"
)

var conf = utils.NewCfg().InitConfig() // 初始化配置文件

func Cmd() {

	models.DB = models.InitSQL()     //初始化sql
	models.Rdb = models.InitClient() //初始化 redis

	//初始化router
	r := gin.Default()
	router.InitRouter(r)
	//pprof.Register(r)
	//使用gin自带的异常恢复中间件，避免出现异常时程序退出
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	host := conf.GetString("ListenOn.Host")
	port := conf.GetString("ListenOn.Port")
	err := r.Run(host + ":" + port)
	if err != nil {
		log.Fatal("服务启动失败 ：", err)
	}
}