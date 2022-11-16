package models

import "gorm.io/gorm"

type Hero struct {
	gorm.Model
	Name string `gorm:"unique;not null;type:VARCHAR(255)"`
	Role string `gorm:"not null;type:VARCHAR(255)"`
	Meta Meta   `gorm:"foreignKey:HeroName;references:name"`
}
