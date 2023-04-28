package model

import "gorm.io/gorm"

type NftMallUserGoods struct {
	ID             uint         `json:"id" gorm:"column:id"`
	GoodsID        uint         `json:"goods_id" gorm:"column:goods_id"`                 // 藏品ID
	GoodsName      string       `json:"goods_name" gorm:"column:goods_name"`             // 藏品名称
	GoodsNo        string       `json:"goods_no" gorm:"column:goods_no"`                 // 藏品编号
	FromUserID     string       `json:"from_user_id" gorm:"column:from_user_id"`         // 出售用户ID
	MallPlatformID uint         `json:"mall_platform_id" gorm:"column:mall_platform_id"` // 平台ID
	UserID         string       `json:"user_id" gorm:"column:user_id"`                   // 用户编号
	GoodsHash      string       `json:"goods_hash" gorm:"column:goods_hash"`             // 藏品HASh
	Status         uint8        `json:"status" gorm:"column:status"`                     // 是否可用
	Goods          NftMallGoods `json:"-" gorm:"foreignKey:goods_id"`
	Base
}

func (m *NftMallUserGoods) TableName() string {
	return "nft_mall_user_goods"
}

func (m *NftMallUserGoods) Table() *gorm.DB {
	return m.Base.Table(m.TableName())
}

func (m *NftMallUserGoods) List(userId string, t, offset, size int) (*[]NftMallUserGoods, error) {
	var list []NftMallUserGoods
	err := m.Table().Preload("Goods").
		Where("user_id=?", userId).Where("type=?", t).
		Offset(offset).Limit(size).
		Order("id d").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (m *NftMallUserGoods) Count(userId string, t int) (int64, error) {
	var count int64
	err := m.Table().Preload("Goods").
		Where("user_id=?", userId).Where("type=?", t).
		Order("id d").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *NftMallUserGoods) Add(info NftMallUserGoods) error {
	err := m.Table().Create(&info).Error
	return err
}

func (m *NftMallUserGoods) FindByIdAndUserId(id uint, userId string) (*NftMallUserGoods, error) {
	var info NftMallUserGoods
	err := m.Table().Where("id=?", id).Where("user_id=?", userId).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}
