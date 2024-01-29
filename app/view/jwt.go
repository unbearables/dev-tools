package view

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateJwtContent() *fyne.Container {
	feedbackMsg := widget.NewLabel("🐌 Waiting for input")

	claimsField := widget.NewEntry()
	claimsField.SetPlaceHolder("{}")
	// claimsField.Disable()
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
			feedbackMsg.SetText("⚠️ Invalid JWT: incorrect number of part")
			return
		}

		decodedStr, err := base64.StdEncoding.DecodeString(jwtParts[1] + "==")
		if err == nil {
			claimPart := string(decodedStr)
			var prettyJson bytes.Buffer
			err := json.Indent(&prettyJson, []byte(claimPart), "", "  ")
			if err != nil {
				feedbackMsg.SetText("⚠️ Invalid JWT: claims JSON is not valid")
				return
			}

			claimsField.SetText(string(prettyJson.Bytes()))
			feedbackMsg.SetText("✅ Valid JWT")
		} else {
			feedbackMsg.SetText("⚠️ Invalid JWT: claims part not a valid base64")
		}
	}

	return container.NewBorder(feedbackMsg, nil, nil, nil, container.NewGridWithColumns(2,
		jwtField,
		claimsField,
	))
}
