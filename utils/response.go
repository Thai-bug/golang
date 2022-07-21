package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code       string
	StatusCode int
	Message    string
	Result map[string]interface{} 
}

func GetResponse(c *gin.Context, params Response) {
	fmt.Print("hello")
	c.JSON(params.StatusCode, gin.H{
		"code":    params.Code,
		"message": params.Message,
		"result":  params.Result,
	})
	c.Abort()
}
