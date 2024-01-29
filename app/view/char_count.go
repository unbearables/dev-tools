package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateCharCountContent() *fyne.Container {
	charCountLabel := widget.NewLabel("Char count: 0")

	textField := widget.NewEntry()
	textField.SetPlaceHolder("Add text to count chars")
	textField.MultiLine = true
	textField.Wrapping = fyne.TextWrapBreak
	textField.OnChanged = func(s string) {
		count := 0
		if s != "" {
			count = len(s)
		}
		charCountLabel.SetText(fmt.Sprintf("Char count: %d", count))
	}

	return container.NewGridWithRows(2,
		textField,
		charCountLabel,
	)
}
