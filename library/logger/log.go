package logger

import (
	"go.uber.org/zap"
	"encoding/json"
)

type BLog struct {
	logger *zap.Logger
}

var log BLog

func init() {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"common key": "common value"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	log = BLog{logger: logger}
	//log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
}
func GetLogger() BLog {
	return log
}

func (log *BLog) Debug(args ...interface{}) {
	log.logger.Sugar().Debug(args)
}

func (log *BLog) Debugf(format string, args ...interface{}) {
	log.logger.Sugar().Debugf(format, args)
}

func (log *BLog) Info(args ...interface{}) {
	log.logger.Sugar().Info(args)
}
func (log *BLog) Infof(format string, args ...interface{}) {
	log.logger.Sugar().Infof(format, args)
}
func (log *BLog) Warning(args ...interface{}) {
	log.logger.Sugar().Warn(args)
}

func (log *BLog) Warningf(format string, args ...interface{}) {
	log.logger.Sugar().Warnf(format, args)
}

func (log *BLog) Error(args ...interface{}) {
	log.logger.Sugar().Error(args)
}

func (log *BLog) Errorf(format string, args ...interface{}) {
	log.logger.Sugar().Errorf(format, args)
}

func (log *BLog) Fatal(args ...interface{}) {
	log.logger.Sugar().Fatal(args)
}

func (log *BLog) Fatalf(format string, args ...interface{}) {
	log.logger.Sugar().Fatalf(format, args)
}

//var Logger *zap.Logger
//
//func init() {
//	logPath := "./app.log"
//	logLevel := "info"
//	js := fmt.Sprintf(`{
//      "level": "%s",
//      "encoding": "json",
//      "outputPaths": ["%s"],
//      "errorOutputPaths": ["%s"]
//      }`, logLevel, logPath, logPath)
//	var cfg zap.Config
//	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
//		panic(err)
//	}
//	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
//	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	Logger, err := cfg.Build()
//	if err != nil {
//		log.Fatal("init logger error:", err)
//	}
//	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
//}
