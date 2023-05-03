package logging

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("C:\\prox_client.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Info("Failed to log to file, using default stderr, err: %v", err)
	} else {
		fmt.Println("File opened with succes")
		Log.Out = file
	}
}
