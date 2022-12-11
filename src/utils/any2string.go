// Package utils
// @Description
// @Author root_wang
// @Date 2022/12/11 13:42
package utils

import (
	"log"
	"reflect"
	"strconv"
)

func Any2string(v interface{}) (str string) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String:
		str = v.(string)
	case reflect.Int64:
		str = strconv.FormatInt(v.(int64), 10)
	case reflect.Bool:
		if v == true {
			str = "true"
		} else {
			str = "false"
		}
	default:
		log.Fatalf("value: %v 2 string failed", v)
	}
	return
}