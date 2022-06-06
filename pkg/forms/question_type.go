package forms

import "errors"

//go:generate stringer -type=QuestionType -trimprefix=QuestionType

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

func (o *QuestionType) MarshalCSV() (data string, err error) {
	return string(rune(*o)), nil
}

func (o *QuestionType) UnmarshalCSV(data string) (err error) {
	if len(data) != 1 {
		return errors.New("unexpected len of data")
	}

	*o = QuestionType(data[0])

	return nil
}
