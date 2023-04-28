package model

import (
	"gorm.io/gorm"
	"nft_platform/global"
)

type NftUserWithdraw struct {
	ID          uint   `json:"id" gorm:"column:id"`
	UserID      string `json:"user_id" gorm:"column:user_id"`           // 用户编号
	Amount      uint   `json:"amount" gorm:"column:amount"`             // 提现金额
	Bank        string `json:"bank" gorm:"column:bank"`                 // 开户行
	BankName    string `json:"bank_name" gorm:"column:bank_name"`       // 户名
	BankAccount string `json:"bank_account" gorm:"column:bank_account"` // 账号
	BankAddr    string `json:"bank_addr" gorm:"column:bank_addr"`       // 开户地址
	Status      uint8  `json:"status" gorm:"column:status"`             // 状态，1 审核中 2 待打款 3 已打款 4 撤销
	Base
}

func (m *NftUserWithdraw) TableName() string {
	return "nft_user_withdraw"
}

func (m *NftUserWithdraw) Table() *gorm.DB {
	return global.DB.Table(m.TableName())
}

func (m *NftUserWithdraw) ListByUserId(userId string, offset, size int) (*[]NftUserWithdraw, error) {
	var list []NftUserWithdraw
	err := m.Table().Where("user_id=?", userId).Offset(offset).Limit(size).Order("dt_create DESC").Find(&list).Error
	return &list, err
}

func (m *NftUserWithdraw) CountByUserId(userId string) (int64, error) {
	var count int64
	err := m.Table().Where("user_id=?", userId).Count(&count).Error
	return count, err
}

func (m *NftUserWithdraw) Add(data NftUserWithdraw) error {
	return m.Table().Create(&data).Error
}
