package model

import (
	"github.com/nullzZ/excel2config/field"
	"github.com/nullzZ/excel2config/pkg/str"
)

type Meta struct {
	Key         string
	Idx         int
	Typ         string
	ClientTyp   string
	Ann         string
	CheckerName string //检测表名称
}

type RowData []string

type ConfigData struct {
	ToPath      string
	SheetName   string //sheet名称
	StructName  string //结构体名称
	PriType     string //主键类型
	PriName     string //主键名称
	MetaList    []*Meta
	PackageName string //package名
	DataRow     []RowData
}

func NewConfigData(toPath, packageName, sheetName string, metaList []*Meta, dataList []RowData) *ConfigData {
	structName := GenStructName(sheetName)
	priType := field.ConvertType(metaList[0].Typ)
	priName := str.FirstUpper(metaList[0].Key)
	model := &ConfigData{
		ToPath:      toPath,
		SheetName:   sheetName,
		StructName:  structName,
		PriType:     priType,
		PriName:     priName,
		PackageName: packageName,
		MetaList:    metaList,
		DataRow:     dataList,
	}
	return model
}

func GenStructName(sheetName string) string {
	structName := str.FirstUpper(sheetName)
	return structName
}
