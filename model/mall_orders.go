package model

import (
	"fmt"
	"gorm.io/gorm"
)

type NftMallOrders struct {
	ID             uint         `json:"id" gorm:"column:id"`
	OrderID        string       `json:"order_id" gorm:"column:order_id"`                 // 订单号
	GoodsID        uint         `json:"goods_id" gorm:"column:goods_id"`                 // 藏品ID
	MallPlatformID uint         `json:"mall_platform_id" gorm:"column:mall_platform_id"` // 平台ID
	FromUserID     string       `json:"from_user_id" gorm:"column:from_user_id"`         // 出售用户编号
	UserID         string       `json:"user_id" gorm:"column:user_id"`                   // 用户ID
	Num            uint         `json:"num" gorm:"column:num"`                           // 购买数量
	Amount         uint         `json:"amount" gorm:"column:amount"`                     // 订单金额 单位 分
	PayChannelID   uint         `json:"pay_channel_id" gorm:"column:pay_channel_id"`     // 支付渠道
	PayEndtime     uint         `json:"pay_endtime" gorm:"column:pay_endtime"`           // 支付超时时间
	PayTime        uint         `json:"pay_time" gorm:"column:pay_time"`                 // 支付超时时间
	Status         uint8        `json:"status" gorm:"column:status"`                     // 订单状态 1待支付 2 已支付 3 已取消 4 超时 5 已完成
	Goods          NftMallGoods `json:"-" gorm:"foreignKey:goods_id"`
	Base
}

func (m *NftMallOrders) TableName() string {
	return "nft_mall_orders"
}

func (m *NftMallOrders) Table() *gorm.DB {
	return m.Base.Table(m.TableName())
}

func (m *NftMallOrders) List(userId string) (*[]NftMallGoods, error) {
	var list []NftMallGoods
	err := m.Table().Where("user_id=?", userId).
		Preload("Goods").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (m *NftMallOrders) Add(order NftMallOrders) error {
	return m.Table().Create(&order).Error
}

func (m *NftMallOrders) ChangeStatus(orderId, userId string, status int) error {
	err := m.Table().
		Where("order_id=?", orderId).
		Where("user_id=?", userId).
		Update("status", status).Error
	return err
}

func (m *NftMallOrders) UpdateByOrderId(data NftMallOrders, orderId string) error {
	return m.Table().Where("order_id=?", orderId).Updates(&data).Error
}

func (m *NftMallOrders) FindByOrderIdAndUserId(orderId, userId string) (*NftMallOrders, error) {
	var info NftMallOrders
	err := m.Table().
		Preload("Goods").
		Where("order_id=?", orderId).
		Where("user_id=?", userId).
		First(&info).Error
	return &info, err
}

func (m *NftMallOrders) FindByOrderId(orderId string) (*NftMallOrders, error) {
	var info NftMallOrders
	err := m.Table().
		Preload("Goods").
		Where("order_id=?", orderId).
		First(&info).Error
	return &info, err
}

func (m *NftMallOrders) ListByUserIdAndStatus(column, userId string, status, offset, size int) (*[]NftMallOrders, error) {
	var list []NftMallOrders
	db := m.Table().Preload("Goods").Where(fmt.Sprintf("%s=?", column), userId)
	if status > 0 {
		db.Where("status=?", status)
	}
	err := db.Order("dt_create DESC").Offset(offset).Limit(size).Find(&list).Error
	return &list, err
}
