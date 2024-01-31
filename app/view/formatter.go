package view

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/enescakir/emoji"
)

func CreateJsonFormatterContent() *fyne.Container {
	textArea := widget.NewMultiLineEntry()
	textArea.SetPlaceHolder(fmt.Sprintf("%v Insert JSON to format", emoji.Snail))
	errMsg := widget.NewLabel("")
	formatBtn := widget.NewButton("Format", func() {
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, []byte(textArea.Text), "", "  ")
		if err == nil {
			textArea.SetText(prettyJSON.String())
			errMsg.SetText(fmt.Sprintf("All good %v", emoji.ThumbsUp))
		} else {
			errMsg.SetText(fmt.Sprintf("%v Invalid JSON: %s", emoji.Warning, err.Error()))
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
