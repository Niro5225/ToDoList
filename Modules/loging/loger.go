package loging

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type Loging struct {
	Loger *logrus.Logger
}

func NewLoging(filename string) Loging {
	os.Mkdir("logs", os.ModePerm)
	var loger = logrus.New()
	loger.Formatter = new(logrus.JSONFormatter)
	loger.Formatter = new(logrus.TextFormatter)
	loger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	file, err := os.OpenFile(fmt.Sprintf("logs/%s", filename), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		loger.Out = file
	} else {
		loger.Info("Failed to log to file, using default stderr")
	}
	return Loging{Loger: loger}
}
