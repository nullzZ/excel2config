package excel2conf

var goTmpl = `// Package {{.Package}} generate. DO NOT EDIT.
package {{.Package}}

type {{.Name}} struct {
{{- range .Fields}}
	// {{.Annotation}}
	{{.Name}} {{.Kind}} {{.JsonStr}}
{{- end}}
}
`

var genRawdataConfTmpl = `// Package {{.Package}} generate. DO NOT EDIT.
package {{.Package}}


type {{.Name}}Conf struct {
	{{.Name}}s         map[{{.PriType}}]*{{.Name}}
}

func (m *{{.Name}}Conf) Get{{.Name}}s() map[{{.PriType}}]*{{.Name}} {
	if m != nil {
		return m.{{.Name}}s
	}
	return nil
}
`

var genLoadTmpl = `// Package {{.Package}} generate. DO NOT EDIT.
package {{.Package}}

import (
	"encoding/json"
	"errors"
	"sync"
)

var (
	LoadedCallBackWithName func(path string)
	ErrNotExist            = errors.New("Not Found by your ID")

{{- range .Fields }}
	Name{{.Name}}Conf	= "{{.Name}}_conf.gen.json"
	count{{.Name}}		= 0
	lock{{.Name}}		sync.RWMutex
	conf{{.Name}}		*{{.Name}}Conf
	checker{{.Name}}	*{{.Name}}ConfChecker
{{- end}}
)

func loadDataWithPath(buff []byte, msg interface{}) error {
	err := json.Unmarshal(buff, msg)
	if err != nil {
		return err
	}
	return nil
}

func InitCheckerFunc(checkerFunc map[string]func(i interface{}, param string) bool) {
	checkers=checkerFunc
}

func InitWithLoader(f func(key string, f func(b []byte) error)) {
	addLoaderFunc = f

{{- range .Fields}}
	addLoaderFunc(Name{{.Name}}Conf, func(buff []byte) error {
		tmp := new({{.Name}}Conf)
		err := loadDataWithPath(buff, tmp)
		if err != nil {
			return err
		}
		lock{{.Name}}.Lock()
		conf{{.Name}} = tmp
		count{{.Name}} = len(tmp.{{.Name}}s)
		lock{{.Name}}.Unlock()
		if LoadedCallBackWithName != nil {
			LoadedCallBackWithName(Name{{.Name}}Conf)
		}
		return nil
	})
{{- end}}
	initGetFunc()
}

func Checker() bool{
	ok := true
	ret := true
{{- range .Fields}}
	ret = checker{{.Name}}.Check()
	if !ret {
		ok = false
	}
{{- end}}
	return ok
}

func initGetFunc() {
{{- range .Fields}}
	getFunc["{{.Name}}"] = func(key interface{}) (interface{}, bool) {
		var id {{.PriTyp}}
		switch key.(type) {
		case int:
			k2, ok := key.(int)
			if !ok {
				return nil, false
			}
			id = {{.PriTyp}}(k2)
		default:
			k3, ok := key.({{.PriTyp}})
			if !ok {
				return nil, false
			}
			id = k3
		}
		c, h := Get{{.Name}}(id)
		return c, h
	}
{{- end}}

}


func Get(confName string, key interface{}) (interface{}, bool) {
	f, h := getFunc[confName]
	if !h {
		return nil, false
	}
	return f(key)
}

{{- range .Fields}}

func Get{{.Name}}(key {{.PriTyp}}) (*{{.Name}}, bool) {
	if item, ok := get{{.Name}}Map()[key]; ok {
		return item, ok
	}
	return nil, false
}

func get{{.Name}}Map() map[{{.PriTyp}}]*{{.Name}} {
	return Get{{.Name}}Conf().{{.Name}}s
}

func Get{{.Name}}Conf() (ret *{{.Name}}Conf) {
	lock{{.Name}}.RLock()
	ret = conf{{.Name}}
	lock{{.Name}}.RUnlock()
	return
}

func Count{{.Name}}Map() (c int) {
	lock{{.Name}}.RLock()
	c = count{{.Name}}
	lock{{.Name}}.RUnlock()
	return
}
{{- end}}
`

var genInitTmpl = `// Package {{.Package}} generate. DO NOT EDIT.
package {{.Package}}


var addLoaderFunc = func(key string, f func(b []byte) error) { panic("please init addLoaderFunc first.") }
var checkers = make(map[string]func(i interface{}, param string) bool)
var getFunc = make(map[string]func(b interface{}) (interface{}, bool), 0)

`

var genJsonTmpl = `{
  "{{.Name}}s": 
{{- range $i,$v := .Datas}}
{{- if eq $i 0}}
	{
{{- end}}
    "{{$v.PriKey}}": 
		{
{{- range $j,$vv := $v.Fields}}
		  "{{$vv.FieldName}}": {{$vv.FieldVal}}
{{- if lt $j $v.FieldsLen}}
				,
{{- end}}
{{- end}}
    	}
{{- if lt $i $.DatasLen}}
	,
{{- end}}
{{- end}}
	}
}
`

var genCheckerTmpl = `// Package {{.Package}} generate. DO NOT EDIT.
package {{.Package}}

import (
	"log"
)

type {{.Name}}ConfChecker struct {
}

func (m *{{.Name}}ConfChecker) Check() (ok bool) {
	ok = true
{{- range $i,$v := .Checkers}}
	for _, data := range get{{$.Name}}Map() {
		f, h := checkers["{{$v.ConfName}}"]
		if !h {
			log.Printf("Error config check  {{$.Name}} %s annotation error  \n", "{{$v.FieldName}}")
			ok = false
		} else {
			if h := f(data.{{$v.FieldName}}, "{{$v.Param}}"); !h {
				log.Printf("Error config check  %s id:%d %s, data error \n", "{{$.Name}}", data.{{$v.PriName}}, "{{$v.FieldName}}")
				ok = false
			}
		}	
	}
{{- end}}
	if ok {
		log.Println("check {{$.Name}} success")
	}
	return ok
}

`
