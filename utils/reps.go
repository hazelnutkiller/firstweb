package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code int, msg string, err error) {

	statusCode := code

	// check if dedicated status code
	errorResponseCode := c.PostForm("errorResponseCode")
	if errorResponseCode != "" && errorResponseCode == "200" {
		s, e := strconv.Atoi(errorResponseCode)
		if e == nil {
			statusCode = s
		}
	}

	errorMsg := msg
	if err != nil {
		errorMsg = fmt.Sprintf("%s: %v", msg, err)
	}
	c.Set("ErrorMsg", errorMsg)
	c.JSON(statusCode, gin.H{"error": msg, "code": code})
}
