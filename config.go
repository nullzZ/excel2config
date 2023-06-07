package excel2conf

import "errors"

const (
	FieldName      = 1
	ClientFieldTyp = 2
	ServerFieldTyp = 3
	FieldAnn       = 4 //注释
	Packaged       = "config"
	GenJsonName    = "_conf.gen.json"
)

var (
	ErrorSourcePath = errors.New("ReadExcel sourcePath nil")
	ErrorToPath     = errors.New("ReadExcel stoPath nil")
	ErrorMaxRow     = errors.New("ReadExcel row<4")
)
