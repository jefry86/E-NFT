package service

import (
	"fmt"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/utils"
	"nft_platform/vm/response"
	"time"
)

var userBalance model.NftUserBalanceLog

type UserBalance struct {
}

func (u *UserBalance) List(userId string, pageNo, pageSize int) (*response.UserBalanceListRes, error) {
	offset, pageSize := utils.PageOffset(pageNo, pageSize)

	count, err := userBalance.CountByUserId(userId)
	if err != nil {
		global.SLogger.Errorf("count error:%s", err)
		return nil, fmt.Errorf("系统错误！")
	}

	if count == 0 {
		return nil, nil
	}

	list, err := userBalance.ListByUserId(userId, offset, pageSize)
	if err != nil {
		global.SLogger.Errorf("List error:%s", err)
		return nil, fmt.Errorf("系统错误！")
	}

	res := make([]response.UserBalanceList, 0)
	for _, log := range *list {
		res = append(res, response.UserBalanceList{
			Amount:   log.Amount,
			Balance:  log.Balance,
			Info:     log.Info,
			Type:     log.Type,
			DateTime: time.Time{},
		})
	}

	return &response.UserBalanceListRes{
		List: res,
		Page: response.Page{
			Total: count,
			Size:  pageSize,
		},
	}, nil
}

func (u *UserBalance) Add() {

}
