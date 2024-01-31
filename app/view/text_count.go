package view

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateCharCountContent() *fyne.Container {
	charCountLabel := widget.NewLabel("Char count: 0")
	wordCountLabel := widget.NewLabel("Word count: 0")

	textField := widget.NewEntry()
	textField.SetPlaceHolder("Add text to count chars")
	textField.MultiLine = true
	textField.Wrapping = fyne.TextWrapBreak
	textField.OnChanged = func(s string) {
		charCount := 0
		wordCount := 0
		if s != "" {
			charCount = len(s)
			wordCount = len(strings.Fields(s))
		}
		charCountLabel.SetText(fmt.Sprintf("Char count: %d", charCount))
		wordCountLabel.SetText(fmt.Sprintf("Word count: %d", wordCount))
	}

	return container.NewBorder(nil, container.NewVBox(
		charCountLabel,
		wordCountLabel,
	), nil, nil,
		textField,
	)
}
