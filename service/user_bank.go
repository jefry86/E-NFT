package service

import (
	"fmt"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/vm/response"
)

var userBankModel model.NftUserBank

type UserBank struct {
}

func (u *UserBank) List(userId string) (*[]response.UserBankInfo, error) {
	list, err := userBankModel.ListByUserId(userId)
	if err != nil {
		global.SLogger.Errorf("list blank err:%s", err.Error())
		return nil, fmt.Errorf("系统错误！")
	}

	res := make([]response.UserBankInfo, 0)
	for _, bank := range *list {
		res = append(res, response.UserBankInfo{
			Bank:        bank.Bank,
			BankName:    bank.BankName[:1],
			BankAccount: bank.BankAccount[:4] + "****" + bank.BankAccount[len(bank.BankAccount)-4:],
			BankAddr:    bank.BankAddress[:5],
		})
	}
	return &res, nil

}

func (u *UserBank) Add(data response.UserBankInfo, userId string) (bool, error) {
	bankData := model.NftUserBank{
		UserID:      userId,
		Bank:        data.Bank,
		BankName:    data.BankName,
		BankAccount: data.BankAccount,
		BankAddress: data.BankAddr,
		Status:      1,
	}

	err := userBankModel.Add(bankData)
	if err != nil {
		global.SLogger.Errorf("add blank  err:%s", err.Error())
		return false, fmt.Errorf("系统错误！")
	}
	return true, nil
}

func (u *UserBank) Delete(id int, userId string) (bool, error) {
	err := userBankModel.Delete(id, userId)
	if err != nil {
		global.SLogger.Errorf("add blank  err:%s", err.Error())
		return false, fmt.Errorf("系统错误！")
	}
	return true, nil
}
