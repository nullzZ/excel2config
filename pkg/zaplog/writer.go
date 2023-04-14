package zaplog

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,    //最大M数，超过则切割
		MaxBackups: 5,     //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     30,    //保存30天
		Compress:   false, //是否压缩
	}
}

//func getWriter2(filename string) io.Writer {
//	// 生成rotatelogs的Logger 实际生成的文件名 filename.YYmmddHH
//	// filename是指向最新日志的链接
//	hook, err := rotatelogs.New(
//		filename+".%Y%m%d%H",
//		rotatelogs.WithLinkName(filename),
//		rotatelogs.WithMaxAge(time.Hour*24*30),    // 保存30天
//		rotatelogs.WithRotationTime(time.Hour*24), //切割频率 24小时
//	)
//	if err != nil {
//		log.Println("日志启动异常")
//		panic(err)
//	}
//	return hook
//}
