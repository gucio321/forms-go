package formswidget

import (
	"fmt"
	"github.com/AllenDang/giu"
	"github.com/gucio321/forms-go/pkg/forms"
	"strings"
)

var _ giu.Disposable = &formWidgetState{}

type formWidgetState struct {
	currentPage int
}

func (f *formWidgetState) Dispose() {
	// noop
}

func (f *FormsWidget) newState() *formWidgetState {
	return &formWidgetState{
		currentPage: 0,
	}
}

func (f *FormsWidget) getState() (state *formWidgetState) {
	if s := giu.Context.GetState(f.id); s == nil {
		state = f.newState()
		giu.Context.SetState(f.id, state)
	} else {
		state = s.(*formWidgetState)
	}

	return state
}

// static check if FormsWidget implemented giu.Widget interface
var _ giu.Widget = &FormsWidget{}

type FormsWidget struct {
	id        string
	formPages [][]*forms.Question
}

func Form(form *forms.Form) *FormsWidget {
	result := &FormsWidget{
		id:        giu.GenAutoID("FormsWidget"),
		formPages: make([][]*forms.Question, 0),
	}

	page := make([]*forms.Question, 0)
	for _, question := range form.Questions {
		if question == nil {
			break
		}

		if question.Type == forms.QuestionTypeSeparator {
			result.formPages = append(result.formPages, page)
			page = make([]*forms.Question, 0)
		} else {
			page = append(page, question)
		}
	}

	result.formPages = append(result.formPages, page)

	return result
}

func (f *FormsWidget) Build() {
	state := f.getState()
	if (state.currentPage < 0) || (state.currentPage >= len(f.formPages)) {
		panic("invalid page index!")
	}

	currentPage := f.formPages[state.currentPage]
	giu.ProgressBar(float32(state.currentPage) / float32(len(f.formPages))).Build()
	for _, question := range currentPage {
		switch question.Type {
		case forms.QuestionTypeSeparator:
			panic("fatal: smething went wrong here")
		case forms.QuestionTypeText:
			giu.InputText(&question.Answer).Build()
		case forms.QuestionTypeTextArea:
			giu.InputTextMultiline(&question.Answer).Build()
		case forms.QuestionTypeCheckbox:
			answersStr := strings.ReplaceAll(question.Answer, " ", "")
			answers := strings.Split(answersStr, "/\\")
			for ; len(answers) < len(question.Options); answers = append(answers, "") {
			}

			for i, option := range question.Options {
				answer := i < len(answers) && answers[i] == "true"
				giu.Checkbox(option, &answer).OnChange(func() {
					answers[i] = fmt.Sprintf("%v", answer)
					answersStr = strings.Join(answers, "/\\")
					question.Answer = answersStr
				}).Build()
			}
		case forms.QuestionTypeRadio:
			//giu.RadioButton(question.Answer).Build()
		case forms.QuestionTypeSelect:
			//giu.ComboBox(question.Answer, question.Options).Build()
		}
	}
	giu.Row(
		giu.Button("Previous").OnClick(func() {
			state.currentPage--
		}).Disabled(state.currentPage == 0),
		giu.Condition(state.currentPage == len(f.formPages)-1,
			giu.Layout{
				giu.Button("Submit").OnClick(func() {
					// noop
				}),
			},
			giu.Layout{
				giu.Button("Next").OnClick(func() {
					state.currentPage++
				}),
			},
		),
	).Build()
}
