package excel2conf

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

const (
	stringTyp = "string"

	intTyp   = "int"
	int32Typ = "int32"
	int64Typ = "long"

	float32Typ = "float"
	float64Typ = "double"

	repeatedStringTyp  = "string[]"
	repeatedIntTyp     = "int[]"
	repeatedInt32Typ   = "int32[]"
	repeatedInt64Typ   = "int64[]"
	repeatedFloat32Typ = "float[]"
	repeatedFloat64Typ = "double[]"
	repeatedLongTyp    = "long[]"

	repeatedInt2Typ    = "int[][]"
	repeatedString2Typ = "string[][]"
	repeatedFloat2Typ  = "float[][]"
	repeatedDouble2Typ = "double[][]"
	repeatedLong2Typ   = "long[][]"

	mapInt32String  = "map<int32,string>"
	mapInt32In32    = "map<int32,int32>"
	mapInt64String  = "map<int64,string>"
	mapInt64Int32   = "map<int64,int32>"
	mapInt64Int64   = "map<int64,int64>"
	mapStringString = "map<string,string>"
	mapStringInt32  = "map<string,int32>"
	mapIntInt       = "map<int,int>"
	mapLongLong     = "map<long,long>"
	mapLongInt      = "map<long,int>"
	mapIntString    = "map<int,string>"
	mapInt32Float32 = "map<int32,float>"
	mapIntFloat32   = "map<int,float>"
	mapInt32Float64 = "map<int32,double>"

	//repeatedStringTyp  = "repeated<string>"
	//repeatedIntTyp     = "repeated<int>"
	//repeatedInt32Typ   = "repeated<int32>"
	//repeatedInt64Typ   = "repeated<int64>"
	//repeatedFloat32Typ = "repeated<float32>"
	//repeatedFloat64Typ = "repeated<float64>"
	//
	//repeatedInt2Typ    = "repeated<repeated<int>>"
	//repeatedString2Typ = "repeated<repeated<string>>"
)

func checkFieldType(t string) bool {
	switch t {
	case stringTyp:
		return true
	case intTyp, int32Typ, int64Typ:
		return true
	case float32Typ, float64Typ:
		return true
	case repeatedStringTyp, repeatedIntTyp, repeatedInt32Typ, repeatedInt64Typ,
		repeatedFloat32Typ, repeatedFloat64Typ, repeatedLongTyp, repeatedLong2Typ:
		return true
	case mapInt32String, mapInt32In32, mapInt64String, mapInt64Int32,
		mapInt64Int64, mapInt32Float32, mapInt32Float64, mapStringString,
		mapStringInt32, mapIntInt, mapIntString, mapIntFloat32, mapLongLong, mapLongInt:
		return true
	case repeatedInt2Typ, repeatedString2Typ, repeatedFloat2Typ, repeatedDouble2Typ:
		return true
	default:
		return false
	}
}

func isString(typ string) bool {
	return typ == stringTyp
}

func parseString(val string) string {
	if val == "" {
		return "\"\""
	} else {
		b, _ := json.Marshal(val)
		return string(b)
	}
}

func isInt(typ string) bool {
	switch typ {
	case intTyp, int32Typ, int64Typ:
		return true
	default:
		return false
	}
}

func parseInt(typ, val string) (int64, error) {
	switch typ {
	case intTyp, int32Typ:
		return strconv.ParseInt(val, 10, 32)
	case int64Typ:
		return strconv.ParseInt(val, 10, 64)
	default:
		return 0, errors.New("parseInt err")
	}
}

func isFloat(typ string) bool {
	switch typ {
	case float32Typ:
	case float64Typ:
	default:
		return false
	}
	return true
}

func parseFloat(typ, val string) (float64, error) {
	switch typ {
	case float32Typ:
		return strconv.ParseFloat(val, 32)
	case float64Typ:
		return strconv.ParseFloat(val, 64)
	default:
		return 0, errors.New("parseFloat err")
	}
}

// isRepeated 1,2,3,4
func IsRepeated(typ string) bool {
	if typ == "" || !strings.Contains(typ, "[]") {
		return false
	}
	switch typ {
	case repeatedFloat64Typ, repeatedStringTyp, repeatedIntTyp, repeatedInt32Typ,
		repeatedInt64Typ, repeatedFloat32Typ, repeatedLongTyp:
		return true
	default:
		return false
	}
}

func IsRepeated2(typ string) bool {
	if typ == "" || !strings.Contains(typ, "[][]") {
		return false
	}
	switch typ {
	case repeatedInt2Typ, repeatedString2Typ, repeatedFloat2Typ, repeatedDouble2Typ, repeatedLong2Typ:
		return true
	default:
		return false
	}
}

func parseRepeatedJson(typ, val string) string {
	val = strings.TrimFunc(val, func(r rune) bool { //处理掉特殊字符
		if r == '[' {
			return true
		} else if r == ']' {
			return true
		}
		return false
	})
	builder := &strings.Builder{}
	builder.WriteString("[")
	if val == "" {
		builder.WriteString("]")
		return builder.String()
	}
	strs := strings.Split(val, ",")
	builder2 := &strings.Builder{}
	if typ == repeatedStringTyp {
		for _, str := range strs {
			if str == "" {
				str = ""
			}
			builder2.WriteString("\"")
			builder2.WriteString(str)
			builder2.WriteString("\"")
			builder2.WriteString(",")
		}
	} else {
		for _, str := range strs {
			builder2.WriteString(str)
			builder2.WriteString(",")
		}
	}
	s := strings.TrimRight(builder2.String(), ",")
	builder.WriteString(s)
	builder.WriteString("]")
	return builder.String()
}

func parseRepeated2Json(typ, val string) string {
	//val = strings.TrimFunc(val, func(r rune) bool { //处理掉特殊字符
	//	if r == '[[' {
	//		return true
	//	} else if r == ']' {
	//		return true
	//	}
	//	return false
	//})
	val = strings.TrimLeft(val, "[[")
	val = strings.TrimRight(val, "]]")

	builder := &strings.Builder{}
	builder.WriteString("[")
	if val == "" {
		builder.WriteString("]")
		return builder.String()
	}
	strs := strings.Split(val, "],[")
	if typ == repeatedString2Typ {
		builder3 := &strings.Builder{}
		for _, str := range strs {
			builder2 := &strings.Builder{}
			builder2.WriteString("[")
			sstr := strings.Split(str, ",")
			for _, v := range sstr {
				if v == "" {
					v = ""
				}
				builder2.WriteString("\"")
				builder2.WriteString(v)
				builder2.WriteString("\"")
				builder2.WriteString(",")
			}
			ss := strings.TrimRight(builder2.String(), ",")
			builder3.WriteString(ss)
			builder3.WriteString("]")
			builder3.WriteString(",")
		}
		s := strings.TrimRight(builder3.String(), ",")
		builder.WriteString(s)
	} else {
		builder3 := &strings.Builder{}
		for _, str := range strs {
			builder2 := &strings.Builder{}
			builder2.WriteString("[")
			sstr := strings.Split(str, ",")
			for _, v := range sstr {
				if v == "" {
					v = ""
				}
				builder2.WriteString(v)
				builder2.WriteString(",")
			}
			ss := strings.TrimRight(builder2.String(), ",")
			builder3.WriteString(ss)
			builder3.WriteString("]")
			builder3.WriteString(",")
		}
		s := strings.TrimRight(builder3.String(), ",")
		builder.WriteString(s)

	}
	//builder.WriteString(s)
	builder.WriteString("]")
	return builder.String()
}

func convertType(typ string) string {
	switch typ {
	case intTyp:
		return "int32"
	case int64Typ:
		return "int64"
	case float32Typ:
		return "float32"
	case float64Typ:
		return "float64"
	case repeatedIntTyp:
		return "[]int32"
	case repeatedInt32Typ:
		return "[]int32"
	case repeatedInt64Typ:
		return "[]int64"
	case repeatedStringTyp:
		return "[]string"
	case repeatedFloat32Typ:
		return "[]float32"
	case repeatedFloat64Typ:
		return "[]float64"
	case repeatedLongTyp:
		return "[]int64"
	//---------
	case repeatedInt2Typ:
		return "[][]int32"
	case repeatedString2Typ:
		return "[][]string"
	case repeatedFloat2Typ:
		return "[][]float32"
	case repeatedDouble2Typ:
		return "[][]float64"
	case repeatedLong2Typ:
		return "[][]int64"
	//----------
	case mapInt32String:
		return "map[int32]string"
	case mapInt32In32:
		return "map[int32]int32"
	case mapInt64String:
		return "map[int64]string"
	case mapInt64Int32:
		return "map[int64]int32"
	case mapInt64Int64:
		return "map[int64]int64"
	case mapInt32Float32:
		return "map[int32]float32"
	case mapInt32Float64:
		return "map[int32]float64"
	case mapStringString:
		return "map[string]string"
	case mapStringInt32:
		return "map[string]int32"
	case mapIntFloat32:
		return "map[int32]float32"
	case mapIntInt:
		return "map[int32]int32"
	case mapLongLong:
		return "map[int64]int64"
	case mapLongInt:
		return "map[int64]int32"
	case mapIntString:
		return "map[int32]string"

	}
	return typ
}

func IsMap(typ string) bool {
	switch typ {
	case mapInt32String, mapInt32In32, mapInt64String, mapInt64Int32,
		mapInt64Int64, mapInt32Float32, mapInt32Float64, mapStringString, mapStringInt32, mapIntInt, mapIntString, mapIntFloat32:
		return true
	default:
		return false
	}
}

func parseMapJson(typ, val string) string {
	b1 := &strings.Builder{}
	b1.WriteString("{")
	if val == "" {
		b1.WriteString("}")
		return b1.String()
	}
	str1 := strings.Split(val, ",")
	b2 := &strings.Builder{}
	for _, s := range str1 {
		str2 := strings.Split(s, ":")
		key := str2[0]
		v := str2[1]
		b2.WriteString("\"")
		b2.WriteString(key)
		b2.WriteString("\"")
		b2.WriteString(":")
		if typ == mapInt32String || typ == mapInt64String || typ == mapStringString || typ == mapIntString {
			b2.WriteString("\"")
			b2.WriteString(v)
			b2.WriteString("\"")
		} else {
			b2.WriteString(v)
		}
		b2.WriteString(",")
	}
	ss := strings.TrimRight(b2.String(), ",")
	b1.WriteString(ss)
	b1.WriteString("}")
	return b1.String()
}
