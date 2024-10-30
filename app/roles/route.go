package roles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	roles, err := GetRolesData()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if roles == nil {
		c.IndentedJSON(http.StatusNoContent, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, roles)
}
