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

type RenderSettings struct{}

func (r *RenderSettings) RenderSettingsWindow(config *settings.Configs, a fyne.App) fyne.Window {
	w := a.NewWindow("Prox Settings")
	w.Resize(fyne.NewSize(500, 500))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	heightAux := float32(10)

	label_name := canvas.NewText("Configurações", color.Black)
	label_name.Resize(fyne.NewSize(480, 30))
	label_name.TextSize = 24
	label_name.Alignment = fyne.TextAlignCenter
	label_name.Move(fyne.NewPos(10, heightAux))
	heightAux = heightAux + label_name.MinSize().Height + 20

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
	heightAux = heightAux + inboundPath_display.MinSize().Height + 10

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
	heightAux = heightAux + backupPath_btn.MinSize().Height + 10

	report_check := widget.NewCheck("Relatório", func(b bool) {
		fmt.Println(config.ReceiveReport)
		config.ReceiveReport = b
		fmt.Println(config.ReceiveReport)
	})
	report_check.Resize(fyne.NewSize(470, 30))
	report_check.Move(fyne.NewPos(10, heightAux))
	heightAux = heightAux + report_check.MinSize().Height + 10

	save_btn := widget.NewButton("Salvar", func() {
		config.Save()
	})
	save_btn.Resize(fyne.NewSize(100, 40))
	save_btn.Move(fyne.NewPos(380, 440))
	save_btn.Importance = widget.HighImportance
	save_btn.Refresh()

	wrapperContainer := container.NewWithoutLayout(
		label_name,
		inboundPath_label,
		inboundPath_display,
		inboundPath_btn,
		backupPath_label,
		backupPath_display,
		backupPath_btn,
		report_check,
		save_btn,
	)

	w.SetContent(wrapperContainer)

	w.Show()

	return w

}
