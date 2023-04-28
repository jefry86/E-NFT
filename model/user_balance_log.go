package model

import (
	"gorm.io/gorm"
	"nft_platform/global"
)

type NftUserBalanceLog struct {
	ID         uint   `json:"id" gorm:"column:id"`
	UserID     string `json:"user_id" gorm:"column:user_id"`         // 用户编号
	Amount     int    `json:"amount" gorm:"column:amount"`           // 变动金额 + 出售 - 提现
	Balance    uint   `json:"balance" gorm:"column:balance"`         // 账号余额
	GoodsID    uint   `json:"goods_id" gorm:"column:goods_id"`       // 藏品ID
	GoodsName  string `json:"goods_name" gorm:"column:goods_name"`   // 藏品名称
	GoodsImage string `json:"goods_image" gorm:"column:goods_image"` // 藏品图片
	Info       string `json:"info" gorm:"column:info"`               // 说明
	Type       uint8  `json:"type" gorm:"column:type"`               // 类型 1 藏品转售 2 提现 3 支付
	Base
}

func (m *NftUserBalanceLog) TableName() string {
	return "nft_user_balance_log"
}

func (m *NftUserBalanceLog) Table() *gorm.DB {
	return global.DB.Table(m.TableName())
}

func (m *NftUserBalanceLog) Add(log NftUserBalanceLog) error {
	return m.Table().Create(&log).Error
}

func (m *NftUserBalanceLog) ListByUserId(userId string, offset, size int) (*[]NftUserBalanceLog, error) {
	var list []NftUserBalanceLog
	err := m.Table().Where("user_id=?", userId).Offset(offset).Limit(size).Order("dt_create DESC").Find(&list).Error
	return &list, err
}

func (m *NftUserBalanceLog) CountByUserId(userId string) (int64, error) {
	var count int64
	err := m.Table().Where("user_id=?", userId).Count(&count).Error
	return count, err
}
