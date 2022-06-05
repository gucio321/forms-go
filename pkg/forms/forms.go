package forms

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/gucio321/forms-go/pkg/formswidget"
)

// Form represents a parsable form.
type Form struct {
	Questions []*formswidget.Question
}

// NewForm returns a new form
func NewForm() *Form {
	return &Form{
		Questions: make([]*formswidget.Question, 0),
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
