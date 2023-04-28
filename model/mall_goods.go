package model

import (
	"fmt"
	"gorm.io/gorm"
)

type NftMallGoods struct {
	ID            uint   `json:"id" gorm:"column:id"`
	Name          string `json:"name" gorm:"column:name"`                     // 藏品名称
	Image         string `json:"image" gorm:"column:image"`                   // 藏品首图
	Detail        string `json:"detail" gorm:"column:detail"`                 // 藏品介绍
	Price         uint   `json:"price" gorm:"column:price"`                   // 藏品出售价格
	OriginalPrice uint   `json:"original_price" gorm:"column:original_price"` // 购买价格
	PlatformID    uint   `json:"platform_id" gorm:"column:platform_id"`       // 所属平台
	No            string `json:"no" gorm:"column:no"`                         // 编号
	Source        string `json:"source" gorm:"column:source"`                 // 藏品资源
	SourceType    uint8  `json:"source_type" gorm:"column:source_type"`       // 资源类型 1 图片 2 3D
	Hash          string `json:"hash" gorm:"column:hash"`                     // HASH地址
	Label         string `json:"label" gorm:"column:label"`                   // Label
	Status        uint8  `json:"status" gorm:"column:status"`                 // 是否可售 1是 0下架 2 已售
	Type          uint8  `json:"type" gorm:"column:type"`                     //类型 1普通藏品 2 盲盒
	UserId        string `json:"user_id" gorm:"column:user_id"`
	Base
}

func (m *NftMallGoods) TableName() string {
	return "nft_mall_goods"
}

func (m *NftMallGoods) Table() *gorm.DB {
	return m.Base.Table(m.TableName())
}

func (m *NftMallGoods) List(status []int, sortColumn string, sort string, offset, size int) (*[]NftMallGoods, error) {
	var list []NftMallGoods
	err := m.Table().
		Where("status in ?", status).
		Offset(offset).Limit(size).
		Order(fmt.Sprintf("%s %s", sortColumn, sort)).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}
func (m *NftMallGoods) Count(status []int) (int64, error) {
	var count int64
	err := m.Table().Where("status in ?", status).Count(&count).Error
	return count, err
}

func (m *NftMallGoods) FindById(id uint) (*NftMallGoods, error) {
	var info NftMallGoods
	err := m.Table().
		Where("id=?", id).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}
