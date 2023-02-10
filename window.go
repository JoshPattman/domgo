package domgo

import "syscall/js"

type Window struct {
	value js.Value
}

func GetWindow() *Window {
	return &Window{
		value: js.Global().Get("window"),
	}
}

func (w *Window) GetInnerWidth() int {
	return w.value.Get("innerWidth").Int()
}

func (w *Window) GetInnerHeight() int {
	return w.value.Get("innerHeight").Int()
}

func Alert(txt string) {
	js.Global().Call("alert", txt)
}
