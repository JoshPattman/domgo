package domgo

import (
	"errors"
	"syscall/js"
)

type documentElement interface {
	setValue(js.Value) error
	getValue() js.Value
}

func GetDocument() *Document {
	return &Document{
		value: js.Global().Get("document"),
	}
}

type Document struct {
	value js.Value
}

func (d *Document) GetElementById(id string, v documentElement) error {
	elem := d.value.Call("getElementById", id)
	if elem.IsNull() {
		return errors.New("Could not find element with id '" + id + "'")
	}
	return v.setValue(elem)
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Document) AddMouseMoveListener(f func(me MouseEvent)) func() {
	jsf := makeJSMouseFunc(f)
	d.value.Call("addEventListener", "mousemove", jsf)
	return func() { jsf.Release() }
}

func (d *Document) GetBody() *Element {
	return &Element{d.value.Get("body")}
}

func makeJSKeyboardFunc(f func(KeyEvent)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ke := KeyEvent{args[0]}
		f(ke)
		return nil
	})
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Document) AddKeyDownListener(f func(KeyEvent)) func() {
	jsf := makeJSKeyboardFunc(f)
	d.value.Call("addEventListener", "keydown", jsf)
	return func() { jsf.Release() }
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Document) AddKeyUpListener(f func(KeyEvent)) func() {
	jsf := makeJSKeyboardFunc(f)
	d.value.Call("addEventListener", "keyup", jsf)
	return func() { jsf.Release() }
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Document) AddKeyPressListener(f func(KeyEvent)) func() {
	jsf := makeJSKeyboardFunc(f)
	d.value.Call("addEventListener", "keypress", jsf)
	return func() { jsf.Release() }
}
