package gen_go

import (
	_ "embed"
	"github.com/nullzZ/excel2config/model"
	"github.com/nullzZ/excel2config/pkg/file"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed gen_load.tmpl
var genLoadTmpl string

type GenGlobalLoad struct{}

func (GenGlobalLoad) Gen(packaged, toPath string, configDatas *[]*model.ConfigData) error {
	gen := NewGenLoad(packaged)
	for _, v := range *configDatas {
		gen.Fields = append(gen.Fields, &GenLoadField{
			Name:   v.StructName,
			PriTyp: v.PriType,
		})
	}
	dirPath := filepath.Join(toPath, "config")
	if !file.Exists(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	toPath2 := filepath.Join(dirPath, GenLoadName)
	file, err := os.OpenFile(toPath2, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	t := template.Must(template.New("").Parse(genLoadTmpl))
	err = t.Execute(file, gen)
	if err != nil {
		return err
	}
	return nil
}
