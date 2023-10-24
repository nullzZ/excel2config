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
	skipRow := flag.Int("skipRow", 5, "skipRow")
	skipCol := flag.Int("skipCol", 1, "skipCol")
	flag.Parse()
	zaplog.Init(zap.DebugLevel)
	gen := excel2conf.NewGenerateExcel(*sourcePath, *toPath,
		excel2conf.WithSkipRow(*skipRow),
		excel2conf.WithSkipCol(*skipCol))
	excel2conf.Gen(gen)
}
