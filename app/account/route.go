package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req ReqLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	res, err := GetLoginData(req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if res == nil {
		c.IndentedJSON(http.StatusNoContent, gin.H{})
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}
