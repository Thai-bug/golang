package main

import (
	routes "demo/Routes"
	"demo/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func main() {
	db.Init()
	// fmt.Print(db.DB)
	getRoutes()
	router.Run("localhost:8080")
}

func getRoutes() {
	v1 := router.Group("/api/v1")
	routes.UserRoute(v1)
}
