package main

import (
	"fmt"
	"web_app/controllers"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/pkg/snowflake"
	"web_app/routers"
	"web_app/settings"

	"go.uber.org/zap"
)

//Go Web开发通用模版

// @title BlueBell
// @version 1.0
// @description 第一个gin项目
// asd
// @termOfService https://www.baidu.com

// @Host 127.0.0.1:8081
// @BasePath /api/v1

func main() {
	//1. 加载配置
	err := settings.Init()
	if err != nil {
		zap.L().Debug("Init settings failed", zap.Error(err))
		fmt.Printf("错误:%v\n", err)
		return
	}
	//2. 初始化日志
	err = logger.Init(settings.Cfg.LoggerConf, settings.Cfg.Mode)
	if err != nil {
		zap.L().Debug("Init log failed", zap.Error(err))
		return
	}
	defer zap.L().Sync() //将缓冲区的日志追加到文件中

	zap.L().Debug("开始")
	//3. 初始化Mysql
	err = mysql.Init(settings.Cfg.MysqlConf)
	if err != nil {
		zap.L().Debug("Init mysql failed", zap.Error(err))
		return
	}
	defer mysql.Close()

	//4. 初始化Redis
	err = redis.Init(settings.Cfg.RedisConf)
	if err != nil {
		zap.L().Debug("Init redis failed", zap.Error(err))
		return
	}
	defer redis.Close()

	//5. 注册路由
	r := routers.Setup()

	// ID生成器
	snowflake.Init(settings.Cfg.StartTime, settings.Cfg.MachineId)

	// 初始化gin内置的校验器翻译器
	err = controllers.InitTrans("zh")
	if err != nil {
		zap.L().Error("Init validator tran failed", zap.Error(err))
		return
	}
	//6. 启动服务(优雅关机)

	r.Run(fmt.Sprintf(":%d", settings.Cfg.Port))

	//优雅关机
	//quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	//
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	//<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	//
	//zap.L().Info("Shutdown Server ...")
	//// 创建一个5秒超时的context
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	//if err := srv.Shutdown(ctx); err != nil {
	//	zap.L().Fatal("Server Shutdown", zap.Error(err))
	//}
	//
	//log.Println("Server exiting")
	//select {}

}
