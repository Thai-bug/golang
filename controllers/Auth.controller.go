package controllers

import (
	"demo/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyJWT(c *gin.Context) {
	headers := c.Request.Header.Clone()
	token := strings.Split(headers.Get("Authorization"), " ")
	if len(token) < 2 {
		utils.GetResponse(c, utils.Response{
			"INVALID_TOKEN",
			401, "",
			map[string]interface{}{
				"hello": "hello",
			}})
		return
	}
}
