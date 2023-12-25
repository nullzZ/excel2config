package excel2conf

import (
	"github.com/nullzZ/excel2config/gen_go"
	"github.com/nullzZ/excel2config/pkg/zaplog"
	"go.uber.org/zap"
)

func Gen(gen *GenerateExcel) {
	zaplog.Init(zap.DebugLevel)
	gen.AddGen(&gen_go.GenConfigStruct{})
	gen.AddGen(&gen_go.GenRawdataConf{})
	gen.AddGen(&gen_go.GenChecker{})
	gen.AddGlobalGen(&gen_go.GenGlobalLoad{})
	gen.AddGlobalGen(&gen_go.GenGlobalInit{})
	//gen.AddGlobalGen(&gen_go.GenGlobalLoader{})
	gen.AddGlobalGen(&gen_go.GenGlobalErr{})
	err := gen.ReadFile()
	if err != nil {
		zaplog.SugaredLogger.Panic(err)
	}
}
