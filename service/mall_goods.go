package service

import (
	"fmt"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/utils"
	"nft_platform/vm/request"
	"nft_platform/vm/response"
	"strings"
	"time"
)

var mallGoodsModel model.NftMallGoods

type MallGoods struct {
}

func (m *MallGoods) List(keyword, platformId, sort string, goodsType, pageNo, pageSize int) (*response.MallGoodsListRes, error) {
	offset, pageSize := utils.PageOffset(pageNo, pageSize)
	status := []int{
		1, 4,
	}

	var sortColumn, sortType string

	if sort != "" {
		sorts := strings.Split(sort, "_")
		sortColumn = sorts[0]
		sortType = sorts[1]
	}
	if (sortColumn != "price" && sortColumn != "time") || sortColumn == "time" {
		sortColumn = "dt_create"
	}

	if sortType != "desc" && sortType != "asc" {
		sortType = "desc"
	}
	ids := strings.Split(platformId, ",")
	platformIds := make([]int, 0)
	for _, id := range ids {
		if utils.AtoInt(id) > 0 {
			platformIds = append(platformIds, utils.AtoInt(id))
		}
	}
	count, err := mallGoodsModel.Count(keyword, goodsType, platformIds, status)
	if err != nil {
		global.SLogger.Error("get mall goods count err:%s", err.Error())
		return nil, fmt.Errorf("系统错误！")
	} else if count == 0 {
		return &response.MallGoodsListRes{
			List: nil,
			Page: response.Page{
				Total: count,
				Size:  pageSize,
			},
		}, nil
	}

	list, err := mallGoodsModel.List(keyword, goodsType, platformIds, status, sortColumn, sortType, offset, pageSize)
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
			PlatformName: goods.Platform.Name,
			PlatformLogo: goods.Platform.Logo,
			Price:        goods.Price,
			DateTime:     time.Unix(int64(goods.DtCreate), 0).Format(time.DateOnly),
			Status:       goods.Status,
		})
	}
	return &response.MallGoodsListRes{
		List: res,
		Page: response.Page{
			Total: count,
			Size:  pageSize,
		},
	}, nil
}

func (m *MallGoods) Info(id int, userId string) (*response.MallGoodsInfo, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id 有误")
	}
	info, err := mallGoodsModel.Info(id, userId)
	if err != nil {
		return nil, err
	}
	if info != nil {
		return &response.MallGoodsInfo{
			Id:           info.ID,
			Name:         info.Name,
			Image:        info.Image,
			Type:         info.Type,
			Label:        info.Label,
			No:           info.No,
			Detail:       info.Detail,
			Hash:         info.Hash,
			PlatformId:   info.PlatformID,
			PlatformName: info.Platform.Name,
			PlatformLogo: info.Platform.Logo,
			DateTime:     time.Unix(int64(info.DtCreate), 0).Format(time.DateOnly),
		}, nil
	}
	return nil, nil

}

func (m *MallGoods) ListByUserId(userId string, status, pageNo, pageSize int) (*response.MallGoodsListRes, error) {
	offset, pageSize := utils.PageOffset(pageNo, pageSize)
	count, err := mallGoodsModel.CountByUserAndStatus(userId, status)
	if err != nil {
		global.SLogger.Error("get mall goods count err:%s", err.Error())
		return nil, fmt.Errorf("系统错误！")
	} else if count == 0 {
		return nil, nil
	}

	list, err := mallGoodsModel.ListByUserId(userId, status, offset, pageSize)
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
	return &response.MallGoodsListRes{
		List: res,
		Page: response.Page{
			Total: count,
			Size:  pageSize,
		},
	}, nil
}

func (m *MallGoods) Add(goods request.MallGoodsAdd, userId string) (bool, error) {
	data := model.NftMallGoods{
		Name:          goods.Name,
		Image:         goods.Image,
		Detail:        goods.Detail,
		Price:         goods.Price,
		OriginalPrice: goods.OriginalPrice,
		PlatformID:    goods.PlatformID,
		No:            goods.No,
		Source:        goods.Source,
		SourceType:    goods.SourceType,
		Hash:          goods.Hash,
		UserId:        userId,
		Status:        1,
		Type:          1,
	}

	err := mallGoodsModel.Add(data)
	if err != nil {
		global.SLogger.Errorf("create db err:%s", err)
		return false, err
	}
	return true, nil
}
