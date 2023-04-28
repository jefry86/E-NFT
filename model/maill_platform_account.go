package model

import (
	"gorm.io/gorm"
	"nft_platform/global"
)

type NftMallPlatformAccount struct {
	ID                 uint   `json:"id" gorm:"column:id"`
	UserID             string `json:"user_id" gorm:"column:user_id"`                           // 用户编号
	PlatformWalletHash string `json:"platform_wallet_hash" gorm:"column:platform_wallet_hash"` // 钱包地址
	PlatformUserID     string `json:"platform_user_id" gorm:"column:platform_user_id"`         // 用户编号
	PlatformID         int    `json:"platform_id" gorm:"column:platform_id"`                   // 平台ID
	Status             int8   `json:"status" gorm:"column:status"`                             // 是否可用
	DtCreate           uint   `json:"dt_create" gorm:"column:dt_create"`                       // 创建时间
	DtUpdate           uint   `json:"dt_update" gorm:"column:dt_update"`                       // 更新时间
}

func (m *NftMallPlatformAccount) TableName() string {
	return "nft_platform_account"
}

func (m *NftMallPlatformAccount) Table() *gorm.DB {
	return global.DB.Table(m.TableName())
}

func (m *NftMallPlatformAccount) FindByUserId(userId string) (*NftMallPlatformAccount, error) {
	var info NftMallPlatformAccount
	err := m.Table().Where("user_id=?", userId).First(&info).Error
	return &info, err
}
