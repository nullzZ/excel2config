package excel2conf

type IGen interface {
	Gen(structModel *ConfigData) error
}

type IGlobalGen interface {
	Gen(packaged, toPath string, datas map[string]*ConfigData) error
}
