package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerCheck(c *gin.Context) {
	c.String(http.StatusOK, "service is ready")
}