package field

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	StringTyp = "string"

	IntTyp   = "int"
	Int32Typ = "int32"
	Int64Typ = "long"

	Float32Typ = "float"
	Float64Typ = "double"

	RepeatedStringTyp  = "string[]"
	RepeatedIntTyp     = "int[]"
	RepeatedInt32Typ   = "int32[]"
	RepeatedInt64Typ   = "int64[]"
	RepeatedFloat32Typ = "float[]"
	RepeatedFloat64Typ = "double[]"
	RepeatedLongTyp    = "long[]"

	RepeatedInt2Typ    = "int[][]"
	RepeatedString2Typ = "string[][]"
	RepeatedFloat2Typ  = "float[][]"
	RepeatedDouble2Typ = "double[][]"
	RepeatedLong2Typ   = "long[][]"

	MapInt32String  = "map<int32,string>"
	MapInt32In32    = "map<int32,int32>"
	MapInt64String  = "map<int64,string>"
	MapInt64Int32   = "map<int64,int32>"
	MapInt64Int64   = "map<int64,int64>"
	MapStringString = "map<string,string>"
	MapStringInt32  = "map<string,int32>"
	MapIntInt       = "map<int,int>"
	MapLongLong     = "map<long,long>"
	MapLongInt      = "map<long,int>"
	MapIntString    = "map<int,string>"
	MapInt32Float32 = "map<int32,float>"
	MapIntFloat32   = "map<int,float>"
	MapInt32Float64 = "map<int32,double>"
)

func CheckFieldType(t string) bool {
	switch t {
	case StringTyp:
		return true
	case IntTyp, Int32Typ, Int64Typ:
		return true
	case Float32Typ, Float64Typ:
		return true
	case RepeatedStringTyp, RepeatedIntTyp, RepeatedInt32Typ, RepeatedInt64Typ,
		RepeatedFloat32Typ, RepeatedFloat64Typ, RepeatedLongTyp, RepeatedLong2Typ:
		return true
	case MapInt32String, MapInt32In32, MapInt64String, MapInt64Int32,
		MapInt64Int64, MapInt32Float32, MapInt32Float64, MapStringString,
		MapStringInt32, MapIntInt, MapIntString, MapIntFloat32, MapLongLong, MapLongInt:
		return true
	case RepeatedInt2Typ, RepeatedString2Typ, RepeatedFloat2Typ, RepeatedDouble2Typ:
		return true
	default:
		return false
	}
}

func IsString(typ string) bool {
	return typ == StringTyp
}

func ParseString(val string) string {
	if val == "" {
		return "\"\""
	} else {
		b, _ := json.Marshal(val)
		return string(b)
	}
}

func IsInt(typ string) bool {
	switch typ {
	case IntTyp, Int32Typ, Int64Typ:
		return true
	default:
		return false
	}
}

func ParseInt(typ, val string) (int64, error) {
	switch typ {
	case IntTyp, Int32Typ:
		return strconv.ParseInt(val, 10, 32)
	case Int64Typ:
		return strconv.ParseInt(val, 10, 64)
	default:
		return 0, errors.New("parseInt err")
	}
}

func IsFloat(typ string) bool {
	switch typ {
	case Float32Typ:
	case Float64Typ:
	default:
		return false
	}
	return true
}

func ParseFloat(typ, val string) (float64, error) {
	switch typ {
	case Float32Typ:
		return strconv.ParseFloat(val, 32)
	case Float64Typ:
		return strconv.ParseFloat(val, 64)
	default:
		return 0, errors.New("parseFloat err")
	}
}

// IsRepeated 1,2,3,4
func IsRepeated(typ string) bool {
	if typ == "" || !strings.Contains(typ, "[]") {
		return false
	}
	switch typ {
	case RepeatedFloat64Typ, RepeatedStringTyp, RepeatedIntTyp, RepeatedInt32Typ,
		RepeatedInt64Typ, RepeatedFloat32Typ, RepeatedLongTyp:
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
	case RepeatedInt2Typ, RepeatedString2Typ, RepeatedFloat2Typ, RepeatedDouble2Typ, RepeatedLong2Typ:
		return true
	default:
		return false
	}
}

func ParseRepeatedJson(typ, val string) string {
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
	if typ == RepeatedStringTyp {
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

func ParseRepeated2Json(typ, val string) string {
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
	if typ == RepeatedString2Typ {
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

func ConvertType(typ string) string {
	switch typ {
	case IntTyp:
		return "int32"
	case Int64Typ:
		return "int64"
	case Float32Typ:
		return "float32"
	case Float64Typ:
		return "float64"
	case RepeatedIntTyp:
		return "[]int32"
	case RepeatedInt32Typ:
		return "[]int32"
	case RepeatedInt64Typ:
		return "[]int64"
	case RepeatedStringTyp:
		return "[]string"
	case RepeatedFloat32Typ:
		return "[]float32"
	case RepeatedFloat64Typ:
		return "[]float64"
	case RepeatedLongTyp:
		return "[]int64"
	//---------
	case RepeatedInt2Typ:
		return "[][]int32"
	case RepeatedString2Typ:
		return "[][]string"
	case RepeatedFloat2Typ:
		return "[][]float32"
	case RepeatedDouble2Typ:
		return "[][]float64"
	case RepeatedLong2Typ:
		return "[][]int64"
	//----------
	case MapInt32String:
		return "map[int32]string"
	case MapInt32In32:
		return "map[int32]int32"
	case MapInt64String:
		return "map[int64]string"
	case MapInt64Int32:
		return "map[int64]int32"
	case MapInt64Int64:
		return "map[int64]int64"
	case MapInt32Float32:
		return "map[int32]float32"
	case MapInt32Float64:
		return "map[int32]float64"
	case MapStringString:
		return "map[string]string"
	case MapStringInt32:
		return "map[string]int32"
	case MapIntFloat32:
		return "map[int32]float32"
	case MapIntInt:
		return "map[int32]int32"
	case MapLongLong:
		return "map[int64]int64"
	case MapLongInt:
		return "map[int64]int32"
	case MapIntString:
		return "map[int32]string"

	}
	return typ
}

func IsMap(typ string) bool {
	switch typ {
	case MapInt32String, MapInt32In32, MapInt64String, MapInt64Int32,
		MapInt64Int64, MapInt32Float32, MapInt32Float64, MapStringString, MapStringInt32, MapIntInt, MapIntString, MapIntFloat32:
		return true
	default:
		return false
	}
}

func ParseMapJson(typ, val string) string {
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
		if typ == MapInt32String || typ == MapInt64String || typ == MapStringString || typ == MapIntString {
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

func ParseMetaField(rowData string, typ string, isPri bool) (string, error) {
	if IsString(typ) {
		if isPri {
			return rowData, nil
		}
		return ParseString(rowData), nil
	} else if IsInt(typ) {
		if rowData == "" {
			return "0", nil
		}
		v, err := ParseInt(typ, rowData)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", v), nil
	} else if IsFloat(typ) {
		if rowData == "" {
			return "0", nil
		}
		v, err := ParseFloat(typ, rowData)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%f", v), nil
	} else if IsRepeated(typ) {
		return ParseRepeatedJson(typ, rowData), nil
	} else if IsRepeated2(typ) {
		return ParseRepeated2Json(typ, rowData), nil
	} else if IsMap(typ) {
		return ParseMapJson(typ, rowData), nil
	}

	return "", errors.New("parseInt err")
}
