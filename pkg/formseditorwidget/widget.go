package formseditorwidget

import (
	"github.com/AllenDang/giu"
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
			if state.selectedQuestion == -1 {
				f.form.Questions = append(f.form.Questions, &forms.Question{})
				return
			}

			f.form.Questions = append(f.form.Questions[:state.selectedQuestion],
				append([]*forms.Question{&forms.Question{}}, f.form.Questions[state.selectedQuestion:]...)...)
		}),
		giu.Button("Remove QUestion").OnClick(func() {
			f.form.Questions = append(f.form.Questions[:state.selectedQuestion], f.form.Questions[state.selectedQuestion+1:]...)
		}).Disabled(state.selectedQuestion < 0 || state.selectedQuestion >= len(f.form.Questions)),
		giu.Button("Move Up").OnClick(func() {
			f.form.Questions[state.selectedQuestion], f.form.Questions[state.selectedQuestion-1] =
				f.form.Questions[state.selectedQuestion-1], f.form.Questions[state.selectedQuestion]
			state.selectedQuestion--
		}).Disabled(state.selectedQuestion <= 0 || state.selectedQuestion > len(f.form.Questions)),
		giu.Button("Move Down").OnClick(func() {
			f.form.Questions[state.selectedQuestion], f.form.Questions[state.selectedQuestion+1] =
				f.form.Questions[state.selectedQuestion+1], f.form.Questions[state.selectedQuestion]
			state.selectedQuestion++
		}).Disabled(state.selectedQuestion < 0 || state.selectedQuestion >= len(f.form.Questions)-1),
	).Build()

	giu.SplitLayout(giu.DirectionVertical, availableH/2, giu.Custom(func() {
		for i, question := range f.form.Questions {
			i := i
			giu.Selectable(question.Text).OnClick(func() {
				state.selectedQuestion = i
			}).Selected(state.selectedQuestion == i).Build()
		}
	}), giu.Custom(func() {

	})).Build()
}
