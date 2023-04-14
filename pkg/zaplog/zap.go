/*
@Author: nullzz
@Date: 2022/3/23 7:37 下午
@Version: 1.0
@DEC:
*/
package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var SugaredLogger *zap.SugaredLogger
var Logger *zap.Logger

func Init(logLevel zapcore.Level) {
	config := zapcore.EncoderConfig{
		MessageKey:   "msg",                            //结构化（json）输出：msg的key
		LevelKey:     "level",                          //结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
		TimeKey:      "ts",                             //结构化（json）输出：时间的key（INFO，WARN，ERROR等）
		CallerKey:    "file",                           //结构化（json）输出：打印日志的文件对应的Key
		EncodeLevel:  zapcore.CapitalColorLevelEncoder, //将日志级别转换成大写（INFO，WARN，ERROR等）
		EncodeCaller: zapcore.ShortCallerEncoder,       //采用短文件路径编码输出（test/main.go:14 ）
		//EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		//	enc.AppendString(t.Format("2006-01-02 15:04:05"))
		//},//输出的时间格式
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		}, //
	}
	// 实现多个输出
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel), //同时将日志输出到控制台，
		//NewJSONEncoder 是结构化输出
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	SugaredLogger = Logger.Sugar()
}

// InitLog 初始化日志 logger
func InitLog(logPath, errPath string, logLevel zapcore.Level) {
	config := zapcore.EncoderConfig{
		MessageKey:   "msg",                            //结构化（json）输出：msg的key
		LevelKey:     "level",                          //结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
		TimeKey:      "ts",                             //结构化（json）输出：时间的key（INFO，WARN，ERROR等）
		CallerKey:    "file",                           //结构化（json）输出：打印日志的文件对应的Key
		EncodeLevel:  zapcore.CapitalColorLevelEncoder, //将日志级别转换成大写（INFO，WARN，ERROR等）
		EncodeCaller: zapcore.ShortCallerEncoder,       //采用短文件路径编码输出（test/main.go:14 ）
		//EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		//	enc.AppendString(t.Format("2006-01-02 15:04:05"))
		//},//输出的时间格式
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		}, //
	}
	//自定义日志级别：自定义Info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})
	//自定义日志级别：自定义Warn级别
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})
	// 获取io.Writer的实现
	infoWriter := getWriter(logPath)
	warnWriter := getWriter(errPath)
	// 实现多个输出
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(infoWriter), infoLevel),                            //将info及以下写入logPath，NewConsoleEncoder 是非结构化输出
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(warnWriter), warnLevel),                            //warn及以上写入errPath
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel), //同时将日志输出到控制台，
		//NewJSONEncoder 是结构化输出
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	SugaredLogger = Logger.Sugar()
}

//func Init(prod bool, level zapcore.Level) {
//	var cfg zapcore.EncoderConfig
//	if prod {
//		cfg = zap.NewProductionEncoderConfig()
//	} else {
//		cfg = zap.NewDevelopmentEncoderConfig()
//	}
//	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
//	cfg.EncodeLevel = zapcore.LowercaseColorLevelEncoder //这里可以指定颜色
//
//	InitByConfig(cfg, level)
//}

//func InitByConfig(cfg zapcore.EncoderConfig, level zapcore.Level) {
//	writer := zapcore.AddSync(os.Stdout)
//	encoder := zapcore.NewConsoleEncoder(cfg)
//	core := zapcore.NewCore(encoder, writer, level)
//	Logger = zap.New(core, zap.AddCaller())
//	SugaredLogger = Logger.Sugar()
//}

// defer
func Sync() {
	Logger.Sync()
	SugaredLogger.Sync()
}
