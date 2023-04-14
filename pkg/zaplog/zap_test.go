package zaplog

import (
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	InitLog("./info.log", "./error.log", zap.DebugLevel)
	defer SugaredLogger.Sync()
	SugaredLogger.Infof("测试测试%d", 1)
	SugaredLogger.Errorf("测试测试%d", 2)
	SugaredLogger.Warnf("测试测试%d", 3)

}
