package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"ink-web/src/config"
	"ink-web/src/model"
	"ink-web/src/router"
	"ink-web/src/util/log"
	"runtime"
	"runtime/debug"
	"strings"
)

var (
	//配置文件 可通过flag定义路径
	cfg = flag.String("config", "", "lp config file path.")
)

var ctx = context.Background()

func main() {

	//0 异常处理
	defer func() {
		if err := recover(); err != nil {
			log.WithContext(ctx).WithError(errors.Errorf("%v", err)).Error("project run failed, stack: %s", string(debug.Stack()))
		}
	}()

	if cpu := runtime.NumCPU(); cpu == 1 {
		runtime.GOMAXPROCS(2)
	} else {
		runtime.GOMAXPROCS(cpu)
	}
	//1 初始化模块配置 读取配置
	if err := config.Config(ctx, *cfg); err != nil {
		log.WithContext(ctx).Error("init.setupSetting err: %s", err)
		panic(err)
	}
	//log.Info(config.ServerSetting.HttpPort)

	log.Info("LogsConfiguration.Category:%s", config.LogsConfiguration.Category)
	// get runtime root
	_, file, _, _ := runtime.Caller(0)
	config.RuntimeRoot = strings.TrimSuffix(file, "main.go")
	log.Info("程序运行根目录:%s", config.RuntimeRoot)
	logConfigStr, _ := json.Marshal(config.LogsConfiguration)
	sysConfigStr, _ := json.Marshal(config.ServerSetting)
	dbConfigStr, _ := json.Marshal(config.DatabaseSetting)
	log.Info("日志配置 %s", string(logConfigStr))
	log.Info("系统配置 %s", string(sysConfigStr))
	log.Info("数据库配置 %s", string(dbConfigStr))
	//1 log设置
	// change default logger
	log.DefaultWrapper = log.NewWrapper(log.New(
		log.WithCategory(config.LogsConfiguration.Category),
		log.WithLevel(config.LogsConfiguration.Level),
		log.WithJson(config.LogsConfiguration.Json),
		log.WithLineNumPrefix(config.RuntimeRoot),
		log.WithLineNum(!config.LogsConfiguration.LineNum.Disable),
		log.WithLineNumLevel(config.LogsConfiguration.LineNum.Level),
		log.WithLineNumVersion(config.LogsConfiguration.LineNum.Version),
		log.WithLineNumSource(config.LogsConfiguration.LineNum.Source),
	))

	err := model.InitDataBases(ctx)
	if err != nil {
		log.Error("数据库初始化失败")
	}
	//日志颜色
	gin.ForceConsoleColor()
	g := gin.Default()
	//1 加载路由api
	g = router.RegisterServers(g)
	err2 := g.Run(config.ServerSetting.HttpPort)
	if err2 != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务

	//logger.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	//logger.Infof("Start to listening the incoming requests on http address: %s","127.0.0.1")
	//logger.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
