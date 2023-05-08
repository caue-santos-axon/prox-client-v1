package main

import (
	"os"
	guis "proxclient/internal/gui/startApp"
	"proxclient/internal/logging"
	"proxclient/internal/mutex"
	"proxclient/internal/setboot"
	"proxclient/internal/settings"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/sirupsen/logrus"
)

var config = &settings.Configs{
	InboundPath: "",
	BackupPath:  "none",
	CreatedOn:   time.Now().UTC().String(),
	UpdatedOn:   time.Now().UTC().String(),
}

var accounts = []settings.Account{
	{Name: "Empresa_1", PortaLogin: "axon"},
	{Name: "Empresa_2", PortaLogin: "axon"},
	{Name: "Empresa_3", PortaLogin: "axon"},
	{Name: "Empresa_4", PortaLogin: "axon"},
	{Name: "Empresa_5", PortaLogin: "axon"},
}

func init() {
	err := setboot.SetStartWithOS()
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Coundn't register app to start with OS")
	}

	_, err = mutex.CreateMutex("Global\\proxGUI")
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"msg": err,
		}).Error("Init error")
		os.Exit(0)
	}
}

func main() {

	a := app.New()

	if settings.ValidateConfigs() {
		err := config.RecieveStoragedData()
		if err != nil {
			logging.Log.WithFields(logrus.Fields{
				"err": err,
			}).Error("Coundn't getting storage data")
		}
		guis.StartApp(a, config, accounts)

	} else {
		guis.StartApp(a, config, accounts)
	}

}
