package gui

import (
	"proxclient/internal/settings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type RenderRegister struct {
}

func (r *RenderRegister) RenderNewAccountWindow(config *settings.Configs, accounts []settings.Account, a fyne.App, list *widget.List) fyne.Window {
	var AuthorizedAccounts []string

	registerWindow := a.NewWindow("Registrar Cliente")
	registerWindow.Resize(fyne.NewSize(400, 300))
	registerWindow.CenterOnScreen()

	for _, c := range accounts {
		if config.Contains(c) {
			continue
		} else {
			AuthorizedAccounts = append(AuthorizedAccounts, c.Name)
		}
	}

	heightAux := float32(20)

	label_select_company := widget.NewLabel("Cliente")
	label_select_company.TextStyle = fyne.TextStyle{Bold: true}
	select_company := widget.NewSelect(AuthorizedAccounts, func(s string) {
	})
	if len(AuthorizedAccounts) > 0 {
		select_company.Selected = AuthorizedAccounts[0]
	} else {
		select_company.PlaceHolder = "Sem clientes"
		select_company.Disable()
	}
	container_select_company := container.NewVBox(label_select_company, select_company)
	container_select_company.Resize(fyne.NewSize(356, 0))
	container_select_company.Move(fyne.NewPos(10, heightAux))
	heightAux = heightAux + container_select_company.MinSize().Height + 10

	label_entry_portalLogin := widget.NewLabel("Chave de acesso")

	label_entry_portalLogin.TextStyle = fyne.TextStyle{Bold: true}
	entry_portalLogin := widget.NewPasswordEntry()
	feedback_entry_portalLogin := widget.NewLabel("")
	container_entry_portalLogin := container.NewVBox(label_entry_portalLogin, entry_portalLogin, feedback_entry_portalLogin)
	container_entry_portalLogin.Resize(fyne.NewSize(356, 0))
	container_entry_portalLogin.Move(fyne.NewPos(10, heightAux))
	heightAux = heightAux + container_entry_portalLogin.MinSize().Height + 10

	btn_submit := widget.NewButton("Resgistrar", func() {
		newCompany := &settings.Account{
			Name:       select_company.Selected,
			PortaLogin: entry_portalLogin.Text,
			CreatedOn:  time.Now().UTC().String(),
			UpdatedOn:  time.Now().UTC().String(),
		}
		config.AddAccount(*newCompany)
		config.Save()
		list.Refresh()

		registerWindow.Close()
	})
	btn_submit.Resize(fyne.NewSize(356, 40))
	btn_submit.Move(fyne.NewPos(10, heightAux))
	//heightAux = heightAux + btn_submit.MinSize().Height +10

	registerWindow.SetContent(
		container.NewPadded(
			container.NewWithoutLayout(
				container_select_company,
				container_entry_portalLogin,
				btn_submit,
			),
		),
	)

	registerWindow.Show()

	return registerWindow
}
