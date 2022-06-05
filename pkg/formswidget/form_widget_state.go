package formswidget

import "github.com/AllenDang/giu"

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
