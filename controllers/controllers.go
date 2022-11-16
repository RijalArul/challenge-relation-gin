package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateHero(ctx *gin.Context) {
	name, role := ctx.PostForm("name"), ctx.PostForm("role")

	fmt.Println(name, " ==== ", role)
}
