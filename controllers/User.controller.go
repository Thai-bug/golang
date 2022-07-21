package controllers

import (
	models "demo/Models"
	"demo/db"
	"demo/utils"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	users := []models.User{}
	var count int64
	db.DB.Find(&users).Count(&count)
	utils.GetResponse(context, utils.Response{
		Code: "OK",
		StatusCode: 200, 
		Message: "",
		Result: map[string]interface{}{
			"content": users,
			"total": count,
		}})
	return
}
