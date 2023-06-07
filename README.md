# excel gen

## 运行
./darwin_gen_excel -s="../data" -t="../gen" -skipRow=5 -skipCol=1

### excel文件规则

* 忽略不是xlsx扩展名的、开头#号的 ，如 忽略a.exl、#a.xlsx

### sheet规则

#### 忽略规则

* 默认名Sheet
* 中文
* #开头
* _c前端使用
* _s或者不带_c、_s为后端使用，_c前端使用

#### 跳过规则
* 跳过行数 skipRow
* 跳过列数 skipCol

### 字段规则

#### 忽略规则

* 没有配置字段属性
* 没有字段名



* 第1行策划描述
* 第2行字段名 如 id
* 第3行字段类型(后端) 如 int ,不配置就忽略此字段
* 第4行字段类型(前端) 如 int
* 第5行备注描述
* map类型 map<int,int>//1:2,3:4
    * 类型：map<int,long>,map<int,string>等
* 一维数组 1,2,3,4或者[1,2,3,4]
    * 类型：int[],long[],string[],float[],double[]等
* 二维数组 [[1,2],[3,4]],单数据可以[[1,2]]或[1,2]
    * 类型：int[][],long[][],string[][],float[][],double[][]等
* 基本类型 int(备注等于int32) int32 long float double string

### 类型

* 主键必须是 int int32 int64 string
* 字段支持的类型

```go

stringTyp = "string"

intTyp = "int"
int32Typ = "int32"
int64Typ = "int64"

float32Typ = "float"
float64Typ = "double"

repeatedStringTyp = "string[]"
repeatedIntTyp = "int[]"
repeatedInt32Typ = "int32[]"
repeatedInt64Typ = "int64[]"
repeatedFloat32Typ = "float[]"
repeatedFloat64Typ = "double[]"

repeatedInt2Typ = "int[][]"
repeatedString2Typ = "string[][]"
repeatedFloat2Typ = "float[][]"
repeatedDouble2Typ = "double[][]"

mapInt32String = "map<int32,string>"
mapInt32In32 = "map<int32,int32>"
mapInt64String = "map<int64,string>"
mapInt64Int32 = "map<int64,int32>"
mapInt64Int64   = "map<int64,int64>"
mapStringString = "map<string,string>"
mapStringInt32 = "map<string,int32>"
mapIntInt = "map<int,int>"
mapIntString = "map<int,string>"
mapInt32Float32 = "map<int32,float>"
mapIntFloat32 = "map<int,float>"
mapInt32Float64 = "map<int32,double>"
```

### 表数据检测

服务器字段后加入@ 例如 int@自定义注解名称

?增加额外参数

例如 map<int,int>@Resource?1,2

### golang 使用

* 生成配置文件 bin中可执行文件 -s excel源地址 -t gen目标地址
* 加载
    ```
    zap_log.Init(false, zap.DebugLevel)
    config.InitWithLoader(conf_loader.AddLoader)
    conf_loader.MustInitLocal("/Users/malei/works/null-kit/excel2conf/gen/rawdata", true, zap_log.Logger)
   ```
* 获取数据 c, ok := config.GetTest("1")
* 另一种获取数据的方式 c,ok:=config.Get("SheetName",id),
    * id需要注意对应的类型
    * 如果索引是int32 c,ok:=config.Get("SheetName",id)
    * 最好不要使用此方法 此方法可用于表检查
* 自定义注解检测字段内数据

```
  m := make(map[string]func(i interface{}, param string) bool)
  m["ItemData"] = func(i interface{}, param string) bool {
      d, ok := i.(map[int32]int32)
      if !ok {
          return false
      }
      for k, _ := range d {
          _, h := config.GetItemData(k)
          if !h {
              return false
          }
      }
      return true
  }
  m["InConf"] = func(i interface{}, param string) bool {
      d, ok := i.(map[int32]int32)
      if !ok {
          return false
      }
      for k, _ := range d {
          _, h := config.Get(param, k)
          if !h {
              return false
          }
      }
      return true
  }
  config.InitCheckerFunc(m)//注册自定义注解函数
  conf_loader.AddChecker(config.Checker)//加载检测方法
  ```

### build

##### mac和linux

* CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o linux_gen_excel main.go
* CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s' -o darwin_gen_excel main.go
* CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o win_gen_excel.exe main.go

##### windows编译

* go build -ldflags '-w -s' main.go .\win_gen.exe -s=C:\null-kit\excel2conf\data -t=C:\null-kit\excel2conf\gen

### 遇到问题，记录一下

### 待开发
