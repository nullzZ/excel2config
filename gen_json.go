package excel2conf

import (
	"github.com/nullzZ/excel2config/pkg/file"
	"github.com/nullzZ/excel2config/pkg/str"
	"os"
	"path/filepath"
	"text/template"
)

type GenJson struct{}

func (g *GenJson) Gen(structModel *ConfigData) error {
	sheetName := structModel.SheetName
	genJsonData := &GenJsonData{
		Name: sheetName,
	}
	for _, row := range structModel.DataRow {
		genJsonData2 := &GenJsonData2{}
		for idx, meta := range structModel.MetaList {
			if idx == 0 {
				priKey, err := ParseMetaField(row[idx], meta.Typ, true)
				if err != nil {
					return err
				}
				genJsonData2.PriKey = priKey
			}
			field, err := ParseMetaField(row[idx], meta.Typ, false)
			if err != nil {
				return err
			}
			genJsonData2.Fields = append(genJsonData2.Fields, &GenJsonField{
				FieldName: meta.Key,
				FieldVal:  field,
			})
		}
		genJsonData2.FieldsLen = len(genJsonData2.Fields) - 1
		genJsonData.Datas = append(genJsonData.Datas, genJsonData2)
	}
	err := CheckPriRepeat(genJsonData.Datas)
	if err != nil {
		return err
	}
	genJsonData.DatasLen = len(genJsonData.Datas) - 1
	dirPath := filepath.Join(structModel.ToPath, "rawdata")
	toPath2 := filepath.Join(dirPath, g.genRawdataName(sheetName)+GenJsonName)
	if !file.Exists(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(toPath2, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	t := template.Must(template.New("").Parse(genJsonTmpl))
	err = t.Execute(file, genJsonData)
	if err != nil {
		return err
	}
	return nil
}

func (g *GenJson) genRawdataName(name string) string {
	goStructName := str.FirstUpper(name)
	return goStructName
}
