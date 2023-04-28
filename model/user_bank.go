package model

import "gorm.io/gorm"

type NftUserBank struct {
	ID          uint   `json:"id" gorm:"column:id"`
	UserID      string `json:"user_id" gorm:"column:user_id"`           // 用户Id
	Bank        string `json:"bank" gorm:"column:bank"`                 // 开户行
	BankName    string `json:"bank_name" gorm:"column:bank_name"`       // 银行户名
	BankAccount string `json:"bank_account" gorm:"column:bank_account"` // 银行账号
	BankAddress string `json:"bank_address" gorm:"column:bank_address"` // 开户地址
	Status      uint8  `json:"status" gorm:"column:status"`             // 是否可用 1 可用 0 不可用
	Base
}

func (m *NftUserBank) TableName() string {
	return "nft_user_bank"
}

func (m *NftUserBank) Table() *gorm.DB {
	return m.Base.Table(m.TableName())
}

func (m *NftUserBank) ListByUserId(userId string) (*[]NftUserBank, error) {
	var list []NftUserBank
	err := m.Table().Where("user_id=?", userId).Where("status=1").Order("dt_create DESC").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (m *NftUserBank) FindById(id int) (*NftUserBank, error) {
	var info NftUserBank
	err := m.Table().Where("id=?", id).Where("status=1").Order("dt_create DESC").First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (m *NftUserBank) Add(data NftUserBank) error {
	return m.Table().Create(&data).Error
}

func (m *NftUserBank) Delete(id int, userId string) error {
	var model NftUserBank
	return m.Table().Where("id=?", id).Where("user_id=?", userId).Delete(&model).Error
}
