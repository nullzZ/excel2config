package excel2conf

import "github.com/nullzZ/excel2config/model"

type IGen interface {
	Gen(structModel *model.ConfigData) error
}

type IGlobalGen interface {
	Gen(packaged, toPath string, datas map[string]*model.ConfigData) error
}
