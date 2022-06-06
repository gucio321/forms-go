package formseditorwidget

import (
	"fmt"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/gucio321/forms-go/pkg/forms"
)

var _ giu.Widget = &FormsEditorWidget{}

type FormsEditorWidget struct {
	id   string
	form *forms.Form
}

func FormsEditor(form *forms.Form) *FormsEditorWidget {
	return &FormsEditorWidget{
		id:   giu.GenAutoID("forms-editor-widget"),
		form: form,
	}
}

func (f *FormsEditorWidget) Build() {
	_, availableH := giu.GetAvailableRegion()
	state := f.getState()

	giu.Row(
		giu.Button("Add Question").OnClick(func() {
			newQuestion := &forms.Question{
				Text: "New Question",
				Type: forms.QuestionTypeText,
			}
			if state.selectedQuestion == -1 {
				f.form.Questions = append(f.form.Questions, newQuestion)
				return
			}

			f.form.Questions = append(
				f.form.Questions[:state.selectedQuestion+1],
				append(
					[]*forms.Question{
						newQuestion,
					},
					f.form.Questions[state.selectedQuestion+1:]...,
				)...,
			)

			state.selectedQuestion++
		}),
		giu.Button("Remove QUestion").OnClick(func() {
			f.form.Questions = append(f.form.Questions[:state.selectedQuestion], f.form.Questions[state.selectedQuestion+1:]...)
		}).Disabled(state.selectedQuestion < 0 || state.selectedQuestion >= len(f.form.Questions)),
		giu.Button("Move Up").OnClick(func() {
			f.form.Questions[state.selectedQuestion], f.form.Questions[state.selectedQuestion-1] = f.form.Questions[state.selectedQuestion-1], f.form.Questions[state.selectedQuestion]
			state.selectedQuestion--
		}).Disabled(state.selectedQuestion <= 0 || state.selectedQuestion > len(f.form.Questions)),
		giu.Button("Move Down").OnClick(func() {
			f.form.Questions[state.selectedQuestion], f.form.Questions[state.selectedQuestion+1] = f.form.Questions[state.selectedQuestion+1], f.form.Questions[state.selectedQuestion]
			state.selectedQuestion++
		}).Disabled(state.selectedQuestion < 0 || state.selectedQuestion >= len(f.form.Questions)-1),
	).Build()

	giu.SplitLayout(giu.DirectionVertical, availableH/2, giu.Custom(func() {
		pageID := 1
		open := imgui.TreeNode("Page 1")
		for i, question := range f.form.Questions {
			if question == nil {
				break
			}

			i := i
			if question.Type == forms.QuestionTypeSeparator {
				if open {
					imgui.TreePop()
				}
				pageID++
				open = imgui.TreeNode(fmt.Sprintf("Page %d##%v", pageID, f.id))
				if giu.IsItemClicked(giu.MouseButtonLeft) {
					state.selectedQuestion = i
				}

				continue
			}

			if open {
				giu.Selectable(question.Text).OnClick(func() {
					state.selectedQuestion = i
				}).Selected(state.selectedQuestion == i).Build()
			}
		}
		if open {
			imgui.TreePop()
		}
	}), giu.Custom(func() {
		if state.selectedQuestion < 0 || state.selectedQuestion >= len(f.form.Questions) {
			giu.Label("Please select a quesstion above").Build()
			return
		}

		question := f.form.Questions[state.selectedQuestion]
		questionTypes := make([]string, 0)
		for i := forms.QuestionTypeSeparator; i <= forms.QuestionTypeSelect; i++ {
			questionTypes = append(questionTypes, i.String())
		}

		qt := int32(question.Type)
		giu.Combo("Question Type", questionTypes[question.Type], questionTypes, &qt).OnChange(func() {
			question.Type = forms.QuestionType(qt)
		}).Build()

		if question.Type == forms.QuestionTypeSeparator {
			return
		}

		giu.InputText(&question.Text).Hint("What do you think about GO?").Label("Question Title##" + f.id).Build()
		switch question.Type {
		case forms.QuestionTypeText:
		case forms.QuestionTypeTextArea:
		case forms.QuestionTypeCheckbox,
			forms.QuestionTypeRadio,
			forms.QuestionTypeSelect:
			giu.TreeNode("Options").Layout(
				giu.Custom(func() {
					for i := range question.Options {
						giu.InputText(&question.Options[i]).Hint("Option").Labelf("Option %d##%v", i, f.id).Build()
					}
				}),
				giu.Row(
					giu.Button("Add Option").OnClick(func() {
						question.Options = append(question.Options, "")
					}),
					giu.Button("Remove Option").OnClick(func() {
						question.Options = question.Options[:len(question.Options)-1]
					}),
				),
			).Build()
		}
	})).Build()
}
