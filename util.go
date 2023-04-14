package excel2conf

import (
	"errors"
	"fmt"
	"github.com/nullzZ/excel2config/pkg/str"
)

func GenStructName(sheetName string) string {
	structName := str.FirstUpper(sheetName)
	return structName
}

func ParseMetaField(rowData string, typ string, isPri bool) (string, error) {
	if isString(typ) {
		if isPri {
			return rowData, nil
		}
		return parseString(rowData), nil
	} else if isInt(typ) {
		if rowData == "" {
			return "0", nil
		}
		v, err := parseInt(typ, rowData)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", v), nil
	} else if isFloat(typ) {
		if rowData == "" {
			return "0", nil
		}
		v, err := parseFloat(typ, rowData)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%f", v), nil
	} else if IsRepeated(typ) {
		return parseRepeatedJson(typ, rowData), nil
	} else if IsRepeated2(typ) {
		return parseRepeated2Json(typ, rowData), nil
	} else if IsMap(typ) {
		return parseMapJson(typ, rowData), nil
	}

	return "", errors.New("parseInt err")
}

// CheckPriRepeat 检测主键是否重复
func CheckPriRepeat(data []*GenJsonData2) error {
	m := make(map[string]bool)
	for _, d := range data {
		if _, ok := m[d.PriKey]; ok {
			return fmt.Errorf("Duplicate primary key id=%s", d.PriKey)
		}
		m[d.PriKey] = true
	}
	return nil
}
