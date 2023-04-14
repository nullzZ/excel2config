package excel2conf

import (
	"github.com/nullzZ/excel2config/pkg/file"
	"os"
	"path/filepath"
	"text/template"
)

type GenGlobalInit struct{}

func (GenGlobalInit) Gen(packaged, toPath string, datas map[string]*ConfigData) error {
	gen := NewGenInit(packaged)

	dirPath := filepath.Join(toPath, "config")
	if !file.Exists(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	toPath2 := filepath.Join(dirPath, GenInitConfigFileName)
	file, err := os.OpenFile(toPath2, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	t := template.Must(template.New("").Parse(genInitTmpl))
	err = t.Execute(file, gen)
	if err != nil {
		return err
	}
	return nil
}
