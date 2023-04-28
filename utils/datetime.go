package utils

import "time"

func NowUnix() uint {
	unix := time.Now().Unix()
	return uint(unix)
}
