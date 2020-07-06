package gorm

import (
	"github.com/jinzhu/gorm"
	"go-admin/models"
	"go-admin/models/tools"
)

func AutoMigrate(db *gorm.DB) error {
	db.SingularTable(true)
	err:=db.AutoMigrate(
		new(models.CasbinRule),
		new(models.SysDept),
		new(models.SysConfig),
		new(tools.SysTables),
		new(tools.SysColumns),
		new(models.Menu),
		new(models.LoginLog),
		new(models.SysOperLog),
		new(models.RoleMenu),
		new(models.SysRoleDept),
		new(models.SysUser),
		new(models.SysRole),
		new(models.Post),
		new(models.DictData),
		new(models.DictType),
	).Error

	models.DataInit()
	return err
}


func CustomMigrate(db *gorm.DB,table interface{}) error {
	db.SingularTable(true)
	return db.AutoMigrate(&table).Error
}

