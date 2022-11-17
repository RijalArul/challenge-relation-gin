package models

import "gorm.io/gorm"

type Meta struct {
	gorm.Model
	IsMeta   bool `gorm:"default:false"`
	HeroName string
}
