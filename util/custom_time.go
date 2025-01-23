package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

/**
 * 自定义时间类型
 */
type CustomTime struct {
	time.Time
}

/**
 * 实现json序列化
 */
func (t CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))), nil
}

/**
 * 实现数据库查询映射到自定义的类型
 */
func (t *CustomTime) Scan(value interface{}) error {
	v := value.(time.Time)
	t.Time = v
	return nil
}

/**
 * 实现自定义类型转换为数据库可识别类型
 */
func (t CustomTime) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}
