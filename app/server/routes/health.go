package routes

import "github.com/gin-gonic/gin"

func healthHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	}

}
