package service

import (
	"fmt"
	"github.com/google/uuid"
	idvalidator "github.com/guanguans/id-validator"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/utils"
	"nft_platform/utils/tcloud"
	"nft_platform/vm/response"
	"time"
)

var smsService Sms
var usersModel model.NftUsers
var userBalanceLogModel model.NftUserBalanceLog

type Users struct {
}

func (u *Users) LoginByMobile(mobile, code string) (*response.Login, error) {
	ok, err := smsService.ChkCode(mobile, code, 1)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("验证码错误，请重新输入！")
	}
	//用户不存在、用户存在
	userInfo, err := usersModel.FindByMobile(mobile)
	if err != nil {
		return nil, err
	}
	token := u.generateToken()
	redisKey := fmt.Sprintf("%s%s", global.RedisUserTokenPrefix, token)
	var userId, nickname string
	//用户不存在，注册
	if userInfo == nil {
		userId = utils.RandString(12)
		nickname = fmt.Sprintf("nf-%s", mobile[8:])
		userInfo := model.NftUsers{
			UserID:   userId,
			Username: mobile,
			Password: "",
			Nickname: nickname,
			Mobile:   mobile,
			Avatar:   global.Conf.Server.DefaultAvatar,
			Token:    token,
		}
		if err := usersModel.Add(userInfo); err != nil {
			global.SLogger.Error("add user err:%s", err)
			return nil, fmt.Errorf("系统异常，请联系客服！")
		}
		err := utils.RedisSetString(redisKey, userId, time.Duration(180*24)*time.Hour)
		if err != nil {
			global.SLogger.Error("set token to redis err:%s", err)
			return nil, fmt.Errorf("系统异常，请联系客服！")
		}

	} else {
		//用户存在
		data := model.NftUsers{Token: token}
		if err := usersModel.UpdateByMobile(data, mobile); err != nil {
			global.SLogger.Error("add user err:%s", err)
			return nil, fmt.Errorf("系统异常，请联系客服！")
		}
		userId = userInfo.UserID
		nickname = userInfo.Nickname
	}

	if err := utils.RedisSetString(redisKey, userId, time.Duration(180*24)*time.Hour); err != nil {
		global.SLogger.Error("set token to redis err:%s", err)
		return nil, fmt.Errorf("系统异常，请联系客服！")
	}
	return &response.Login{
		UserId:   userId,
		Token:    token,
		Nickname: nickname,
		Mobile:   fmt.Sprintf("%s****%s", mobile[0:3], mobile[8:]),
	}, nil
}

func (u *Users) generateToken() string {
	return uuid.New().String()
}

// Auth 实名认证
func (u *Users) Auth(userId, realName, card string) error {
	cardInfo, err := idvalidator.GetInfo(card, true)
	if err != nil {
		return fmt.Errorf("身份号码错误")
	}
	birthday := cardInfo.Birthday
	userInfo, err := usersModel.FindByUserId(userId)
	if err != nil {
		global.SLogger.Error("FindByUserId err:%s", err)
		return fmt.Errorf("系统错误！")
	}
	if userInfo == nil {
		return fmt.Errorf("用户不存在")
	}

	var real tcloud.RealName
	ok, err := real.Check(realName, card)
	if err != nil {
		global.SLogger.Error("real name check err:%s", err)
		return fmt.Errorf("系统错误！")
	}
	if !ok {
		return fmt.Errorf("身份号码有误")
	}
	data := model.NftUsers{
		RealName: realName,
		CardNo:   card,
		Birthday: birthday,
	}

	if err = usersModel.UpdateByMobile(data, userId); err != nil {
		global.SLogger.Error("UpdateByMobile err:%s", err)
		return fmt.Errorf("系统错误！")
	}
	return nil
}

func (u *Users) GetUserByToken(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("token为空！")
	}
	redisKey := fmt.Sprintf("%s%s", global.RedisUserTokenPrefix)
	userId, err := utils.RedisGetString(redisKey)
	if err != nil {
		global.SLogger.Error("get user id by redis,err:%s", err)
		return "", fmt.Errorf("系统错误！")
	}
	if userId != "" {
		return userId, nil
	}

	users, err := usersModel.FindByToken(token)
	if err != nil {
		global.SLogger.Error("get user id by db,err:%s", err)
		return "", fmt.Errorf("系统错误！")
	}
	if users == nil {
		return "", fmt.Errorf("token错误！")
	}
	return users.UserID, nil
}

func (u *Users) Logout(token string) error {
	userId, err := u.GetUserByToken(token)
	if err != nil {
		global.SLogger.Error("get userId by db ,err:%s", err)
		return err
	}

	redisKey := fmt.Sprintf("%s%s", global.RedisUserTokenPrefix, token)
	if err := utils.RedisUnlink(redisKey); err != nil {
		global.SLogger.Error("redis unlink token ,err:%s", err)
		return fmt.Errorf("系统错误！")
	}
	data := model.NftUsers{Token: ""}
	return usersModel.UpdateByUserId(data, userId)

}

// VerifyOldMobile 验证旧手机
func (u *Users) VerifyOldMobile(userId, code string) (bool, error) {
	users, err := usersModel.FindByUserId(userId)
	if err != nil {
		global.SLogger.Error("get user err:%s", err)
		return false, err
	}
	ok, err := smsService.ChkCode(users.Mobile, code, 2)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, fmt.Errorf("验证码错误！")
	}

	redisKey := fmt.Sprintf("%s%s", global.RedisVerifyOldMobilePrefix, userId)
	if err = utils.RedisSetString(redisKey, userId, time.Duration(24)*time.Hour); err != nil {
		return false, err
	}
	return true, nil
}

func (u *Users) ChangeMobile(userId, mobile, code string) (bool, error) {
	ok, err := smsService.ChkCode(mobile, code, 3)
	if err != nil {
		global.SLogger.Error("check code err:%s", err)
		return false, err
	}

	if !ok {
		return false, fmt.Errorf("验证码错误！")
	}

	redisKey := fmt.Sprintf("%s%s", global.RedisVerifyOldMobilePrefix, userId)
	val, err := utils.RedisGetString(redisKey)
	if err != nil {
		global.SLogger.Errorf("get redis err:%s", err)
		return false, fmt.Errorf("系统错误！")
	}
	if val == "" {
		return false, fmt.Errorf("错误的操作！")
	}

	data := model.NftUsers{Mobile: mobile}
	err = usersModel.UpdateByUserId(data, userId)
	if err != nil {
		global.SLogger.Errorf("update user db err:%s", err)
		return false, fmt.Errorf("系统错误！")
	}
	return true, nil
}

func (u *Users) Info(userId string) (*response.UserInfo, error) {
	users, err := usersModel.FindByUserId(userId)
	if err != nil {
		global.SLogger.Error("get user info err:%s", err)
		return nil, err
	}

	return &response.UserInfo{
		UserId:     userId,
		Nickname:   users.Nickname,
		Mobile:     fmt.Sprintf("%s****%s", users.Mobile[:3], users.Mobile[8:]),
		Avatar:     users.Avatar,
		WalletAddr: users.WalletAddr,
	}, nil
}

func (u *Users) Balance(userId, goodsName, goodsImage, info string, amount int, goodsId, t uint) (bool, error) {
	users, err := usersModel.FindByUserId(userId)
	if err != nil {
		global.SLogger.Error("get user info err:%s", err)
		return false, fmt.Errorf("系统错误！")
	}
	var balance uint
	if t == 1 {
		balance = uint(int(users.Amount) + amount)
	} else {
		amount = -1 * amount
		balance = uint(int(users.Amount) - amount)
	}

	data := model.NftUserBalanceLog{
		UserID:     userId,
		Amount:     amount,
		Balance:    balance,
		GoodsID:    goodsId,
		GoodsName:  goodsName,
		GoodsImage: goodsImage,
		Info:       info,
		Type:       uint8(t),
	}

	if err = userBalanceLogModel.Add(data); err != nil {
		global.SLogger.Error("user balance log err:%s", err)
		return false, fmt.Errorf("系统错误！")
	}

	userData := model.NftUsers{
		Amount: uint(int(users.Amount) + amount),
	}
	if t == 1 {
		// 转卖藏品
		userData.Earnings = uint(int(userData.Earnings) + amount)
	} else if t == 3 {
		// 提现
		userData.FreezeAmount = uint(int(userData.FreezeAmount) + amount)
	}

	err = usersModel.UpdateByUserId(userData, userId)
	if err = userBalanceLogModel.Add(data); err != nil {
		global.SLogger.Error("user update err:%s", err)
		return false, fmt.Errorf("系统错误！")
	}

	return true, nil
}
