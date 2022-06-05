package forms

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
