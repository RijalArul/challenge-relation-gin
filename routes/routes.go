package routes

import "github.com/gin-gonic/gin"

func Routes() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Run()
}
