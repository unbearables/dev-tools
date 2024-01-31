package main

import (
	"fmt"
	"image/color"
	"net/url"

	"github.com/enescakir/emoji"
	"github.com/unbearables/dev-tools/app/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow(fmt.Sprintf("%v Dev tools", emoji.HammerAndWrench))
	size := fyne.Size{Width: 900, Height: 600}
	w.Resize(size)
	w.SetContent(makeUI(w))
	w.ShowAndRun()
}

const VERSION = "0.1.0"

func makeUI(w fyne.Window) fyne.CanvasObject {
	workspace := container.New(layout.NewPaddedLayout())

	welcomeMsg := canvas.NewText(fmt.Sprintf("Welcome, fellow developer %v", emoji.ConstructionWorker), theme.PrimaryColor())
	welcomeMsg.TextSize = 30
	welcomeMsg.Alignment = fyne.TextAlignCenter

	versionMsg := canvas.NewText("You are runnning Dev Tool "+VERSION, color.White)

	projectMsg := canvas.NewText("If you want to report a bug, suggest a feature or contribute, you can find this project at", color.White)
	u, _ := url.Parse("https://github.com/unbearables/dev-tools")
	gtihubLink := widget.NewHyperlinkWithStyle("github.com/unbearables/dev-tools", u, fyne.TextAlignCenter, fyne.TextStyle{})

	workspace.Add(container.NewCenter(container.NewVBox(welcomeMsg, versionMsg, projectMsg, gtihubLink)))

	menu := view.CreateMenu(w, workspace)

	return container.NewBorder(nil, nil, menu, nil, workspace)
}
