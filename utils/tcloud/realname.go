package tcloud

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/buger/jsonparser"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"nft_platform/global"
	"strings"
	"time"
)

type RealName struct {
}

func (r *RealName) Check(name, card string) (bool, error) {
	// 云市场分配的密钥Id
	secretId := global.Conf.TCloud.Realname.SecretId
	// 云市场分配的密钥Key
	secretKey := global.Conf.TCloud.Realname.SecretKey
	source := "market"

	// 签名
	auth, datetime, _ := r.calcAuthorization(source, secretId, secretKey)

	// 请求方法
	method := "POST"
	// 请求头
	headers := map[string]string{"X-Source": source, "X-Date": datetime, "Authorization": auth}

	// 查询参数
	queryParams := make(map[string]string)

	// body参数
	bodyParams := make(map[string]string)
	bodyParams["cardNo"] = card
	bodyParams["realName"] = name
	// url参数拼接
	url := "https://service-18c38npd-1300755093.ap-beijing.apigateway.myqcloud.com/release/idcard/VerifyIdcardv2"
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s?%s", url, r.urlencode(queryParams))
	}

	bodyMethods := map[string]bool{"POST": true, "PUT": true, "PATCH": true}
	var body io.Reader = nil
	if bodyMethods[method] {
		body = strings.NewReader(r.urlencode(bodyParams))
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return false, err
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	errorCode, err := jsonparser.GetInt(bodyBytes, "error_code")
	if err != nil {
		return false, err
	}
	if errorCode != 0 {
		return false, fmt.Errorf("return error code:%d", errorCode)
	}

	isOk, err := jsonparser.GetBoolean(bodyBytes, "result", "isok")
	if err != nil {
		return false, err
	}
	if !isOk {
		return false, nil
	}
	return true, nil

}

func (r *RealName) calcAuthorization(source string, secretId string, secretKey string) (auth string, datetime string, err error) {
	timeLocation, _ := time.LoadLocation("Etc/GMT")
	datetime = time.Now().In(timeLocation).Format("Mon, 02 Jan 2006 15:04:05 GMT")
	signStr := fmt.Sprintf("x-date: %s\nx-source: %s", datetime, source)

	// hmac-sha1
	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(signStr))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	auth = fmt.Sprintf("hmac id=\"%s\", algorithm=\"hmac-sha1\", headers=\"x-date x-source\", signature=\"%s\"",
		secretId, sign)

	return auth, datetime, nil
}

func (r *RealName) urlencode(params map[string]string) string {
	var p = url.Values{}
	for k, v := range params {
		p.Add(k, v)
	}
	return p.Encode()
}
