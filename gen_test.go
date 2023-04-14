package excel2conf

import (
	"encoding/json"
	"github.com/nullzZ/excel2config/pkg/zaplog"
	"go.uber.org/zap"
	"log"
	"testing"
)

func TestGen(t *testing.T) {
	sourcePath := "/Users/malei/works/excel2config/data"
	toPath := "/Users/malei/works/excel2config/gen"
	zaplog.Init(zap.DebugLevel)
	gen := NewGenerateExcel(sourcePath, toPath)
	err := gen.ReadFile()
	if err != nil {
		t.Errorf("err=%q", err)
	}
}

func TestIsRepeated(t *testing.T) {
	ok := IsRepeated("<1,2,3>")
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
	str := parseRepeated2Json(repeatedInt2Typ, "[1,2]")
	t.Log(str)
	str2 := parseRepeated2Json(repeatedInt2Typ, "[[1,2]]")
	t.Log(str2)
	str3 := parseRepeated2Json(repeatedInt2Typ, "[[1,2],[1]]")
	t.Log(str3)
}
