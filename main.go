package main

import (
	"fmt"
	guis "proxclient/internal/gui/startApp"
	"proxclient/internal/settings"
	"time"

	"fyne.io/fyne/v2/app"
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

func main() {

	a := app.New()

	if settings.ValidateConfigs() {
		err := config.RecieveStoragedData()
		if err != nil {
			fmt.Println(err)
		}
		guis.StartApp(a, config, accounts)

	} else {
		guis.StartApp(a, config, accounts)
	}

}
