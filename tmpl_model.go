package excel2conf

type GoStruct struct {
	Package string
	Name    string
	Fields  []*GoField
	PriType string
	PriName string
}

type RawdataConf struct {
	Package string
	Name    string
	PriType string
}

type Checker struct {
	Package  string
	Name     string
	PriName  string
	Checkers []*CheckerField
}

type CheckerField struct {
	ConfName  string
	FieldName string
	Param     string
	PriName   string
}

type GoField struct {
	Name       string
	Kind       string
	Annotation string
	JsonStr    string
}

type GenLoad struct {
	Package string
	Fields  []*GenLoadField
}
type GenLoadField struct {
	Name   string
	PriTyp string
}

type GenInit struct {
	Package string
}

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

func NewGoField(name, kind, annotation, jsonStr string) *GoField {
	return &GoField{
		Name:       name,
		Kind:       convertType(kind),
		Annotation: annotation,
		JsonStr:    jsonStr,
	}
}

func NewGoStruct(pack, name string) *GoStruct {
	return &GoStruct{
		Package: pack,
		Name:    name,
	}
}

func NewGenLoad(pack string) *GenLoad {
	return &GenLoad{
		Package: pack,
	}
}

func NewGenInit(pack string) *GenInit {
	return &GenInit{
		Package: pack,
	}
}

func NewRawdataConf(pack, name, priType string) *RawdataConf {
	return &RawdataConf{
		Package: pack,
		Name:    name,
		PriType: priType,
	}
}

func NewChecker(pack, name, priName string) *Checker {
	return &Checker{
		Package: pack,
		Name:    name,
		PriName: priName,
	}
}
