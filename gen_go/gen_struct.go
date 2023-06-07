package gen_go

import (
	"github.com/nullzZ/excel2config/model"
	"github.com/nullzZ/excel2config/pkg/file"
	"github.com/nullzZ/excel2config/pkg/str"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type GenConfigStruct struct{}

func (g *GenConfigStruct) Gen(structModel *model.ConfigData) error {
	goStructName := structModel.StructName
	pack := structModel.PackageName
	toPath := structModel.ToPath
	sheetName := structModel.SheetName
	goStruct := NewGoStruct(pack, goStructName)
	goStruct.PriType = structModel.PriType
	goStruct.PriName = structModel.PriName
	for _, m := range structModel.MetaList {
		goFieldName := str.FirstUpper(m.Key)
		jsonsStr := g.toJsonStr(m.Key)
		goStruct.Fields = append(goStruct.Fields, NewGoField(goFieldName, m.Typ, m.Ann, jsonsStr))
	}
	dirPath := filepath.Join(toPath, "config")
	if !file.Exists(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	toPath2 := filepath.Join(dirPath, sheetName+GenRawdataName)
	file, err := os.OpenFile(toPath2, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	t := template.Must(template.New("").Parse(goTmpl))
	err = t.Execute(file, goStruct)
	if err != nil {
		return err
	}
	return nil
}

func (g *GenConfigStruct) toJsonStr(s string) string {
	builder := &strings.Builder{}
	builder.WriteString("`json:\"")
	builder.WriteString(s)
	builder.WriteString("\"`")
	return builder.String()
}
