package excel2conf

import (
	"errors"
	"github.com/nullzZ/excel2config/pkg/zaplog"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"path"
	"strings"
	"unicode"
)

/**
 * 将excel中的前四列转化为struct
 * 第一列字段名		如 id
 * 第二列字段类型		如 int
 * 第三列备注
 * 第四列s,c,all 	s表示服务端使用 c表示客户端使用 all表示都使用
 */
type GenerateExcel struct {
	toPath        string //存储路径
	sourcePath    string //源路径
	genList       []IGen
	genGlobalList []IGlobalGen
	cfgDatas      map[string]*ConfigData
}

func NewGenerateExcel(sourcePath, toPath string) *GenerateExcel {
	gen := &GenerateExcel{
		toPath:        toPath,
		sourcePath:    sourcePath,
		genList:       make([]IGen, 0, 10),
		genGlobalList: make([]IGlobalGen, 0, 10),
		cfgDatas:      make(map[string]*ConfigData, 1),
	}
	gen.AddGen(&GenConfigStruct{})
	gen.AddGen(&GenRawdataConf{})
	gen.AddGen(&GenChecker{})

	gen.AddGlobalGen(&GenGlobalLoad{})
	gen.AddGlobalGen(&GenGlobalInit{})
	return gen
}

func (g *GenerateExcel) AddGen(gen IGen) {
	g.genList = append(g.genList, gen)
}
func (g *GenerateExcel) AddGlobalGen(gen IGlobalGen) {
	g.genGlobalList = append(g.genGlobalList, gen)
}

func (g *GenerateExcel) ReadFile() error {
	if g.sourcePath == "" {
		return ErrorSourcePath
	}
	if g.toPath == "" {
		return ErrorToPath
	}

	files, err := ioutil.ReadDir(g.sourcePath)
	if err != nil {
		return err
	}
	sheetNames := make([]string, 0)
	for _, file := range files {
		if g.isFileContinue(file.Name()) {
			continue
		}

		p1 := path.Join(g.sourcePath, file.Name())
		f, err := xlsx.OpenFile(p1)
		if err != nil {
			return err
		}
		zaplog.SugaredLogger.Debugf("gen---file=%s", file.Name())
		for _, sheet := range f.Sheets {
			if g.isContinue(sheet) { //跳过sheet
				continue
			}

			if sheet.MaxRow < LineNumber+1 {
				zaplog.SugaredLogger.Errorf("ReadFile sheet.MaxRow < %d sheet=%s", LineNumber+1, sheet.Name)
				return ErrorMaxRow
			}

			zaplog.SugaredLogger.Debugf("gen---sheet%s", sheet.Name)
			sheetName := sheet.Name
			if strings.Contains(sheet.Name, "_s") { //去掉后缀
				sheetName = strings.TrimRight(sheetName, "_s")
			} else if strings.Contains(sheet.Name, "_c") {
				sheetName = strings.TrimRight(sheetName, "_c")
			}
			err2 := g.GenSheet(sheet, sheetName, Packaged)
			if err2 != nil {
				zaplog.SugaredLogger.Panicf("ReadFile sheet=%s gen err %v", sheetName, err2)
				return err2
			}
			sheetNames = append(sheetNames, sheetName)
		}
	}
	for _, gen := range g.genGlobalList {
		err2 := gen.Gen(Packaged, g.toPath, g.cfgDatas)
		if err2 != nil {
			zaplog.SugaredLogger.Panicf("genGlobalList gen err %v", err2)
		}
	}
	zaplog.SugaredLogger.Debugf("gen---success!!!")
	return nil
}

func (g *GenerateExcel) GenSheet(sheet *xlsx.Sheet, sheetName, pack string) error {
	colNum := sheet.MaxCol
	contunies := make(map[int]bool)
	metaList := make([]*Meta, 0, colNum)
	dataList := make([]RowData, 0, len(sheet.Rows)-LineNumber)
	for i := 0; i < colNum; i++ {
		if i < SkipColNumber {
			continue
		}
		fieldName := sheet.Cell(FieldName, i)           //字段名字
		fieldTyp := sheet.Cell(ServerFieldTyp, i)       //后端字段类型
		clientFieldTyp := sheet.Cell(ClientFieldTyp, i) //前端字段类型
		fieldAnn := sheet.Cell(FieldAnn, i)             //备注
		if g.isMetaContinue(fieldName.Value) {
			contunies[i] = true
			continue
		}
		if g.isMetaContinue(fieldTyp.Value) {
			contunies[i] = true
			continue
		}
		typs := strings.Split(fieldTyp.Value, "@") //用@符号分割
		m := &Meta{
			Key:       fieldName.Value,
			Idx:       i,
			Typ:       typs[0], //fieldTyp.Value
			Ann:       fieldAnn.Value,
			ClientTyp: clientFieldTyp.Value,
		}

		if len(typs) > 1 {
			m.CheckerName = typs[1]
		}
		if !checkFieldType(m.Typ) {
			zaplog.SugaredLogger.Errorf("ReadFile sheet=%s m= %v", sheetName, m)
			return errors.New("type err")
		}
		metaList = append(metaList, m)
	}

	for i, _ := range sheet.Rows {
		if i < LineNumber {
			continue
		}
		data := make(RowData, 0, colNum)
		for j := 0; j < colNum; j++ {
			if j < SkipColNumber {
				continue
			}
			c := sheet.Cell(i, j)
			if _, ok := contunies[j]; ok {
				continue
			}
			data = append(data, c.Value)
		}
		dataList = append(dataList, data)
	}

	configData := NewConfigData(g.toPath, pack, sheetName, metaList, dataList)
	g.cfgDatas[sheetName] = configData //存一下 后面使用
	for _, gen := range g.genList {
		err := gen.Gen(configData)
		if err != nil {
			zaplog.SugaredLogger.Errorf("ReadFile sheet=%s err= %v", sheetName, err)
			return err
		}
	}
	//fmt.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@%v", metaList)
	//fmt.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@%v", dataList)
	return nil
}

func (GenerateExcel) isContinue(sheet *xlsx.Sheet) bool {
	if strings.HasPrefix(sheet.Name, "#") {
		return true
	}
	if strings.Index(sheet.Name, "Sheet") != -1 { //排除默认sheet
		return true
	}
	if strings.Contains(sheet.Name, "_c") { //_c 前端使用
		return true
	}

	for _, v := range []rune(sheet.Name) { //排除汉子
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

func (GenerateExcel) isFileContinue(fileName string) bool { //是否跳过excel表格
	if path.Ext(fileName) != ".xlsx" {
		return true
	}
	if strings.Contains(fileName, "#") {
		return true
	}
	return false
}

func (g *GenerateExcel) isMetaContinue(val string) bool {
	if val == "" {
		return true
	}
	//if strings.Contains(name, "_C") {
	//	return true
	//}
	return false
}
