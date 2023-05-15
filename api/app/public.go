package app

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jaevor/go-nanoid"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"nft_platform/global"
	"strings"
	"time"
)

var appid = "1316966420"
var bucket = "cos-1316966420"
var region = "ap-shanghai"

type Public struct {
	UserBase
}

func (p *Public) TCloudSTS(c *gin.Context) {
	client := sts.NewClient(
		// 通过环境变量获取密钥, os.Getenv 方法表示获取环境变量
		global.Conf.TCloud.Sms.SecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考https://cloud.tencent.com/document/product/598/37140
		global.Conf.TCloud.Sms.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考https://cloud.tencent.com/document/product/598/37140
		nil,
		// sts.Host("sts.internal.tencentcloudapi.com"), // 设置域名, 默认域名sts.tencentcloudapi.com
		// sts.Scheme("http"),      // 设置协议, 默认为https，公有云sts获取临时密钥不允许走http，特殊场景才需要设置http
	)
	// 策略概述 https://cloud.tencent.com/document/product/436/18023
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					// 密钥的权限列表。简单上传和分片需要以下的权限，其他权限列表请看 https://cloud.tencent.com/document/product/436/31923
					Action: []string{
						// 简单上传
						"name/cos:PostObject",
						"name/cos:PutObject",
						// 分片上传
						"name/cos:InitiateMultipartUpload",
						"name/cos:ListMultipartUploads",
						"name/cos:ListParts",
						"name/cos:UploadPart",
						"name/cos:CompleteMultipartUpload",
					},
					Effect: "allow",
					Resource: []string{
						// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						// 存储桶的命名格式为 BucketName-APPID，此处填写的 bucket 必须为此格式
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/images",
					},
					// 开始构建生效条件 condition
					// 关于 condition 的详细设置规则和COS支持的condition类型可以参考https://cloud.tencent.com/document/product/436/71306
					Condition: map[string]map[string]interface{}{
						"ip_equal": map[string]interface{}{
							"qcs:ip": []string{
								"10.217.182.3/24",
								"111.21.33.72/24",
							},
						},
					},
				},
			},
		},
	}

	// case 1 请求临时密钥
	res, err := client.GetCredential(opt)
	if err != nil {
		panic(err)
	}

	/*
		cosHost: cosHost,
		cosKey: cosKey,
		policy: Buffer.from(policy).toString('base64'),
		qSignAlgorithm: qSignAlgorithm,
		qAk: config.secretId,
		qKeyTime: qKeyTime,
		qSignature: qSignature,
	*/
	p.JsonSuccessWithData(c, res)
}

func (p *Public) TCloudPolicy(c *gin.Context) {
	ext := c.Query("ext")
	secretId := global.Conf.TCloud.Sms.SecretId
	secretKey := global.Conf.TCloud.Sms.SecretKey
	//bucket := "your_bucket_name"
	//region := "your_bucket_region"
	key := generateCosKey(ext)

	// 过期时间 1 小时
	newTime := time.Now()
	expireTime := newTime.Add(2 * time.Hour).Unix()

	keyTime := fmt.Sprintf("%d;%d", newTime.Unix(), expireTime)

	// 生成 policy
	policy := generatePolicy(expireTime, secretId, bucket, region, key, keyTime)

	// 对 policy 进行 base64 编码
	policyBase64 := base64.StdEncoding.EncodeToString([]byte(policy))

	// 生成签名
	sign := generateSign(policy, secretKey, keyTime)

	// 返回给前端的数据
	responseData := map[string]interface{}{
		"policy":         policyBase64,
		"qSignature":     sign,
		"cosHost":        fmt.Sprintf("%s.cos.%s.myqcloud.com", bucket, region),
		"cosKey":         key,
		"qSignAlgorithm": "sha1",
		"qAk":            secretId,
		"qKeyTime":       keyTime,
	}

	// 将 responseData 转为 JSON 格式并返回
	p.JsonSuccessWithData(c, responseData)
}

func generatePolicy(expireTime int64, secretId, bucket, region, key, keyTime string) string {
	policyMap := map[string]interface{}{
		"expiration": time.Unix(expireTime, 0).Format("2006-01-02T15:04:05.999Z"), //2019-08-30T09:38:12.414Z
		"conditions": []interface{}{
			map[string]string{"bucket": bucket},
			map[string]string{"key": key},
			//[]string{"starts-with", "$key", key},
			map[string]string{"q-sign-algorithm": "sha1"},
			map[string]string{"q-ak": secretId},
			map[string]string{"q-sign-time": keyTime},
			//map[string]string{"q-key-time": fmt.Sprintf("%d;%d", time.Now().Unix(), expireTime)},
			//map[string]string{"q-header-list": "content-type;host"},
			//map[string]string{"q-url-param-list": ""},
			//map[string]string{"q-signature": ""},
		},
	}
	/*
	   {'q-sign-algorithm': qSignAlgorithm},
	               {'q-ak': config.secretId},
	               {'q-sign-time': qKeyTime},
	               {'bucket': config.bucket},
	               {'key': cosKey},
	*/

	policyJson, _ := json.Marshal(policyMap)
	global.SLogger.Debugf("id:%s key:%s", global.Conf.TCloud.Sms.SecretId, global.Conf.TCloud.Sms.SecretKey)
	global.SLogger.Debugf("policy:%s\n", policyJson)

	return strings.ReplaceAll(string(policyJson), "\\", "")
}

func generateSign(policy, secretKey, signTime string) string {
	signKey := hmacSha1(signTime, secretKey)
	stringToSign := SHA1(policy)
	sign := hmacSha1(stringToSign, signKey)
	global.SLogger.Debugf("policy:%s \nsignKey:%s \nstringToSign:%s", policy, signKey, stringToSign)
	return sign
}

func SHA1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func hmacSha1(data, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func generateCosKey(ext string) string {
	nid, err := nanoid.Standard(20)
	if err != nil {
		return ""
	}
	var cosKey = fmt.Sprintf("images/%s.%s", nid(), ext)
	return cosKey
}
