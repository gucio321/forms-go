package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/AllenDang/giu"
	"github.com/gucio321/forms-go/pkg/forms"
	"github.com/gucio321/forms-go/pkg/formswidget"
)

var (
	title    = "PlaceHolder"
	formText = ""
)

var form = forms.NewForm()

func loop() {
	giu.SingleWindow().Layout(
		formswidget.Form(form),
	)
}

func main() {
	wnd := giu.NewMasterWindow(title, 640, 480, 0)
	formBytes := strings.Split(formText, "x")
	formData := make([]byte, 0)
	for _, x := range formBytes {
		result, err := strconv.ParseInt(x, 10, 8)
		if err != nil {
			log.Panicf("error parsing data: %v, string is %v", err, formText)
		}
		formData = append(formData, byte(result))
	}
	if err := form.Parse([]byte(formData)); err != nil {
		log.Panicf("error loading form")
	}
	wnd.Run(loop)
}
