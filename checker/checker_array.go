package checker

import "github.com/nullzZ/excel2config/gen/config"

// ArrayExist 数组里的数据是否存在于另一配置中 例子:ArrayExist?OutputInfoData
func ArrayExist(i interface{}, param string) bool {
	arr, ok := i.([]int32)
	if !ok {
		return false
	}
	if param == "" {
		return false
	}
	for _, id := range arr {
		_, ok2 := config.Get(param, id)
		if !ok2 {
			return false
		}
	}
	return true
}
