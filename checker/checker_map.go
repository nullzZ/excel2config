package checker

import "github.com/nullzZ/excel2config/gen/config"

// MapValExist map中的val 是否存在于另一配置中
func MapValExist(i interface{}, param string) bool {
	m, ok := i.(map[int32]int32)
	if !ok {
		return false
	}
	if param == "" {
		return false
	}
	for _, val := range m {
		_, ok2 := config.Get(param, val)
		if !ok2 {
			return false
		}
	}
	return true
}
