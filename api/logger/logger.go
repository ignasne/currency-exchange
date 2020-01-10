package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

const defaultLogLevel = logrus.DebugLevel

var standardLogger *logrus.Entry

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(defaultLogLevel)

	if envLogLvl, ok := os.LookupEnv("APP_LOG_LEVEL"); ok {
		logLevel, err := logrus.ParseLevel(envLogLvl)
		if err != nil {
			logrus.WithError(err).Fatal("Could not parse log level")
			os.Exit(1)
		}
		logrus.SetLevel(logLevel)
	}

	standardLogger = logrus.StandardLogger().WithFields(logrus.Fields{})
}

func Get() *logrus.Entry {
	return standardLogger
}
