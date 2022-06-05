package forms

// Question represents a question in a form
type Question struct {
	Text        string       `csv:"text"`
	Type        QuestionType `csv:"type"`
	Placeholder string       `csv:"placeholder"`
	Options     Options
	Answer      string `csv:"answer"`
}
