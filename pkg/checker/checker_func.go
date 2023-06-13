package checker

import (
	"github.com/nullzZ/excel2config/gen/config"
	"strconv"
	"strings"
)

// ResourceMap 奖励判断 map类型
func ResourceMap(i interface{}, param string) bool {
	d, ok := i.(map[int32]int32)
	if !ok {
		return false
	}
	for k, _ := range d {
		_, h := config.GetItemData(k)
		if !h {
			return false
		}
	}
	return true
}

// Resource 奖励判断 int32类型
func Resource(i interface{}, param string) bool {
	id, ok := i.(int32)
	if !ok {
		return false
	}
	_, h := config.GetItemData(id)
	if !h {
		return false
	}
	return true
}

// Scope 区间值判断 例子:Scope?1,2
func Scope(i interface{}, param string) bool {
	d, ok := i.(int32)
	if !ok {
		return false
	}
	if param == "" {
		return false
	}
	ss := strings.Split(param, ",")
	if len(ss) != 2 {
		return false
	}
	min, err := strconv.ParseInt(ss[0], 10, 32)
	if err != nil {
		return false
	}
	max, err := strconv.ParseInt(ss[1], 10, 32)
	if err != nil {
		return false
	}
	if d < int32(min) {
		return false
	}
	if d > int32(max) {
		return false
	}
	return true
}

// Exist 例子:Exist?ItemData
func Exist(i interface{}, param string) bool {
	id, ok := i.(int32)
	if !ok {
		return false
	}
	if param == "" {
		return false
	}
	_, ok2 := config.Get(param, id)
	if !ok2 {
		return false
	}
	return true
}

//func isIn(param string, id int32) bool {
//	switch param {
//	case "OutputInfoData":
//		_, h := config.GetOutputInfoData(id)
//		return h
//	}
//	return false
//}
