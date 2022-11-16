package models

import "gorm.io/gorm"

type Meta struct {
	gorm.Model
	IsMeta   bool `gorm:"type:bool;default:false"`
	HeroName string
}
