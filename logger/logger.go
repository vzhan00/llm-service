package logger

import (
	"os"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
	LogFormatter := &logrus.TextFormatter{
		ForceColors: true,
		FullTimestamp: true,
	}
	Log.SetFormatter(LogFormatter)
}
