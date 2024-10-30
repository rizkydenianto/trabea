package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req ReqLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "GAGAL")
	}

	res, err := GetLoginData(req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "GAGAL")
	}

	c.IndentedJSON(http.StatusOK, res)
}
