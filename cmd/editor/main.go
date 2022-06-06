package main

import (
	_ "embed"

	"github.com/AllenDang/giu"

	"github.com/gucio321/forms-go/pkg/forms"
	"github.com/gucio321/forms-go/pkg/formseditorwidget"
	"github.com/gucio321/forms-go/pkg/formswidget"
)

const (
	windowW, windowH = 640, 480
)

//go:embed form.csv
var data []byte

var (
	form   *forms.Form
	layout int
)

func getMenubar() giu.Widget {
	return giu.Layout{
		giu.Menu("View").Layout(
			giu.Menu("View Mode").Layout(
				giu.RadioButton("Editor", layout == 0).OnChange(func() {
					layout = 0
				}),
				giu.RadioButton("Preview", layout == 1).OnChange(func() {
					layout = 1
				}),
				giu.RadioButton("Mixed", layout == 2).OnChange(func() {
					layout = 2
				}),
			),
		),
	}
}

func loop() {
	switch layout {
	case 0:
		giu.SingleWindowWithMenuBar().Layout(
			giu.MenuBar().Layout(getMenubar()),
			formseditorwidget.FormsEditor(form),
		)
	case 1:
		giu.SingleWindowWithMenuBar().Layout(
			giu.MenuBar().Layout(getMenubar()),
			formswidget.Form(form),
		)
	case 2:
		giu.MainMenuBar().Layout(getMenubar()).Build()
		giu.Window("Editor").Size(windowW/2, windowH).Layout(
			formseditorwidget.FormsEditor(form),
		)
		giu.Window("Preview").Size(windowW/2, windowH).Pos(windowW/2, 0).Layout(
			formswidget.Form(form),
		)
	}
}

func main() {
	form = forms.NewForm()
	form.Parse(data)
	wnd := giu.NewMasterWindow("Form editor", windowW, windowH, 0)
	wnd.Run(loop)
}
