package model

import (
	"gorm.io/gorm"
	"nft_platform/global"
)

type NftMallPlatform struct {
	ID           uint   `json:"id" gorm:"column:id"`
	Name         string `json:"name" gorm:"column:name"`                     // 平台名称
	Logo         string `json:"logo" gorm:"column:logo"`                     // 平台logo
	Site         string `json:"site" gorm:"column:site"`                     // 平台web地址
	ApiUser      string `json:"api_user" gorm:"column:api_user"`             // 登记api
	ApiGoodsList string `json:"api_goods_list" gorm:"column:api_goods_list"` // 藏品列表API
	ApiTransfer  string `json:"api_transfer" gorm:"column:api_transfer"`     // 藏品转移API
	ApiSales     string `json:"api_sales" gorm:"column:api_sales"`           // 藏品上架销售查询API
	Status       uint8  `json:"status" gorm:"column:status"`                 // 是否可用 1可用 0 不可用
	HasOnline    uint8  `json:"has_online" gorm:"column:has_online"`         // 是否已上线
	AppKey       string `json:"app_key" gorm:"column:app_key"`
	AppSecret    string `json:"app_secret" gorm:"column:app_secret"`
	Base
}

func (m *NftMallPlatform) TableName() string {
	return "nft_mall_platform"
}

func (m *NftMallPlatform) Table() *gorm.DB {
	return global.DB.Table(m.TableName())
}

func (m *NftMallPlatform) FindById(id int) (*NftMallPlatform, error) {
	var info NftMallPlatform
	err := m.Table().Where("id=?", id).First(&info).Error
	return &info, err
}

func (m *NftMallPlatform) FindByAppKey(appKey string) (*NftMallPlatform, error) {
	var info NftMallPlatform
	err := m.Table().Where("app_key=?", appKey).First(&info).Error
	return &info, err
}

func (m *NftMallPlatform) ListByStatus(status int, offset, pageSize int) (*[]NftMallPlatform, error) {
	var list []NftMallPlatform
	err := m.Table().Where("status=?", status).
		Offset(offset).Limit(pageSize).
		Order("sort DESC,dt_create DESC").
		Find(&list).Error
	return &list, err
}

func (m *NftMallPlatform) CountByStatus(status int) (int64, error) {
	var count int64
	err := m.Table().Where("status=?", status).
		Count(&count).Error
	return count, err
}
