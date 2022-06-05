package main

import (
	"github.com/gucio321/forms-go/pkg/formswidget"
	"os"

	"github.com/gucio321/forms-go/pkg/forms"
)

func main() {
	form := forms.NewForm()
	form.Questions = append(form.Questions,
		&formswidget.Question{
			Text: "Hi there!",
			Type: formswidget.QuestionTypeText,
		},
		&formswidget.Question{
			Type: formswidget.QuestionTypeSeparator,
		},
		&formswidget.Question{
			Text:    "Check us!",
			Type:    formswidget.QuestionTypeCheckbox,
			Options: []string{"Me!", "And me too!"},
		},
	)

	data, err := form.Marshal()
	if err != nil {
		panic(err)
	}
	os.WriteFile("./form.csv", data, 0o644)
}
