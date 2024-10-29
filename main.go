package main

import (
	"trabea/app/roles"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/login/roles", roles.GetRoles)

	router.Run("localhost:8080")
}
