package forms

import (
	"fmt"
	"strings"

	"github.com/gocarina/gocsv"
)

// QuestionType represents a type of question
type QuestionType byte

// types of question in a form
const (
	QuestionTypeSeparator QuestionType = iota
	QuestionTypeText
	QuestionTypeTextArea
	QuestionTypeCheckbox
	QuestionTypeRadio
	QuestionTypeSelect
)

type Options []string

func (o *Options) MarshalCSV() (data string, err error) {
	return strings.Join(*o, "/\\"), nil
}

func (o *Options) UnmarshalCSV(data string) (err error) {
	*o = strings.Split(data, "/\\")
	return nil
}

// Question represents a question in a form
type Question struct {
	Text        string       `csv:"text"`
	Type        QuestionType `csv:"type"`
	Placeholder string       `csv:"placeholder"`
	Options     Options
	Answer      string `csv:"answer"`
}

// Form represents a parsable form.
type Form struct {
	Questions []*Question
}

// NewForm returns a new form
func NewForm() *Form {
	return &Form{
		Questions: make([]*Question, 0),
	}
}

// Parse parses a form from a byte array
func (f *Form) Parse(data []byte) error {
	if err := gocsv.UnmarshalBytes(data, &f.Questions); err != nil {
		return fmt.Errorf("error parsing data: %w", err)
	}

	return nil
}

// Marshal marshals a form to a byte array
func (f *Form) Marshal() (data []byte, err error) {
	if data, err = gocsv.MarshalBytes(f.Questions); err != nil {
		return nil, fmt.Errorf("error marshalling data: %w", err)
	}

	return data, nil
}
