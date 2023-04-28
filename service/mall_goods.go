package service

import (
	"fmt"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/vm/response"
	"time"
)

var mallGoodsModel model.NftMallGoods

type MallGoods struct {
}

func (m *MallGoods) List(sort string, t, pageNo, PageSize int) (*response.MallGoodsRes, error) {
	offset := (pageNo - 1) * PageSize
	status := []int{
		1, 2,
	}
	var sortCloumn string
	switch sort {
	case "price":
		sortCloumn = "price"
	default:
		sortCloumn = "dt_create"
	}

	s := "DESC"
	count, err := mallGoodsModel.Count(status)
	if err != nil {
		global.SLogger.Error("get mall goods count err:%s", err.Error())
		return nil, fmt.Errorf("系统错误！")
	} else if count == 0 {
		return nil, nil
	}

	list, err := mallGoodsModel.List(status, sortCloumn, s, offset, PageSize)
	if err != nil {
		global.SLogger.Error("get mall goods list err:%s", err.Error())
		return nil, fmt.Errorf("系统错误！")
	}

	res := make([]response.MallGoodsList, 0)
	for _, goods := range *list {
		res = append(res, response.MallGoodsList{
			Id:           goods.ID,
			Name:         goods.Name,
			Image:        goods.Image,
			Type:         goods.Type,
			Label:        goods.Label,
			No:           goods.No,
			PlatformId:   goods.PlatformID,
			PlatformName: "-",
			PlatformLogo: "-",
			DateTime:     time.Unix(int64(goods.DtCreate), 0).Format(time.DateOnly),
		})
	}
	return &response.MallGoodsRes{
		List: res,
		Page: response.Page{
			Total: count,
			Size:  PageSize,
		},
	}, nil
}
