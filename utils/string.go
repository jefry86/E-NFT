package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func AtoInt(str string) int {
	if str == "" {
		return 0
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Errorf("string to int err:%s", err))
	}
	return i
}

func AtoInt64(str string) int64 {
	if str == "" {
		return 0
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(fmt.Errorf("string to int64 err:%s", err))
	}
	return i
}

func RandString(n int) string {
	str := "abcdefghijkmnopqrstuvwsyzABCDEFGHIJKMNOPQRSTUVWSYZ0123456789"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func PageOffset(pageNo, pageSize int) (int, int) {
	if pageNo == 0 && pageSize == 0 {
		return 0, 10
	} else if pageNo == 0 {
		return 0, pageSize
	} else if pageSize == 0 {
		pageSize = 10
	}
	return (pageNo - 1) * pageSize, pageSize
}
