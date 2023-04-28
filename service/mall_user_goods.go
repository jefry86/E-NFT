package service

import (
	"fmt"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/vm/response"
	"time"
)

var mallUserGoodsModel model.NftMallUserGoods

type MallUserGoods struct {
}

func (u *MallUserGoods) Goods(userId string, t, pageNo, pageSize int) (*response.MallUserGoodsListRes, error) {
	count, err := mallUserGoodsModel.Count(userId, t)
	if err != nil {
		global.SLogger.Errorf("get mall user goods count for db err:%s", err)
		return nil, fmt.Errorf("系统错误！")
	} else if count == 0 {
		return nil, nil
	}

	list, err := mallUserGoodsModel.List(userId, t, pageNo, pageSize)
	if err != nil {
		global.SLogger.Errorf("get mall user goods list for db err:%s", err)
		return nil, fmt.Errorf("系统错误！")
	}

	if list == nil {
		return nil, nil
	}
	res := make([]response.MallUserGoodsList, 0)
	for _, goods := range *list {
		res = append(res, response.MallUserGoodsList{
			Id:           goods.ID,
			GoodsName:    goods.GoodsName,
			GoodsId:      goods.GoodsID,
			GoodsImage:   goods.Goods.Image,
			GoodsType:    goods.Goods.Type,
			GoodsLabel:   goods.Goods.Label,
			GoodsNo:      goods.GoodsNo,
			PlatformId:   goods.Goods.PlatformID,
			PlatformName: "-",
			PlatformLogo: "-",
			DateTime:     time.Unix(int64(goods.DtCreate), 0).Format(time.DateOnly),
		})
	}
	return &response.MallUserGoodsListRes{
		List: res,
		Page: response.Page{
			Total: count,
			Size:  pageSize,
		},
	}, nil
}

func (u *MallUserGoods) GoodsInfo(userId string, id int) (*response.MallUserGoodsInfo, error) {
	info, err := mallUserGoodsModel.FindByIdAndUserId(uint(id), userId)
	if err != nil {
		global.SLogger.Errorf("get user goods info err:%s", err.Error())
		return nil, fmt.Errorf("系统错误！")
	}
	if info == nil {
		return nil, nil
	}
	return &response.MallUserGoodsInfo{
		Id:           info.ID,
		GoodsId:      info.GoodsID,
		GoodsName:    info.GoodsName,
		GoodsImage:   info.Goods.Image,
		GoodsType:    info.Goods.Type,
		GoodsLabel:   info.Goods.Label,
		GoodsNo:      info.GoodsNo,
		GoodsDetail:  info.Goods.Detail,
		GoodsHash:    info.GoodsHash,
		PlatformId:   info.Goods.PlatformID,
		PlatformName: "-",
		PlatformLogo: "-",
		DateTime:     time.Unix(int64(info.DtCreate), 0).Format(time.DateOnly),
	}, nil
}
