package gui

import (
	"image/color"
	guir "proxclient/internal/gui/register"
	gui "proxclient/internal/gui/settings"
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

// Renders the lading page
func (r *RenderLanding) RenderLandingWindow(a fyne.App, name string, service_status bool, server_status bool, config *settings.Configs, accounts []settings.Account) fyne.Window {

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

	yAxisAuxliar := float32(10)

	label_name := canvas.NewText(name, color.Black)
	label_name.Resize(fyne.NewSize(480, 30))
	label_name.TextSize = 24
	label_name.Alignment = fyne.TextAlignCenter
	label_name.Move(fyne.NewPos(10, yAxisAuxliar))
	yAxisAuxliar = yAxisAuxliar + label_name.MinSize().Height + 10

	service_icon_status := canvas.NewCircle(color.NRGBA{R: 0, G: 255, B: 0, A: 255})
	service_icon_status.Resize(fyne.NewSize(12, 12))
	service_icon_status.Move(fyne.NewPos(10, yAxisAuxliar+10)) // "+10" to fit circle behavior

	service_label_status := widget.NewLabel("Prox Service running")
	service_label_status.Resize(fyne.NewSize(240, 30))
	service_label_status.Move(fyne.NewPos(18, yAxisAuxliar))

	if !service_status {
		service_icon_status.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		service_label_status.SetText("Prox Service not running")
	}

	server_icon_status := canvas.NewCircle(color.NRGBA{R: 0, G: 255, B: 0, A: 255})
	server_icon_status.Resize(fyne.NewSize(12, 12))
	server_icon_status.Move(fyne.NewPos(220, yAxisAuxliar+10)) // "+10" to fit circle behavior

	server_label_status := widget.NewLabel("Prox server running")
	server_label_status.Resize(fyne.NewSize(100, 30))
	server_label_status.Move(fyne.NewPos(228, yAxisAuxliar))

	if !server_status {
		server_icon_status.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		server_label_status.SetText("Prox server not running")
	}

	btn_options := widget.NewButton("", func() {
		s := gui.RenderSettings{}
		if settingsWindow != nil {
			settingsWindow.Close()
		}

		settingsWindow = s.RenderSettingsWindow(config, a)
	})
	btn_options.Icon = theme.SettingsIcon()
	btn_options.Move(fyne.NewPos(424, yAxisAuxliar))
	btn_options.Resize(fyne.NewSize(50, 40))
	btn_options.Importance = widget.LowImportance
	btn_options.Refresh()
	yAxisAuxliar = yAxisAuxliar + service_label_status.MinSize().Height + 10

	label_accounts := canvas.NewText("Clientes", color.Black)
	label_accounts.TextSize = 18
	label_accounts.Move(fyne.NewPos(17, yAxisAuxliar))

	btn_new_account := widget.NewButton("Novo", func() {
		r := guir.RenderRegister{}
		if registerWindow != nil {
			registerWindow.Close()
		}
		registerWindow = r.RenderNewAccountWindow(config, accounts, a, list)
	})
	btn_new_account.Resize(fyne.NewSize(50, 30))
	btn_new_account.Move(fyne.NewPos(424, yAxisAuxliar))
	btn_new_account.Importance = widget.LowImportance
	btn_new_account.Refresh()
	yAxisAuxliar = yAxisAuxliar + btn_new_account.MinSize().Height + 10

	container_accordion := container.NewVScroll(
		list,
	)
	container_accordion.Resize(fyne.NewSize(464, 330))
	container_accordion.Move(fyne.NewPos(10, yAxisAuxliar))
	//heightAux = heightAux + accordion_accounts.MinSize().Height + 10

	wrapperContainer := container.NewWithoutLayout(
		label_name,
		service_icon_status,
		service_label_status,
		server_icon_status,
		server_label_status,
		btn_options,
		label_accounts,
		btn_new_account,
		container_accordion,
	)

	w.SetContent(wrapperContainer)

	w.Show()

	return w

}
