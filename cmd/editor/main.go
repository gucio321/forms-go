package main

import (
	_ "embed"

	"github.com/AllenDang/giu"
	"github.com/gucio321/forms-go/pkg/forms"
	"github.com/gucio321/forms-go/pkg/formseditorwidget"
)

//go:embed form.csv
var data []byte

var form *forms.Form

func loop() {
	giu.SingleWindow().Layout(
		formseditorwidget.FormsEditor(form),
	)
}

func main() {
	form = forms.NewForm()
	form.Parse(data)
	wnd := giu.NewMasterWindow("Form editor", 640, 480, 0)
	wnd.Run(loop)
}
