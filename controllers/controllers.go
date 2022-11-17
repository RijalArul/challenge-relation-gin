package controllers

import (
	"challenge-relation-gin/config"
	"challenge-relation-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHero(ctx *gin.Context) {
	name, role := ctx.PostForm("name"), ctx.PostForm("role")

	tx := config.GetDB().Begin()

	payloadHero := models.Hero{
		Name: name,
		Role: role,
	}

	payloadMeta := models.Meta{
		HeroName: payloadHero.Name,
	}

	var errMessage error

	if err := tx.Create(&payloadHero).Error; err != nil {
		errMessage = err
		tx.Rollback()

	}

	if err := tx.Create(&payloadMeta).Error; err != nil {
		tx.Rollback()
	}

	if errMessage != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": errMessage,
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Success create heroes",
			"hero":    payloadHero,
			"meta":    payloadMeta,
		})
	}

	tx.Commit()
}

// func GetALlHeroes(ctx *gin.Context) {
// 	payload := models.Hero{}
// 	result :=
// }
