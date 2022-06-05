package formseditorwidget

import "github.com/AllenDang/giu"

type formEditorWidgetState struct {
	selectedQuestion int
}

func (f *formEditorWidgetState) Dispose() {
	// noop
}

func (f *FormsEditorWidget) newState() *formEditorWidgetState {
	return &formEditorWidgetState{
		selectedQuestion: -1,
	}
}

func (f *FormsEditorWidget) getState() *formEditorWidgetState {
	if s := giu.Context.GetState(f.id); s != nil {
		return s.(*formEditorWidgetState)
	}

	state := f.newState()
	giu.Context.SetState(f.id, state)
	return state
}
