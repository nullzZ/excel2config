package gen_go

import (
	_ "embed"
	"github.com/nullzZ/excel2config/model"
	"github.com/nullzZ/excel2config/pkg/file"
	"os"
	"path/filepath"
	"text/template"
)

type GenGlobalLoader struct{}

//go:embed loader.tmpl
var loaderTmpl string

func (GenGlobalLoader) Gen(packaged, toPath string, rawdataConfs map[string]*model.ConfigData) error {
	gen := NewGenLoader(packaged)
	dirPath := filepath.Join(toPath, "conf_loader")
	if !file.Exists(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	toPath2 := filepath.Join(dirPath, GenLoaderName)
	file, err := os.OpenFile(toPath2, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	t := template.Must(template.New("").Parse(loaderTmpl))
	err = t.Execute(file, gen)
	if err != nil {
		return err
	}
	return nil
}
