package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/google/uuid"
)

func CreateGeneratorContent(window fyne.Window) *fyne.Container {
	uuidGenBox := createUuidGeneratorBox(window)

	return container.NewVBox(uuidGenBox, layout.NewSpacer())
}

func createUuidGeneratorBox(window fyne.Window) *fyne.Container {
	uuidGenResult := widget.NewLabel(uuid.New().String())
	uuidGenBtn := widget.NewButton("Generate UUID", func() {
		uuidGenResult.SetText(uuid.New().String())
	})
	uuidCopyBtn := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		window.Clipboard().SetContent(uuidGenResult.Text)
	})

	return container.NewGridWithColumns(2,
		uuidGenBtn,
		container.NewBorder(nil, nil, uuidCopyBtn, nil, uuidGenResult),
	)
}
