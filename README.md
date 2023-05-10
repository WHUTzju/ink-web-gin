## Go-web 脚手架

参考：https://github.com/piupuer/gin-web

模块
框架：Gin
日志：logrus
配置：viper


目录结构
```
├── conf
│   └── config.yaml  #配置文件
├── go.mod #包管理
├── go.sum
├── pkg
├── README.md
└── src
    ├── config
    │   ├── config.go    #配置读取
    │   ├── dbConfig.go  #数据库配置信息
    │   ├── logConfig.go #日志配置
    │   └── sysConfig.go #系统配置
    ├── global
    │   └── global.go    #全局常量管理 
    ├── main.go          #入口函数
    ├── model
    │   ├── model.go     #
    │   ├── mysql.go     # 数据库初始化
    │   └── user.go      # 表：用户
    ├── router
    │   ├── middleware   # 中间件
    │   ├── router.go    # 路由
    │   ├── system       # 系统
    │   └── vm           # 
    ├── service
    └── util
        └── log          # 日志
```
