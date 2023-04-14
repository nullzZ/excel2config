package excel2conf

import "errors"

const (
	LineNumber    = 5 //每个sheet开始行数
	SkipColNumber = 1 //跳过列数

	FieldName              = 1
	ClientFieldTyp         = 2
	ServerFieldTyp         = 3
	FieldAnn               = 4 //注释
	Packaged               = "config"
	GenRawdataName         = ".gen.go"
	GenJsonName            = "_conf.gen.json"
	GenLoadName            = "gen_load.gen.go"
	GenRawdataConfFileName = "_conf.gen.go"
	GenInitConfigFileName  = "gen_init.go"
	GenCheckerConf         = "_checker_conf.gen.go"
)

var (
	ErrorSourcePath = errors.New("ReadExcel sourcePath nil")
	ErrorToPath     = errors.New("ReadExcel stoPath nil")
	ErrorMaxRow     = errors.New("ReadExcel row<4")
)
