package routes

import (
	"challenge-relation-gin/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	r.POST("/heroes", controllers.CreateHero)

	r.Run()
}
