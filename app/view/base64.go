package view

import (
	"encoding/base64"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateBase64Content() *fyne.Container {
	textField := widget.NewEntry()
	textField.SetPlaceHolder("Add text to encode")
	textField.MultiLine = true
	textField.Wrapping = fyne.TextWrapBreak

	base64Field := widget.NewEntry()
	base64Field.SetPlaceHolder("Add base64 to decode")
	base64Field.MultiLine = true
	base64Field.Wrapping = fyne.TextWrapBreak

	encodeBtn := widget.NewButtonWithIcon("Encode", theme.MoveDownIcon(), func() {
		if textField.Text != "" {
			out := base64.StdEncoding.EncodeToString([]byte(textField.Text))
			base64Field.SetText(out)
		}
	})
	encodeBtn.Disable()
	textField.OnChanged = func(s string) {
		if s == "" {
			encodeBtn.Disable()
		} else {
			encodeBtn.Enable()
		}
	}

	decodeBtn := widget.NewButtonWithIcon("Decode", theme.MoveUpIcon(), func() {
		if base64Field.Text != "" {
			out, err := base64.StdEncoding.DecodeString(base64Field.Text)
			if err == nil {
				textField.SetText(string(out))
			} else {
				textField.SetText(err.Error())
			}
		}
	})
	decodeBtn.Disable()
	base64Field.OnChanged = func(s string) {
		if s == "" {
			decodeBtn.Disable()
		} else {
			decodeBtn.Enable()
		}
	}

	return container.NewGridWithRows(2,
		container.NewBorder(
			nil,
			container.NewGridWithColumns(4, layout.NewSpacer(), encodeBtn, decodeBtn, layout.NewSpacer()), nil, nil, textField),
		base64Field,
	)
}
