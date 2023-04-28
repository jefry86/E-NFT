package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type NftUsers struct {
	ID               uint      `json:"id" gorm:"column:id"`
	UserID           string    `json:"user_id" gorm:"column:user_id"`                       // 用户ID
	Username         string    `json:"username" gorm:"column:username"`                     // 账号
	Password         string    `json:"password" gorm:"column:password"`                     // 密码
	Nickname         string    `json:"nickname" gorm:"column:nickname"`                     // 昵称
	Mobile           string    `json:"mobile" gorm:"column:mobile"`                         // 手机号
	Avatar           string    `json:"avatar" gorm:"column:avatar"`                         // 头像
	WalletAddr       string    `json:"wallet_addr" gorm:"column:wallet_addr"`               // 钱包地址
	WalletPrivateKey string    `json:"wallet_private_key" gorm:"column:wallet_private_key"` // 钱包私钥
	Status           int8      `json:"status" gorm:"column:status"`                         // 状态，1 正常 0 冻结
	HasRealAuth      int8      `json:"has_real_auth" gorm:"column:has_real_auth"`           // 是否实名认证
	RealName         string    `json:"real_name" gorm:"column:real_name"`                   // 真实姓名
	CardNo           string    `json:"card_no" gorm:"column:card_no"`                       // 身份证
	Amount           uint      `json:"amount" gorm:"column:amount"`                         // 账号余额
	FreezeAmount     uint      `json:"freeze_amount" gorm:"column:freeze_amount"`           // 冻结金额
	Token            string    `json:"token" gorm:"column:token"`
	Birthday         time.Time `json:"birthday" gorm:"column:birthday"`
	Earnings         uint      `json:"earnings" gorm:"column:earnings"`
	Base
}

func (m *NftUsers) TableName() string {
	return "nft_users"
}

func (m *NftUsers) Table() *gorm.DB {
	return m.Base.Table(m.TableName())
}

func (m *NftUsers) FindByMobile(mobile string) (*NftUsers, error) {
	var info NftUsers
	err := m.Table().Where("mobile=?", mobile).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (m *NftUsers) FindByUserId(userId string) (*NftUsers, error) {
	var info NftUsers
	err := m.Table().Where("user_id=?", userId).First(&info).Error
	if err != nil {
		return nil, err
	} else if info.UserID == "" {
		return nil, fmt.Errorf("用户不存在！")
	}
	return &info, nil
}

func (m *NftUsers) FindByToken(token string) (*NftUsers, error) {
	var info NftUsers
	err := m.Table().Where("token=?", token).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (m *NftUsers) ExistByMobile(mobile string) (bool, error) {
	var count int64
	err := m.Table().Where("mobile=?", mobile).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, err
}

func (m *NftUsers) Add(data NftUsers) error {
	return m.Table().Create(&data).Error
}

func (m *NftUsers) UpdateByMobile(users NftUsers, mobile string) error {
	return m.Table().Where("mobile=?", mobile).Updates(&users).Error
}

func (m *NftUsers) UpdateByUserId(user NftUsers, userId string) error {
	return m.Table().Where("user_id=?", userId).Updates(&user).Error
}
