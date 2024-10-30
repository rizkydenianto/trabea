package main

import (
	"trabea/app/account"
	"trabea/app/roles"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent"}
	router.Use(cors.New(config))

	router.GET("/login", roles.GetRoles)
	router.POST("/login", account.Login)

	router.Run("localhost:8080")
}
