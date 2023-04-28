package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"nft_platform/global"
	"sort"
	"strings"
)

func ToSign(data map[string]interface{}, secret string) string {
	keys := make([]string, 0)
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	signStr := make([]string, 0)
	for _, val := range keys {
		signStr = append(signStr, fmt.Sprintf("%s=%s", val, data[val]))
	}

	md5Sign := encode(strings.Join(signStr, "&"), secret)
	return md5Sign
}

func encode(data, secret string) string {
	if len(data) == 0 || len(secret) == 0 {
		return ""
	}
	data = fmt.Sprintf("%s&%s", data, secret)
	global.SLogger.Debugf("sign str:%s", data)
	md := md5.New()
	md.Write([]byte(data))
	return hex.EncodeToString(md.Sum(nil))
}
