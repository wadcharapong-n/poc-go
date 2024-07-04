package repository

import (
	"gorm.io/gorm"
	"poc-go/app/domain/dao"
)

type RoleRepository interface {
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepositoryImpl {
	db.AutoMigrate(&dao.Role{})
	return &RoleRepositoryImpl{
		db: db,
	}
}
