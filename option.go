package excel2conf

type GenCfg struct {
	SkipRowNumber int //每个sheet开始行数
	SkipColNumber int //跳过列数
}

type GenOption func(cfg *GenCfg)

func WithSkipRow(skipRow int) GenOption {
	return func(cfg *GenCfg) {
		cfg.SkipRowNumber = skipRow
	}
}

func WithSkipCol(skipCol int) GenOption {
	return func(cfg *GenCfg) {
		cfg.SkipColNumber = skipCol
	}
}
