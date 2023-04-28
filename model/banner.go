package model

import (
	"fmt"
	"gorm.io/gorm"
	"nft_platform/global"
)

type NftBanner struct {
	ID       uint   `json:"id" gorm:"column:id"`
	Title    string `json:"title" gorm:"column:title"`         // 标题
	Image    string `json:"image" gorm:"column:image"`         // banner图片
	LinkAddr string `json:"link_addr" gorm:"column:link_addr"` // 链接地址
	Status   uint8  `json:"status" gorm:"column:status"`       // 是否启用 1 启用 2 下架
	Mark     string `json:"mark" gorm:"column:mark"`           // 备注
	Type     int8   `json:"type" gorm:"column:type"`           // 类型 1、自营平台 2、转售平台
	Base
}

func init() {
	fmt.Println("init")
}

func (m *NftBanner) TableName() string {
	return "nft_banner"
}

func (m *NftBanner) Table() *gorm.DB {
	return m.Base.Table(m.TableName())
}

func (m *NftBanner) List(t int) (*[]NftBanner, error) {
	var bannerList []NftBanner
	err := m.Table().
		Where("type=?", t).
		Find(&bannerList).
		Order("id DESC").Error
	if err != nil {
		global.SLogger.Error("banner list err:%s", err)
		return nil, err
	}
	return &bannerList, nil
}

func (m *NftBanner) Add(data NftBanner) error {
	return m.Table().Create(&data).Error
}

func (m *NftBanner) Update(data NftBanner, id int) error {
	return m.Table().Where("id=?", id).Updates(&data).Error
}

func (m *NftBanner) Delete(id int) error {
	var banner NftBanner
	return m.Table().Where("id=?", id).Delete(&banner).Error
}
