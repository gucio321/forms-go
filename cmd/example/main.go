package main

import (
	_ "embed"

	"github.com/AllenDang/giu"
	"github.com/gucio321/forms-go/pkg/forms"
	"github.com/gucio321/forms-go/pkg/formswidget"
)

//go:embed form.csv
var data []byte

var form *forms.Form

func loop() {
	giu.SingleWindow().Layout(
		formswidget.Form(form),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Example of Forms-go", 640, 480, 0)
	form = forms.NewForm()
	form.Parse(data)
	wnd.Run(loop)
}
