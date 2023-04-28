package service

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"net/http"
	"net/url"
	"nft_platform/global"
	"nft_platform/model"
	"nft_platform/vm/response"
	"sort"
	"strconv"
	"strings"
	"time"
)

var mallPlatformMode model.NftMallPlatform

type Platform struct {
}

// Register 用户登记
func (p *Platform) Register(platformId int, nickName, mobile, name, cardNo string) (*response.PlatformUser, error) {
	info, err := p.getPlatformInfo(platformId)
	if err != nil {
		global.SLogger.Errorf("get Platform Info err:%s", err)
		return nil, err
	}
	var data url.Values
	timeline := strconv.FormatInt(time.Now().Unix(), 10)
	data.Add("nick_name", nickName)
	data.Add("mobile", mobile)
	data.Add("name", name)
	data.Add("card_no", cardNo)
	data.Add("timeline", timeline)
	authorization := p.toSign(data, info.AppSecret)
	data.Del("timeline")

	var header http.Header
	header.Add("timeline", timeline)
	header.Add("authorization", authorization)
	res, err := p.call(info.ApiUser, http.MethodPost, data, header)
	if err != nil {
		return nil, err
	}
	global.SLogger.Infof("api user res:%s", res)

	var userInfo response.PlatformUser

	err = json.Unmarshal(res, &userInfo)
	if err != nil {
		global.SLogger.Errorf("json Platform Info err:%s", err)
		return nil, err
	}

	return &userInfo, nil
}

// Transfer 售卖通知
func (p *Platform) Transfer(platformId int, transferId, goodsHash, fromWalletHash, toWalletHash, name, cardNo string) (*response.PlatformTransfer, error) {
	info, err := p.getPlatformInfo(platformId)
	if err != nil {
		global.SLogger.Errorf("get Platform Info err:%s", err)
		return nil, err
	}
	var data url.Values
	timeline := strconv.FormatInt(time.Now().Unix(), 10)
	data.Add("transfer_id", transferId)
	data.Add("goods_hash", goodsHash)
	data.Add("from_wallet_hash", fromWalletHash)
	data.Add("to_wallet_hash", toWalletHash)
	data.Add("name", name)
	data.Add("card_no", cardNo)
	data.Add("timeline", timeline)
	authorization := p.toSign(data, info.AppSecret)
	data.Del("timeline")

	var header http.Header
	header.Add("timeline", timeline)
	header.Add("authorization", authorization)
	res, err := p.call(info.ApiUser, http.MethodPut, data, header)
	if err != nil {
		return nil, err
	}
	global.SLogger.Infof("api transfer res:%s", res)
	var result response.PlatformTransfer
	err = json.Unmarshal(res, &result)
	if err != nil {
		global.SLogger.Errorf("json Platform Info err:%s", err)
		return nil, err
	}

	return &result, nil
}

// Sales 上架查询
func (p *Platform) Sales(platformId int, userId, walletHash, goodsId, goodsHash, name, cardNo string) (*response.PlatformSales, error) {
	info, err := p.getPlatformInfo(platformId)
	if err != nil {
		global.SLogger.Errorf("get Platform Info err:%s", err)
		return nil, err
	}
	var data url.Values
	timeline := strconv.FormatInt(time.Now().Unix(), 10)
	data.Add("user_id", userId)
	data.Add("wallet_hash", walletHash)
	data.Add("goods_id", goodsId)
	data.Add("goods_hash", goodsHash)
	data.Add("name", name)
	data.Add("card_no", cardNo)
	data.Add("timeline", timeline)
	authorization := p.toSign(data, info.AppSecret)
	data.Del("timeline")

	var header http.Header
	header.Add("timeline", timeline)
	header.Add("authorization", authorization)
	res, err := p.call(info.ApiUser, http.MethodGet, data, header)
	if err != nil {
		return nil, err
	}
	global.SLogger.Infof("api sales res:%s", res)

	var result response.PlatformSales
	err = json.Unmarshal(res, &result)
	if err != nil {
		global.SLogger.Errorf("json Platform Info err:%s", err)
		return nil, err
	}
	return &result, nil
}

// GoodsList 藏品列表
func (p *Platform) GoodsList(walletHash, name, cardNo string, platformId, pageNo, pageSize int) (*response.PlatformGoodsListRes, error) {
	info, err := p.getPlatformInfo(platformId)
	if err != nil {
		global.SLogger.Errorf("get Platform Info err:%s", err)
		return nil, err
	}
	var data url.Values
	timeline := strconv.FormatInt(time.Now().Unix(), 10)
	data.Add("wallet_hash", walletHash)
	data.Add("name", name)
	data.Add("card_no", cardNo)
	data.Add("page_no", strconv.Itoa(pageNo))
	data.Add("page_size", strconv.Itoa(pageSize))
	data.Add("timeline", timeline)
	authorization := p.toSign(data, info.AppSecret)
	data.Del("timeline")

	var header http.Header
	header.Add("timeline", timeline)
	header.Add("authorization", authorization)
	res, err := p.call(info.ApiUser, http.MethodGet, data, header)
	if err != nil {
		return nil, err
	}
	global.SLogger.Infof("api user res:%s", res)

	var result response.PlatformGoodsListRes
	err = json.Unmarshal(res, &result)
	if err != nil {
		global.SLogger.Errorf("json Platform Info err:%s", err)
		return nil, err
	}
	return &result, nil
}

func (p *Platform) getPlatformInfo(platformId int) (*model.NftMallPlatform, error) {
	info, err := mallPlatformMode.FindById(platformId)
	if err != nil {
		global.SLogger.Errorf("find app key err:%s", err.Error())
		return nil, err
	}

	if info.Status != 1 {
		return nil, fmt.Errorf("平台不可用！")
	} else if info.HasOnline != 1 {
		return nil, fmt.Errorf("平台接口还没上线！")
	}

	return info, nil
}

func (p *Platform) toSign(param url.Values, secret string) string {
	keys := make([]string, 0)
	for k, _ := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sortParam url.Values
	for _, key := range keys {
		sortParam.Add(key, param.Get(key))
	}

	str := sortParam.Encode()
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func (p *Platform) call(url, method string, body url.Values, header http.Header) ([]byte, error) {

	request, err := http.NewRequest(method, url, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}
	header.Add("user-agent", "nft platform v1.0.0")
	request.Header = header
	client := http.Client{
		Timeout: time.Duration(3) * time.Second,
	}
	do, err := client.Do(request)

	resByte := make([]byte, 1024)
	_, err = bufio.NewReader(do.Body).Read(resByte)
	if err != nil {
		return nil, err
	}
	if do.StatusCode != 200 {
		return nil, fmt.Errorf("status:%s", do.Status)
	}
	code, err := jsonparser.GetInt(resByte, "code")
	if err != nil {
		return nil, err
	}
	msg, err := jsonparser.GetString(resByte, "msg")

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("response code:%d,msg:%s", code, msg)
	}

	data, err := jsonparser.GetString(resByte, "data")

	if err != nil {
		return nil, err
	}
	return []byte(data), nil
}
