package gen_go

import (
	_ "embed"
	"github.com/nullzZ/excel2config/model"
	"github.com/nullzZ/excel2config/pkg/file"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed gen_rawdata_conf.tmpl
var genRawdataConfTmpl string

type GenRawdataConf struct{}

func (GenRawdataConf) Gen(structModel *model.ConfigData) error {
	goStructName := structModel.StructName
	pack := structModel.PackageName
	sheetName := structModel.SheetName
	goStruct := NewRawdataConf(pack, goStructName, structModel.PriType)
	dirPath := filepath.Join(structModel.ToPath, "config")
	if !file.Exists(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	toPath2 := filepath.Join(dirPath, sheetName+GenRawdataConfFileName)
	file, err := os.OpenFile(toPath2, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	t := template.Must(template.New("").Parse(genRawdataConfTmpl))
	err = t.Execute(file, goStruct)
	if err != nil {
		return err
	}
	return nil
}
