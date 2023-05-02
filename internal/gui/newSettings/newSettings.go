package gui

import (
	"fmt"
	"image/color"
	"proxclient/internal/settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

var accounts = []settings.Account{
	{Name: "Empresa_1", PortaLogin: "axon"},
	{Name: "Empresa_2", PortaLogin: "axon"},
	{Name: "Empresa_3", PortaLogin: "axon"},
	{Name: "Empresa_4", PortaLogin: "axon"},
	{Name: "Empresa_5", PortaLogin: "axon"},
}

type RenderNewSettings struct{}

func (r *RenderNewSettings) RenderNewSettingsWindow(config *settings.Configs, a fyne.App, key string) fyne.Window {
	w := a.NewWindow("Configurações iniciais")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(500, 500))
	w.CenterOnScreen()

	heightAux := float32(20)

	key_label := canvas.NewText("Chave", color.Black)
	key_label.TextSize = 12
	key_label.Alignment = fyne.TextAlignLeading
	key_label.Resize(fyne.NewSize(470, 40))
	key_label.Move(fyne.NewPos(10, heightAux))

	key_entry := widget.NewEntry()
	key_entry.Resize(fyne.NewSize(470, 40))
	key_entry.Move(fyne.NewPos(10, heightAux))

	key_error := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	key_error.TextSize = 10
	key_error.Alignment = fyne.TextAlignLeading
	key_error.Resize(fyne.NewSize(470, 40))
	key_error.Move(fyne.NewPos(10, heightAux))

	if key != "" {
		key_entry.SetText(key)
		key_entry.Disable()
		key_entry.Refresh()
	}

	heightAux = heightAux + key_error.MinSize().Height + 30

	inboundPath_label := canvas.NewText("Diretório para recebimento de arquivos", color.Black)
	inboundPath_label.Resize(fyne.NewSize(480, 20))
	inboundPath_label.TextSize = 10
	inboundPath_label.Move(fyne.NewPos(24, heightAux+10))
	heightAux = heightAux + inboundPath_label.MinSize().Height + 10

	inboundPath_display := widget.NewLabel("...")
	inboundPath_display.Resize(fyne.NewSize(470, 40))
	inboundPath_display.Move(fyne.NewPos(10, heightAux))

	inboundPath_btn := widget.NewButton("", func() {
		dir, err := dialog.Directory().Title("Selecione um Diretório").Browse()
		if err != nil {
			fmt.Println(err)
		}

		if dir != "" {
			inboundPath_display.SetText(dir)
			inboundPath_display.Refresh()
			config.InboundPath = dir
		}

	})
	inboundPath_btn.Icon = theme.FolderIcon()
	inboundPath_btn.Resize(fyne.NewSize(40, 40))
	inboundPath_btn.Move(fyne.NewPos(440, heightAux))
	inboundPath_btn.Importance = widget.LowImportance
	inboundPath_btn.Refresh()
	heightAux = heightAux + 30

	inboundPath_error := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	inboundPath_error.TextSize = 10
	inboundPath_error.Alignment = fyne.TextAlignLeading
	inboundPath_error.Resize(fyne.NewSize(470, 10))
	inboundPath_error.Move(fyne.NewPos(24, heightAux))
	heightAux = heightAux + inboundPath_error.MinSize().Height + 10

	backupPath_label := canvas.NewText("Diretório para backup de arquivos", color.Black)
	backupPath_label.Resize(fyne.NewSize(480, 20))
	backupPath_label.TextSize = 10
	backupPath_label.Move(fyne.NewPos(24, heightAux+10))
	heightAux = heightAux + backupPath_label.MinSize().Height + 10

	backupPath_display := widget.NewLabel("...")
	backupPath_display.Resize(fyne.NewSize(470, 40))
	backupPath_display.Move(fyne.NewPos(10, heightAux))

	backupPath_btn := widget.NewButton("", func() {
		dir, err := dialog.Directory().Title("Selecione um Diretório").Browse()
		if err != nil {
			fmt.Println(err)
		}

		if dir != "" {
			backupPath_display.SetText(dir)
			backupPath_display.Refresh()
			config.BackupPath = dir
		}

	})
	backupPath_btn.Icon = theme.FolderIcon()
	backupPath_btn.Resize(fyne.NewSize(40, 40))
	backupPath_btn.Move(fyne.NewPos(440, heightAux))
	backupPath_btn.Importance = widget.LowImportance
	backupPath_btn.Refresh()
	heightAux = heightAux + 30

	backupPath_error := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	backupPath_error.TextSize = 10
	backupPath_error.Alignment = fyne.TextAlignLeading
	backupPath_error.Resize(fyne.NewSize(470, 10))
	backupPath_error.Move(fyne.NewPos(24, heightAux))
	heightAux = heightAux + backupPath_error.MinSize().Height + 10

	report_check := widget.NewCheck("Relatório", func(b bool) {
		fmt.Println(config.ReceiveReport)
		config.ReceiveReport = b
		fmt.Println(config.ReceiveReport)
	})
	report_check.Resize(fyne.NewSize(470, 30))
	report_check.Move(fyne.NewPos(10, heightAux))
	heightAux = heightAux + report_check.MinSize().Height + 10

	save_btn := widget.NewButton("Salvar", func() {
		isValid := true
		if inboundPath_display.Text == "..." {
			isValid = false
			inboundPath_error.Text = "* Campo obrigatório"
		}

		if backupPath_display.Text == "..." {
			isValid = false
			backupPath_error.Text = "* Campo obrigatório"
		}

		if isValid {
			config.Save()

			w.Close()

		}

	})
	save_btn.Resize(fyne.NewSize(100, 40))
	save_btn.Move(fyne.NewPos(380, 440))
	save_btn.Importance = widget.HighImportance
	save_btn.Refresh()

	wrapperContainer := container.NewWithoutLayout(
		key_label,
		key_entry,
		key_error,
		inboundPath_label,
		inboundPath_display,
		inboundPath_btn,
		inboundPath_error,
		backupPath_label,
		backupPath_display,
		backupPath_btn,
		backupPath_error,
		report_check,
		save_btn,
	)

	w.SetContent(wrapperContainer)

	w.Show()

	return w

}

func (r *RenderNewSettings) GetSettings(key string, config *settings.Configs, a fyne.App) error {
	//Get existing configs from database using client unique key
	// resp, err := http.Get("....")
	// if err != nil {
	// 	return nil
	// }

	// data, err := ioutil.ReadAll(resp.Body)
	// if err := json.Unmarshal(data, config); err != nil {
	// 	fmt.Println("Can not unmarshal JSON")
	// 	return err
	// }

	config.Key = key //simulate user found

	config.Create()
	config.Save()

	r.RenderNewSettingsWindow(config, a, key)

	return nil
}

func (r *RenderNewSettings) RenderValidateWindow(config *settings.Configs, a fyne.App) fyne.Window {
	w := a.NewWindow("Ativar/Buscar")
	w.Resize(fyne.NewSize(400, 160))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	heightAux := float32(20)

	key_label := canvas.NewText("Chave do produto", color.Black)
	key_label.TextSize = 18
	key_label.Alignment = fyne.TextAlignLeading
	key_label.Resize(fyne.NewSize(370, 40))
	key_label.Move(fyne.NewPos(10, heightAux))
	heightAux = heightAux + key_label.MinSize().Height + 10

	key_entry := widget.NewEntry()
	key_entry.Resize(fyne.NewSize(370, 40))
	key_entry.Move(fyne.NewPos(10, heightAux))

	key_error := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	key_error.TextSize = 10
	key_error.Alignment = fyne.TextAlignLeading
	key_error.Resize(fyne.NewSize(370, 40))
	key_error.Move(fyne.NewPos(10, heightAux))
	heightAux = heightAux + key_entry.MinSize().Height + 14

	save_btn := widget.NewButton("Salvar", func() {
		r.GetSettings(key_entry.Text, config, a)
		w.Close()
	})
	save_btn.Resize(fyne.NewSize(100, 40))
	save_btn.Importance = widget.HighImportance
	save_btn.Move(fyne.NewPos(280, heightAux))
	save_btn.Resize(fyne.NewSize(100, 40))
	save_btn.Refresh()

	wrapperContainer := container.NewWithoutLayout(
		key_label,
		key_entry,
		key_error,
		save_btn,
	)

	w.SetContent(wrapperContainer)

	w.Show()

	return w
}
