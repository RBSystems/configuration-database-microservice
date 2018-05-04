package log

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.SugaredLogger
var CFG zap.Config

func init() {

	CFG = zap.NewDevelopmentConfig()
	//CFG.OutputPaths = append(CFG.OutputPaths, os.Getenv("LOGGING_FILE_LOCATION"))
	CFG.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	l, err := CFG.Build()
	if err != nil {
		log.Printf("Couldn't build config for zap logger: %v", err.Error())
		panic(err)
	}
	L = l.Sugar()
	L.Info("Zap Logger Started")

}
