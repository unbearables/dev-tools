package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateMenu(window fyne.Window, workspace *fyne.Container) *widget.Tree {
	list := map[string][]string{
		"":                {"Generator", "Char count", "Formatter", "Decode / Encode"},
		"Formatter":       {"JSON"},
		"Decode / Encode": {"Base64", "JWT"},
	}
	tree := widget.NewTreeWithStrings(list)
	tree.OpenAllBranches()

	lastSelection := ""
	tree.OnSelected = func(uid string) {
		if lastSelection == uid {
			return // selection has not changed
		}

		var content *fyne.Container = nil
		switch uid {
		case "Generator":
			content = CreateGeneratorContent(window)
		case "Char count":
			content = CreateCharCountContent()
		case "JSON":
			content = CreateJsonFormatterContent()
		case "Base64":
			content = CreateBase64Content()
		case "JWT":
			content = CreateJwtContent()
		}
		if content != nil {
			workspace.Objects = []fyne.CanvasObject{}
			workspace.Add(content)
			lastSelection = uid
		}
	}

	return tree
}
