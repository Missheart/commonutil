package util

import (
	"strings"

	"github.com/google/uuid"
)

func GetUuid() string {
	// 获取当前goroutine ID并记录到zap日志中
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
