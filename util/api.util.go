package util

import "github.com/gin-gonic/gin"

func SendError(c *gin.Context, status int, desc string) {
	c.JSON(status, gin.H{"error": desc})
}
