package gui

import (
	guil "proxclient/internal/gui/landing"
	"proxclient/internal/settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var landingWindow fyne.Window

func StartApp(a fyne.App, config *settings.Configs, accounts []settings.Account) {

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Open",
			fyne.NewMenuItem("Abrir", func() {
				r := guil.RenderLanding{}
				if landingWindow != nil {
					landingWindow.Close()
				}
				landingWindow = r.RenderLandingWindow(a, "Prox Client", true, config, accounts)
			},
			))
		desk.SetSystemTrayMenu(m)
	}

	a.Run()

}
