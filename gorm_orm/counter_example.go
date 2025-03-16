package gorm_orm

import (
	"gorm.io/gorm"
)

type AutoIncrementCounter struct {
	gorm.Model
	Value int "gorm:autoIncrement"
}

//{Model:{ID:1 CreatedAt:2025-03-13 11:24:25.520314 +0300 EAT UpdatedAt:2025-03-13 11:24:25.520314 +0300 EAT DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Value:0}
