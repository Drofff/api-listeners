package util

import "time"

func UnixMillis() int64 {
	unixSec := time.Now().Unix()
	return unixSec * 1000
}