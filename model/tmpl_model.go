package model

type GenJsonData struct {
	Name     string
	Datas    []*GenJsonData2
	DatasLen int
}

type GenJsonData2 struct {
	PriKey    string
	Fields    []*GenJsonField
	FieldsLen int
}

type GenJsonField struct {
	FieldName string
	FieldVal  string
}
