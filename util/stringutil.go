/**
 * 功能描述:
 * @Date: 2021/1/21
 * @author: lixiaoming
 */
package util

import (
	"encoding/json"
	"github.com/xm5646/log"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2Byte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func SplitAndTrim(str, sep string) []string {
	parts := strings.Split(str, sep)
	result := make([]string, 0, len(parts))
	for _, s := range parts {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		result = append(result, s)
	}
	return result
}

// StringInArry check if str exist in arr
func StringInArry(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// UniqArr remove duplicate item
func UniqArr(arrs ...[]string) []string {
	dic := make(map[string]struct{}, 5)
	for _, arr := range arrs {
		for _, item := range arr {
			dic[item] = struct{}{}
		}
	}
	result := make([]string, 0, len(dic))
	for k := range dic {
		result = append(result, k)
	}
	return result
}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func Struct2String(s interface{}) string {
	bytes, err := json.Marshal(s)
	if err != nil {
		return ""
	} else {
		return Byte2String(bytes)
	}
}

func String2Time(s, format string) (*time.Time, error) {
	time, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		log.Errorf(err, "failed to convert string to time.")
		return nil, err
	}
	return &time, nil
}

