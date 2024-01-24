package helpers

import (
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, message interface{}, respCode string) {
	c.AbortWithStatusJSON(code, gin.H{"resp_desc": message, "resp_code": respCode})
}

func RespondWithSuccess(c *gin.Context, code int, message interface{}, respCode string, data ...interface{}) {
	response := gin.H{"resp_desc": message, "resp_code": respCode}

	if len(data) > 0 {
		response["data"] = data[0]
	}

	c.JSON(code, response)
}
