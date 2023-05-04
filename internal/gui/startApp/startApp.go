package gui

import (
	gui "proxclient/internal/gui/landing"
	guins "proxclient/internal/gui/newSettings"
	"proxclient/internal/logging"
	"proxclient/internal/monitor"
	"proxclient/internal/settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/sirupsen/logrus"
)

var landingWindow fyne.Window
var newStttingsWindow fyne.Window

func StartApp(a fyne.App, config *settings.Configs, accounts []settings.Account) {

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Open",
			fyne.NewMenuItem("Abrir", func() {
				if settings.ValidateConfigs() {
					r := gui.RenderLanding{}
					if landingWindow != nil {
						landingWindow.Close()
					}

					status, err := monitor.IsOnline()
					if err != nil {
						logging.Log.WithFields(logrus.Fields{
							"err": err,
						}).Error("Coundn't get service status")
						landingWindow = r.RenderLandingWindow(a, "Prox Client", false, false, config, accounts)
					} else {
						landingWindow = r.RenderLandingWindow(a, "Prox Client", status, false, config, accounts)
					}

				} else {
					ns := guins.RenderNewSettings{}
					if newStttingsWindow != nil {
						newStttingsWindow.Close()
					}
					newStttingsWindow = ns.RenderValidateWindow(config, a)

				}
			},
			))
		desk.SetSystemTrayMenu(m)
	}

	a.Run()

}
