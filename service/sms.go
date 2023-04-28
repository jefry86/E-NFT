package service

import (
	"fmt"
	"math"
	"math/rand"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/utils"
	"nft_platform/utils/tcloud"
	"strconv"
	"time"
)

var sms tcloud.Sms

type Sms struct {
}

// SendCode 发送短信验证码， t: 1 登录/注册 2 修改手机号码 3 短信验证码
func (s *Sms) SendCode(userId, mobile, ip string, t int) error {
	redisKey := fmt.Sprintf("%s%s%d", global.RedisSmsCodePrefix, mobile, t)
	if t != 2 && !utils.ValidatorMobile(mobile) {
		return fmt.Errorf("手机号码错误！")
	}

	if err := s.chkLimitMobile(mobile); err != nil {
		return err
	}

	if err := s.chkLimitIP(ip); err != nil {
		return err
	}

	var err error
	code := s.getCode(6)

	switch t {
	case 1:
		err = s.loginCode(mobile, code)
	case 2:
		err = s.SmsCodeByChkUser(userId, code)
	case 3:
		err = sms.Code(mobile, code)
	default:
		return fmt.Errorf("短信类型不支持！")
	}
	if err != nil {
		return err
	}
	return utils.RedisSetString(redisKey, code, time.Duration(1)*time.Minute)
}

func (s *Sms) ChkCode(mobile, code string, t int) (bool, error) {
	redisKey := fmt.Sprintf("%s%s%d", global.RedisSmsCodePrefix, mobile, t)
	redisCode, err := utils.RedisGetString(redisKey)
	if err != nil {
		return false, fmt.Errorf("系统错误！")
	}
	if redisCode != code {
		return false, nil
	}
	return true, nil
}

// mobile频率限制
func (s *Sms) chkLimitMobile(mobile string) error {
	redisKey := fmt.Sprintf("%s%s", global.RedisSmsLimitMobilePrefix, mobile)
	num, err := utils.RedisGetString(redisKey)
	if err != nil {
		global.SLogger.Errorf("redis获取:%s err:%s", redisKey, err)
		return fmt.Errorf("系统错误！")
	}
	if utils.AtoInt(num) > global.Conf.Server.SmsMobileLimit {
		global.SLogger.Errorf("mobile:%s 短信量超过了:%d 次", mobile, global.Conf.Server.SmsMobileLimit)
		return fmt.Errorf("您的短信发送太过频繁了！")
	}
	return nil
}

// IP频率限制
func (s *Sms) chkLimitIP(ip string) error {
	redisKey := fmt.Sprintf("%s%s", global.RedisSmsLimitIPPrefix, ip)
	num, err := utils.RedisGetString(redisKey)
	if err != nil {
		global.SLogger.Errorf("redis获取:%s err:%s", redisKey, err)
		return fmt.Errorf("系统错误！")
	}
	if utils.AtoInt(num) > global.Conf.Server.SmsIPLimit {
		global.SLogger.Errorf("ip:%s 短信量超过了:%d 次", ip, global.Conf.Server.SmsMobileLimit)
		return fmt.Errorf("您的短信发送太过频繁了！")
	}
	return nil
}

func (s *Sms) setLimit(mobile, ip string) error {
	redisLimitMobileKey := fmt.Sprintf("%s%s", global.RedisSmsLimitMobilePrefix, mobile)
	redisLimitIPKey := fmt.Sprintf("%s%s", global.RedisSmsLimitIPPrefix, ip)

	if err := utils.RedisInc(redisLimitMobileKey); err != nil {
		global.SLogger.Errorf("redis inc:%s err:%s", redisLimitMobileKey, err)
		return fmt.Errorf("系统错误！")
	}

	if err := utils.RedisInc(redisLimitIPKey); err != nil {
		global.SLogger.Errorf("redis inc:%s err:%s", redisLimitIPKey, err)
		return fmt.Errorf("系统错误！")
	}

	return nil
}

func (s *Sms) loginCode(mobile, code string) error {
	return sms.LoginCode(mobile, code)
}

func (s *Sms) getCode(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	max := math.Pow(10, float64(length))
	return strconv.Itoa(r.Intn(int(max)))
}

func (s *Sms) SmsCodeByChkUser(userId, code string) error {
	var userModel model.NftUsers
	users, err := userModel.FindByUserId(userId)
	if err != nil {
		global.SLogger.Errorf("通过user id获取用户信息出错,err:%s", err)
		return fmt.Errorf("系统错误！")
	}
	if users == nil {
		return fmt.Errorf("用户不存在！")
	}
	return sms.Code(users.Mobile, code)
}
