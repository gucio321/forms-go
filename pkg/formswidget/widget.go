package formswidget

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AllenDang/giu"
	"github.com/gucio321/forms-go/pkg/forms"
)

const (
	buttonH = 40
	buttonW = 150
)

type OnSubmitCallback func(form *forms.Form)

// static check if FormsWidget implemented giu.Widget interface
var _ giu.Widget = &FormsWidget{}

type FormsWidget struct {
	id        string
	form      *forms.Form
	formPages [][]*forms.Question
	onSubmit  OnSubmitCallback
}

func Form(form *forms.Form) *FormsWidget {
	result := &FormsWidget{
		id:        giu.GenAutoID("FormsWidget"),
		formPages: make([][]*forms.Question, 0),
		form:      form,
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

func (f *FormsWidget) OnSubmit(cb OnSubmitCallback) *FormsWidget {
	f.onSubmit = cb
	return f
}

func (f *FormsWidget) Build() {
	state := f.getState()
	if (state.currentPage < 0) || (state.currentPage >= len(f.formPages)) {
		panic("invalid page index!")
	}

	currentPage := f.formPages[state.currentPage]

	giu.ProgressBar(float32(state.currentPage)/float32(len(f.formPages))).
		Overlayf("%d/%d", state.currentPage+1, len(f.formPages)).
		Build()

	rows := make([]*giu.TableRowWidget, 0)
	for _, question := range currentPage {
		rows = append(rows, giu.TableRow(giu.Custom(func() {
			giu.Markdown(&question.Text).Build()
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
				answer, err := strconv.Atoi(question.Answer)
				if err != nil {
					answer = -1
				}

				for i, option := range question.Options {
					giu.RadioButton(option, i == answer).OnChange(func() {
						question.Answer = strconv.Itoa(i)
					}).Build()
				}
			case forms.QuestionTypeSelect:
				answerInt, err := strconv.ParseInt(question.Answer, 10, 32)
				if err != nil {
					answerInt = 0
				}

				answer := int32(answerInt)

				giu.Combo("", question.Options[answer], question.Options, &answer).OnChange(func() {
					question.Answer = strconv.Itoa(int(answer))
				}).Build()
			}
		})))
	}

	_, availableH := giu.GetAvailableRegion()
	_, spacingH := giu.GetItemSpacing()
	tableH := availableH - buttonH - spacingH
	giu.Table().Rows(rows...).Size(-1, tableH).Build()

	giu.Row(
		giu.Button("Previous").OnClick(func() {
			state.currentPage--
		}).Disabled(state.currentPage == 0).Size(buttonW, buttonH),
		giu.Condition(state.currentPage == len(f.formPages)-1,
			giu.Layout{
				giu.Button("Submit").OnClick(func() {
					if f.onSubmit != nil {
						f.onSubmit(f.form)
					}
				}).Size(buttonW, buttonH),
			},
			giu.Layout{
				giu.Button("Next").OnClick(func() {
					state.currentPage++
				}).Size(buttonW, buttonH),
			},
		),
	).Build()
}
