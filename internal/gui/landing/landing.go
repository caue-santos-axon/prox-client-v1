package gui

import (
	"fmt"
	"image/color"
	guir "proxclient/internal/gui/register"
	guis "proxclient/internal/gui/settings"
	"proxclient/internal/settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type RenderLanding struct {
}

var registerWindow fyne.Window
var settingsWindow fyne.Window

func (r *RenderLanding) RenderLandingWindow(a fyne.App, name string, status bool, config *settings.Configs, accounts []settings.Account) fyne.Window {

	w := a.NewWindow("Prox Client")
	w.Resize(fyne.NewSize(500, 500))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	list := widget.NewList(
		func() int {
			return len(config.AuthorizedAccounts)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(config.AuthorizedAccounts[lii].Name)
		},
	)

	heightAux := float32(10)

	label_name := canvas.NewText(name, color.Black)
	label_name.Resize(fyne.NewSize(480, 30))
	label_name.TextSize = 24
	label_name.Alignment = fyne.TextAlignCenter
	label_name.Move(fyne.NewPos(10, heightAux))
	heightAux = heightAux + label_name.MinSize().Height + 10

	label_status := widget.NewLabel("Status: Online")
	if !status {
		label_status.SetText("Status: Offline")
		label_status.Refresh()
	}
	label_status.Resize(fyne.NewSize(240, 30))
	label_status.Move(fyne.NewPos(10, heightAux))

	btn_options := widget.NewButton("", func() {
		s := guis.RenderSettings{}
		if settingsWindow != nil {
			settingsWindow.Close()
		}
		settingsWindow = s.RenderSettingsWindow(a)
		fmt.Println("clickou")
	})
	btn_options.Icon = theme.SettingsIcon()
	btn_options.Move(fyne.NewPos(424, heightAux))
	btn_options.Resize(fyne.NewSize(50, 40))
	heightAux = heightAux + label_status.MinSize().Height + 10

	label_accounts := canvas.NewText("Clientes", color.Black)
	label_accounts.TextSize = 18
	label_accounts.Move(fyne.NewPos(17, heightAux))

	btn_new_account := widget.NewButton("Novo", func() {
		r := guir.RenderRegister{}
		if registerWindow != nil {
			registerWindow.Close()
		}
		registerWindow = r.RenderNewAccountWindow(config, accounts, a, list)
	})
	btn_new_account.Resize(fyne.NewSize(50, 30))
	btn_new_account.Move(fyne.NewPos(424, heightAux))
	btn_new_account.Importance = widget.LowImportance
	heightAux = heightAux + btn_new_account.MinSize().Height + 10

	container_accordion := container.NewVScroll(
		list,
	)
	container_accordion.Resize(fyne.NewSize(464, 330))
	container_accordion.Move(fyne.NewPos(10, heightAux))
	//heightAux = heightAux + accordion_accounts.MinSize().Height + 10

	wrapperContainer := container.NewWithoutLayout(
		label_name,
		label_status,
		btn_options,
		label_accounts,
		btn_new_account,
		container_accordion,
	)

	w.SetContent(wrapperContainer)

	w.Show()

	return w

}
