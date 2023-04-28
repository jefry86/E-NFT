package model

import (
	"gorm.io/gorm"
	"nft_platform/global"
	"nft_platform/utils"
)

type Base struct {
	DtCreate uint `json:"-" gorm:"column:dt_create"` // 创建时间
	DtUpdate uint `json:"-" gorm:"column:dt_update"` // 跟新时间
}

func (b *Base) SetDtCreate() {
	b.DtCreate = utils.NowUnix()
}

func (b *Base) SetDtUpdate() {
	b.DtUpdate = utils.NowUnix()
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.SetDtUpdate()
	b.SetDtUpdate()
	return nil
}

func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	b.SetDtUpdate()
	return nil
}

func (b *Base) Table(tableName string) *gorm.DB {
	return global.DB.Table(tableName)
}
