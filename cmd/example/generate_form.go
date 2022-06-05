package main

import (
	"os"

	"github.com/gucio321/forms-go/pkg/forms"
)

func main() {
	form := forms.NewForm()
	form.Questions = append(form.Questions,
		&forms.Question{
			Text: "Hi there!",
			Type: forms.QuestionTypeText,
		},
		&forms.Question{
			Type: forms.QuestionTypeSeparator,
		},
		&forms.Question{
			Text:    "Check us!",
			Type:    forms.QuestionTypeCheckbox,
			Options: []string{"Me!", "And me too!"},
		},
	)

	data, err := form.Marshal()
	if err != nil {
		panic(err)
	}
	os.WriteFile("./form.csv", data, 0o644)
}
