package view

import (
	"bytes"
	"encoding/json"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateJsonFormatterContent() *fyne.Container {
	textArea := widget.NewMultiLineEntry()
	textArea.SetPlaceHolder("Insert JSON to format")
	errMsg := widget.NewLabel("")
	formatBtn := widget.NewButton("Format", func() {
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, []byte(textArea.Text), "", "  ")
		if err == nil {
			textArea.SetText(prettyJSON.String())
			errMsg.SetText("All good 👍")
		} else {
			errMsg.SetText("⚠️ Invalid JSON: " + err.Error())
		}
	})
	formatBtn.Disable()

	textArea.OnChanged = func(newText string) {
		if strings.TrimSpace(newText) == "" {
			formatBtn.Disable()
		} else {
			formatBtn.Enable()
		}
	}

	return container.NewBorder(nil, container.NewGridWithRows(2, errMsg, formatBtn), nil, nil, textArea)
}
