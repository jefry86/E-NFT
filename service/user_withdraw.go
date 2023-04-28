package service

import (
	"fmt"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/utils"
	"nft_platform/vm/response"
	"time"
)

var userWithdrawMode model.NftUserWithdraw

type UserWithdraw struct {
}

func (u *UserWithdraw) List(userId string, pageNo, pageSize int) (*response.UserWithdrawListRes, error) {
	offset, pageSize := utils.PageOffset(pageNo, pageSize)
	count, err := userWithdrawMode.CountByUserId(userId)
	if err != nil {
		global.SLogger.Error("ListByUserId error:%s", err)
		return nil, fmt.Errorf("系统错误！")
	}

	if count == 0 {
		return nil, nil
	}

	list, err := userWithdrawMode.ListByUserId(userId, offset, pageSize)
	if err != nil {
		global.SLogger.Error("ListByUserId error:%s", err)
		return nil, fmt.Errorf("系统错误！")
	}
	res := make([]response.UserWithdrawInfo, 0)

	for _, withdraw := range *list {
		res = append(res, response.UserWithdrawInfo{
			Amount:   withdraw.Amount,
			Bank:     withdraw.Bank,
			Status:   withdraw.Status,
			DateTime: time.Unix(int64(withdraw.DtCreate), 0),
		})
	}

	return &response.UserWithdrawListRes{
		List: res,
		Page: response.Page{
			Total: count,
			Size:  pageSize,
		},
	}, nil
}

func (u *UserWithdraw) Add(userId string, bankId, amount int) (bool, error) {
	userInfo, err := usersModel.FindByUserId(userId)
	if err != nil {
		global.SLogger.Error("Find User error:%s", err)
		return false, fmt.Errorf("系统错误！")
	}
	if userInfo.Amount < uint(amount) {
		return false, fmt.Errorf("余额不足！")
	}

	bank, err := userBankModel.FindById(bankId)
	if err != nil {
		global.SLogger.Error("Find Bank error:%s", err)
		return false, fmt.Errorf("系统错误！")
	}
	if bank == nil {
		return false, fmt.Errorf("银行信息有误！")
	}

	data := model.NftUserWithdraw{
		UserID:      userId,
		Amount:      uint(amount),
		Bank:        bank.Bank,
		BankName:    bank.BankName,
		BankAccount: bank.BankAccount,
		BankAddr:    bank.BankAddress,
		Status:      1,
	}

	err = userWithdrawMode.Add(data)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
