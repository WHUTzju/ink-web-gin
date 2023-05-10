package vm

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// MakeSuccess return success response
func MakeSuccess(c *gin.Context, code int, content interface{}) {
	c.JSON(http.StatusOK, gin.H{"statusCode": code, "data": content})
}

// MakeFail return fail response
func MakeFail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{"statusCode": code, "message": message})
}

type InvokeArgs struct {
	Type       string `form:"type" json:"type" des:"类型"`
	Value      string `form:"value" json:"value" des:"值"`
}