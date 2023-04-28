package service

import (
	"fmt"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/utils"
	"nft_platform/vm/response"
	"time"
)

var mallOrderModel model.NftMallOrders
var userServer Users

type MallOrder struct {
}

// Apply 创建订单
func (m *MallOrder) Apply(userId string, goodsId, num uint) (string, error) {
	goodsInfo, err := mallGoodsModel.FindById(goodsId)
	if err != nil {
		global.SLogger.Error("get mall goods err:%s", err)
		return "", fmt.Errorf("系统错误！")
	} else if goodsInfo == nil {
		return "", fmt.Errorf("藏品不存在！")
	}
	orderId := m.generateOrderId()
	payEndTime := time.Now().Add(time.Minute * time.Duration(global.Conf.Server.PayExpirationTime)).Unix()
	data := model.NftMallOrders{
		OrderID:        orderId,
		GoodsID:        goodsId,
		MallPlatformID: goodsInfo.PlatformID,
		FromUserID:     goodsInfo.UserId,
		UserID:         userId,
		Num:            num,
		Amount:         goodsInfo.Price * num,
		PayChannelID:   0,
		PayEndtime:     uint(payEndTime),
		Status:         0,
	}
	err = mallOrderModel.Add(data)
	if err != nil {
		global.SLogger.Error("add mall order err:%s", err)
		return "", fmt.Errorf("系统错误！")
	}
	return orderId, nil
}

func (m *MallOrder) Info(userId, orderId string) (*response.MallOrderInfo, error) {
	info, err := mallOrderModel.FindByOrderIdAndUserId(orderId, userId)
	if err != nil {
		global.SLogger.Error("find order order err:%s", err)
		return nil, fmt.Errorf("系统错误！")
	}

	return &response.MallOrderInfo{
		OrderId:    info.OrderID,
		Amount:     info.Amount,
		Num:        info.Num,
		GoodsId:    info.GoodsID,
		GoodsName:  info.Goods.Name,
		GoodsImage: info.Goods.Image,
		GoodsType:  info.Goods.Type,
		GoodsLabel: info.Goods.Label,
		GoodsNo:    info.Goods.No,
		DateTime:   time.Unix(int64(info.DtCreate), 0),
		Status:     info.Status,
		PayEndTime: info.PayEndtime - uint(time.Now().Unix()),
	}, nil
}

func (m *MallOrder) Notify(orderId, payTime string, payChannel, amount uint) (bool, error) {
	orderInfo, err := mallOrderModel.FindByOrderId(orderId)
	if err != nil {
		global.SLogger.Error("get order err:%s", err)
		return false, fmt.Errorf("系统错误！")
	} else if orderInfo.Status != 1 {
		return true, nil
	} else if orderInfo.Amount < amount {
		return false, fmt.Errorf("订单金额有误！")
	}

	dataTime, err := time.Parse(time.DateTime, payTime)
	if err != nil {
		global.SLogger.Error("time parse err:%s", err)
		return false, fmt.Errorf("系统错误！")
	}
	updateOrderData := model.NftMallOrders{
		Status:       2,
		PayChannelID: payChannel,
		PayTime:      uint(dataTime.Unix()),
	}
	err = mallOrderModel.UpdateByOrderId(updateOrderData, orderId)

	if err != nil {
		global.SLogger.Error("change order err:%s", err)
		return false, fmt.Errorf("系统错误！")
	}

	_, err = m.Deliver(orderId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Deliver 发货
func (m *MallOrder) Deliver(orderId string) (bool, error) {
	orderInfo, err := mallOrderModel.FindByOrderId(orderId)
	if err != nil {
		global.SLogger.Error("get order err:%s", err)
		return false, fmt.Errorf("系统错误！")
	} else if orderInfo.Status == 3 {
		return true, nil
	} else if orderInfo.Status != 2 {
		return false, fmt.Errorf("订单未支付！")
	}

	data := model.NftMallUserGoods{
		GoodsID:        orderInfo.GoodsID,
		GoodsName:      orderInfo.Goods.Name,
		GoodsNo:        orderInfo.Goods.No,
		FromUserID:     orderInfo.Goods.UserId,
		MallPlatformID: orderInfo.Goods.PlatformID,
		UserID:         orderInfo.UserID,
		GoodsHash:      orderInfo.Goods.Hash,
		Status:         1,
	}

	err = mallUserGoodsModel.Add(data)
	if err != nil {
		global.SLogger.Error("user goods add err:%s", err)
		return false, fmt.Errorf("系统错误！")
	}

	// 给用户调整余额
	_, err = userServer.Balance(orderInfo.UserID, orderInfo.Goods.Name,
		orderInfo.Goods.Image, fmt.Sprintf("转卖藏品：%s", orderInfo.Goods.Name),
		int(orderInfo.Amount), orderInfo.GoodsID, 1)
	if err != nil {
		global.SLogger.Error("user balance err:%s", err)
		return false, fmt.Errorf("系统错误！")
	}

	return true, nil
}

func (m *MallOrder) ListForBuy(userId string, status, pageNo, pageSize int) (*[]response.MallOrderList, error) {
	offset, pageSize := utils.PageOffset(pageNo, pageSize)
	list, err := mallOrderModel.ListByUserIdAndStatus("user_id", userId, status, offset, pageSize)
	if err != nil {
		global.SLogger.Errorf("get list by db err:%s", err)
		return nil, fmt.Errorf("系统错误！")
	}

	res := make([]response.MallOrderList, 0)

	for _, order := range *list {
		res = append(res, response.MallOrderList{
			OrderId:    order.OrderID,
			Amount:     order.Amount,
			Num:        order.Num,
			GoodsId:    order.GoodsID,
			GoodsName:  order.Goods.Name,
			GoodsImage: order.Goods.Image,
			GoodsType:  order.Goods.Type,
			GoodsLabel: order.Goods.Label,
			GoodsNo:    order.Goods.No,
			DateTime:   time.Unix(int64(order.DtCreate), 0),
			Status:     order.Status,
		})
	}

	return &res, nil
}

func (m *MallOrder) ListForSale(userId string, status, pageNo, pageSize int) (*[]response.MallOrderList, error) {
	offset, pageSize := utils.PageOffset(pageNo, pageSize)
	list, err := mallOrderModel.ListByUserIdAndStatus("from_user_id", userId, status, offset, pageSize)
	if err != nil {
		global.SLogger.Errorf("get list by db err:%s", err)
		return nil, fmt.Errorf("系统错误！")
	}
	res := make([]response.MallOrderList, 0)

	for _, order := range *list {
		res = append(res, response.MallOrderList{
			OrderId:    order.OrderID,
			Amount:     order.Amount,
			Num:        order.Num,
			GoodsId:    order.GoodsID,
			GoodsName:  order.Goods.Name,
			GoodsImage: order.Goods.Image,
			GoodsType:  order.Goods.Type,
			GoodsLabel: order.Goods.Label,
			GoodsNo:    order.Goods.No,
			DateTime:   time.Unix(int64(order.DtCreate), 0),
			Status:     order.Status,
		})
	}
	return &res, nil
}

func (m *MallOrder) generateOrderId() string {
	return fmt.Sprintf("%s%d", time.Now().Format("20160102150405"), time.Now().Nanosecond())
}
