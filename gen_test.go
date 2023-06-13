package excel2conf

import (
	"encoding/json"
	"github.com/nullzZ/excel2config/field"
	"github.com/nullzZ/excel2config/gen/conf_loader"
	"github.com/nullzZ/excel2config/gen/config"
	"github.com/nullzZ/excel2config/gen_go"
	"github.com/nullzZ/excel2config/pkg/checker"
	"github.com/nullzZ/excel2config/pkg/zaplog"
	"go.uber.org/zap"
	"log"
	"testing"
)

func TestGenGo(t *testing.T) {
	sourcePath := "/Users/malei/works/excel2config/data"
	toPath := "/Users/malei/works/excel2config/gen"
	zaplog.Init(zap.DebugLevel)
	gen := NewGenerateExcel(sourcePath, toPath, WithSkipRow(5), WithSkipCol(1))
	gen.AddGen(&gen_go.GenConfigStruct{})
	gen.AddGen(&gen_go.GenRawdataConf{})
	gen.AddGen(&gen_go.GenChecker{})

	gen.AddGlobalGen(&gen_go.GenGlobalLoad{})
	gen.AddGlobalGen(&gen_go.GenGlobalInit{})
	gen.AddGlobalGen(&gen_go.GenGlobalLoader{})
	gen.AddGlobalGen(&gen_go.GenGlobalErr{})

	err := gen.ReadFile()
	if err != nil {
		t.Errorf("err=%q", err)
	}
}

func TestIsRepeated(t *testing.T) {
	ok := field.IsRepeated("<1,2,3>")
	t.Logf("@@@%v", ok)
}

func TestRepeated2(t *testing.T) {
	var a [][]int = make([][]int, 5)
	for i := 0; i < 5; i++ {
		var aa []int
		aa = append(aa, 1)
		aa = append(aa, 2)
		a[i] = aa
	}

	bb, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("@@@@@@@@%s", string(bb))
}

func TestParseRepeated2Json(t *testing.T) {
	str := field.ParseRepeated2Json(field.RepeatedInt2Typ, "[1,2]")
	t.Log(str)
	str2 := field.ParseRepeated2Json(field.RepeatedInt2Typ, "[[1,2]]")
	t.Log(str2)
	str3 := field.ParseRepeated2Json(field.RepeatedInt2Typ, "[[1,2],[1]]")
	t.Log(str3)
}

func TestLoader(t *testing.T) {
	zaplog.Init(zap.DebugLevel)
	config.InitWithLoader(conf_loader.AddLoader)
	conf_loader.MustInitLocal("/Users/malei/works/excel2config/gen/rawdata", true, zaplog.SugaredLogger)

	m := make(map[string]func(i interface{}, param string) bool)
	m["ArrayExist"] = checker.ArrayExist
	config.InitCheckerFunc(m)              //注册自定义注解函数
	conf_loader.AddChecker(config.Checker) //加载检测方法
}
