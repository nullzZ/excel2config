package gen_go

import (
	_ "embed"
	"fmt"
	"github.com/nullzZ/excel2config/model"
	"github.com/nullzZ/excel2config/pkg/file"
	"github.com/nullzZ/excel2config/pkg/str"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed checker.tmpl
var genCheckerTmpl string

type GenChecker struct{}

func (GenChecker) Gen(structModel *model.ConfigData) error {
	goStructName := structModel.StructName
	sheetName := structModel.SheetName
	pack := structModel.PackageName
	goStruct := NewChecker(pack, goStructName, structModel.PriName)
	for _, m := range structModel.MetaList {
		if m.CheckerName == "" {
			continue
		}
		ss := strings.Split(m.CheckerName, "?")
		if len(ss) < 1 {
			return fmt.Errorf("genChecker err sheetName=%s", sheetName)
		}
		checkerAnn := ss[0] //注解
		cf := &CheckerField{
			ConfName:  checkerAnn,
			FieldName: str.FirstUpper(m.Key),
			PriName:   goStruct.PriName,
		}
		if len(ss) > 1 {
			cf.Param = ss[1]
		}
		goStruct.Checkers = append(goStruct.Checkers, cf)

	}
	//goStruct.PriType = convertType(metaList[0].Typ)
	dirPath := filepath.Join(structModel.ToPath, "config")
	if !file.Exists(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	toPath2 := filepath.Join(dirPath, sheetName+GenCheckerConf)
	file, err := os.OpenFile(toPath2, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	t := template.Must(template.New("").Parse(genCheckerTmpl))
	err = t.Execute(file, goStruct)
	if err != nil {
		return err
	}
	return nil
}
