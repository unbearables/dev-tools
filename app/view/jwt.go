package view

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/enescakir/emoji"
)

func CreateJwtContent() *fyne.Container {
	feedbackMsg := widget.NewLabel(fmt.Sprintf("%v Waiting for input", emoji.Snail))

	claimsField := widget.NewEntry()
	claimsField.SetPlaceHolder("{}")
	claimsField.MultiLine = true
	claimsField.Wrapping = fyne.TextWrapBreak

	jwtField := widget.NewEntry()
	jwtField.SetPlaceHolder("Add JWT here")
	jwtField.MultiLine = true
	jwtField.Wrapping = fyne.TextWrapBreak
	jwtField.OnChanged = func(s string) {
		if s == "" {
			return
		}
		jwtParts := strings.Split(s, ".")
		if len(jwtParts) != 3 {
			feedbackMsg.SetText(fmt.Sprintf("%v Invalid JWT: incorrect number of JWT parts", emoji.Warning))
			return
		}

		decodedStr, err := base64.StdEncoding.DecodeString(jwtParts[1] + "==")
		if err == nil {
			claimPart := string(decodedStr)
			var prettyJson bytes.Buffer
			err := json.Indent(&prettyJson, []byte(claimPart), "", "  ")
			if err != nil {
				feedbackMsg.SetText(fmt.Sprintf("%v Invalid JWT: claims JSON is not valid", emoji.Warning))
				return
			}

			claimsField.SetText(string(prettyJson.Bytes()))
			feedbackMsg.SetText(fmt.Sprintf("%v Valid JWT", emoji.CheckMark))
		} else {
			feedbackMsg.SetText(fmt.Sprintf("%v Invalid JWT: claims part not a valid base64", emoji.Warning))
		}
	}

	return container.NewBorder(feedbackMsg, nil, nil, nil, container.NewGridWithColumns(2,
		jwtField,
		claimsField,
	))
}
