package main

import (
	"flag"
	excel2conf "github.com/nullzZ/excel2config"
	"github.com/nullzZ/excel2config/pkg/zaplog"
	"go.uber.org/zap"
)

func main() {
	sourcePath := flag.String("s", "", "source")
	toPath := flag.String("t", "", "to")
	flag.Parse()
	zaplog.Init(zap.DebugLevel)
	gen := excel2conf.NewGenerateExcel(*sourcePath, *toPath)
	err := gen.ReadFile()
	if err != nil {
		zaplog.SugaredLogger.Panic(err)
	}
}
