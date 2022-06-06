package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/AllenDang/giu"
	"github.com/sqweek/dialog"

	"github.com/gucio321/forms-go/pkg/forms"
	"github.com/gucio321/forms-go/pkg/formseditorwidget"
	"github.com/gucio321/forms-go/pkg/formswidget"
)

const (
	windowW, windowH = 640, 480
)

var (
	form   *forms.Form
	layout int
)

func getMenubar() giu.Widget {
	return giu.Layout{
		giu.PrepareMsgbox(),
		giu.Menu("File").Layout(
			giu.MenuItem("New").OnClick(func() {
				form = forms.NewForm()
			}),
			giu.MenuItem("Open").OnClick(func() {
				filename, err := dialog.File().Filter("CSV file", "csv").Load()
				if err != nil {
					log.Printf("Error: %v", err)
					giu.Msgbox("Error!", fmt.Sprintf("Error opening load file dialogue: %v", err))
				}

				data, err := os.ReadFile(filename)
				if err != nil {
					log.Printf("Error: %v", err)
					giu.Msgbox("Error!", fmt.Sprintf("Error reading from file: %v", err))
				}

				err = form.Parse(data)
				if err != nil {
					log.Printf("Error: %v", err)
					giu.Msgbox("Error!", fmt.Sprintf("Error loading file: %v", err))
				}
			}),
			giu.MenuItem("Save").OnClick(func() {
				data, err := form.Marshal()
				if err != nil {
					log.Printf("Error: %v", err)
					giu.Msgbox("Error!", fmt.Sprintf("Error converting form's data: %v", err))
				}

				filename, err := dialog.File().Filter("CSV file", "csv").Save()
				if err != nil {
					log.Printf("Error: %v", err)
					giu.Msgbox("Error!", fmt.Sprintf("Error opening save file dialogue: %v", err))
				}

				err = os.WriteFile(filename, data, 0o644)
				if err != nil {
					log.Printf("Error: %v", err)
					giu.Msgbox("Error!", fmt.Sprintf("Error saving file: %v", err))
				}
			}),
			giu.MenuItem("Export").OnClick(func() {
				formBytes, err := form.Marshal()
				if err != nil {
					log.Printf("Error: %v", err)
					giu.Msgbox("Error!", fmt.Sprintf("Error marshaling form data: %v", err))
				}

				formStr := ""
				for i := 0; i < len(formBytes); i++ {
					formStr += strconv.Itoa(int(formBytes[i])) + "x"
				}

				if _, err := os.Stat("./compiler/main.go"); err != nil {
					log.Printf("Error: no required project found, you have to have forms-go project downloaded")
					giu.Msgbox("Error!", fmt.Sprintf("Error You must have forms-go project downloaded and run this app from cmd/editor, elsewhere export feature will not work"))
				}

				formStr = formStr[:len(formStr)-2]

				cmd := exec.Command(
					"bash",
					"-c",
					"go "+
						"build "+
						"-ldflags="+
						"\"-X main.formText="+formStr+"\" "+
						"-o "+
						"output "+
						"./compiler/main.go",
				)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					log.Printf("Error: %v", err)
					giu.Msgbox("Error!", fmt.Sprintf("Error exporting binary: %v", err))
				}
			}),
		),
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
	wnd := giu.NewMasterWindow("Form editor", windowW, windowH, 0)
	wnd.Run(loop)
}
