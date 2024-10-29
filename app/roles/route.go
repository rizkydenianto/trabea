package roles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	roles, err := GetRolesData()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "GAGAL")
	}
	c.IndentedJSON(http.StatusOK, roles)
}
