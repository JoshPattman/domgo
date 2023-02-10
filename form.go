package domgo

import (
	"errors"
	"syscall/js"
)

type Form struct {
	Element
	ctx     *Context
	imgData js.Value
	buf     js.Value
}

func (f *Form) setValue(v js.Value) error {
	if getNodeName(v) != "form" {
		return errors.New("Cannot use that element as a form")
	}
	f.value = v
	return nil
}

func (f *Form) getValue() js.Value {
	return f.value
}

func (f *Form) Submit() {
	f.value.Call("submit")
}

type Input struct {
	value js.Value
}

func (i *Input) setValue(v js.Value) error {
	if getNodeName(v) != "input" {
		return errors.New("Cannot use that element as an input")
	}
	i.value = v
	return nil
}

func (i *Input) getValue() js.Value {
	return i.value
}

func (i *Input) GetFormValue() string {
	return i.value.Get("value").String()
}
