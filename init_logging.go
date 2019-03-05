package medutils

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func configureLogging(loghandler io.Writer, debugMode bool) {
	var logLevel log.Level

	//log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(loghandler)

	switch viper.GetString("app.logLevel") {
	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	case "error":
		logLevel = log.ErrorLevel
	default:
		logLevel = log.InfoLevel
	}
	log.Infof("loglevel: %s", logLevel)
	log.SetLevel(logLevel)
}

func InitLog(fileLocation string) {
	//fileLocation := viper.GetString("app.logPath")
	file, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln("FailedToOpenLogFile", fileLocation, ":", err)
		return
	}
	configureLogging(file, viper.GetBool("app.logLevel"))

	log.Infof("Initialised Logrus")
}
