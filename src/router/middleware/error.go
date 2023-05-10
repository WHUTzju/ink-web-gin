package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ink-web/src/util/log"
	"net/http"
	"runtime/debug"
)

// CatchError catch gin error
func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ers := "系统错误\r\n"
				ers += fmt.Sprintf(fmt.Sprintf("%v\r\n", err)) // output panic info
				ers += fmt.Sprintf("========\r\n")
				ers += fmt.Sprintf(string(debug.Stack())) // output stack info
				log.Error(string(debug.Stack()))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
