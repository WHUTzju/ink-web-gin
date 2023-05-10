package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"ink-web/src/model"
	"ink-web/src/router/middleware"
	"ink-web/src/router/system"
	"net/http"
)

func RegisterServers(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	if viper.GetString("runmode") == "debug" {
		g.GET("/", func(c *gin.Context) {
			c.Redirect(301, "/swagger/index.html")
		})
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//系统健康检查
	svcd := g.Group("/system").Use(middleware.CatchError())
	{
		svcd.GET("/health", system.HealthCheck)
		svcd.GET("/cpu", system.CPUCheck)
		svcd.GET("/ram", system.RAMCheck)
	}

	// hello world
	hw := g.Group("/api/v1").Use(middleware.CatchError())
	{
		hw.GET("/hello/:name", func(c *gin.Context) {

			name := c.Params.ByName("name")
			user := model.GetUserByName(name)

			c.AsciiJSON(http.StatusOK, gin.H{
				"message": "hello world",
				"user":    user,
			})
		})
	}

	return g
}
