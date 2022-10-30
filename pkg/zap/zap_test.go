// @author moqi
// On 2022/10/30 18:32:24
package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
	"time"
)

const url = "https://www.uber.com"

func TestSugaredLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger) // flushes buffer, if any

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

func TestZapLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestCreateLogger(t *testing.T) {
	// 初始化logger
	logger := zap.NewExample()
	// 使用defer logger.Sync()将缓存同步到文件中。
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	// 记录日志
	logger.Info("NewExample",
		zap.String("name", "张三"),
		zap.Int("age", 18),
	)
	productionLogger, _ := zap.NewProduction()
	defer func(productionLogger *zap.Logger) {
		_ = productionLogger.Sync()
	}(productionLogger)
	productionLogger.Info("NewProduction",
		zap.String("name", "张三"),
		zap.Int("age", 18),
	)
	devLogger, _ := zap.NewDevelopment()
	defer func(devLogger *zap.Logger) {
		_ = devLogger.Sync()
	}(devLogger)
	devLogger.Info("NewDevelopment",
		zap.String("name", "张三"),
		zap.Int("age", 18),
	)
}

// 使用默认记录日志
func TestRecordLogWithDefault(t *testing.T) {
	// 初始化记录器（使用默认记录器）
	logger := zap.NewExample()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	// 记录日志
	logger.Debug("这是debug日志")
	logger.Debug("这是debug日志", zap.String("name", "张三"))
	logger.Info("这是info日志", zap.Int("age", 18))
	logger.Error("这是error日志", zap.Int("line", 130), zap.Error(fmt.Errorf("错误示例")))
	logger.Warn("这是Warn日志")
	// 下面两个都会中断程序
	//logger.Fatal("这是Fatal日志")
	//logger.Panic("这是Panic日志")
}

// 使用Sugar记录器
func TestRecordLogWithSugar(t *testing.T) {
	// 初始化记录器
	logger := zap.NewExample()
	// 把日志记录器转成Sugar
	sugarLogger := logger.Sugar()
	defer func(sugarLogger *zap.SugaredLogger) {
		_ = sugarLogger.Sync()
	}(sugarLogger)
	// 记录日志
	sugarLogger.Debug("这是debug日志 ", zap.String("name", "张三"))
	sugarLogger.Debugf("这是Debugf日志 name:%s ", "张三")
	sugarLogger.Info("这是info日志", zap.Int("age", 18))
	sugarLogger.Infof("这是Infof日志  内容:%v", map[string]string{"爱好": "动漫"})
	sugarLogger.Error("这是error日志", zap.Int("line", 130), zap.Error(fmt.Errorf("错误示例")))
	sugarLogger.Errorf("这是Errorf日志，错误信息：%s", "错误报告！")
	sugarLogger.Warn("这是Warn日志")
	sugarLogger.Warnf("这是Warnf日志 %v", []int{1, 2, 4, 5})
	// 下面两个都会中断程序
	//sugarLogger.Fatal("这是Fatal日志")
	//sugarLogger.Panic("这是Panic日志")
}

func Test2File(t *testing.T) {
	// 指定写入文件
	fileHandle, _ := os.Create("./test.log")
	writeFile := zapcore.AddSync(fileHandle)
	// 设置日志输出格式为JSON (参数复用NewDevelopmentEncoderConfig)
	encoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	// 返回zapcore.Core,并指定记录zap.DebugLevel级别及以上日志
	zcore := zapcore.NewCore(encoder, zapcore.Lock(writeFile), zap.DebugLevel)
	// 创建日志记录器
	logger := zap.New(zcore)
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	// 记录日志
	logger.Info("输出日志到文件", zap.String("name", "张三"))
}
