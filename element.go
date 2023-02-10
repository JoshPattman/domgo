package domgo

import (
	"syscall/js"
)

type Element struct {
	value js.Value
}

func (e *Element) setValue(v js.Value) error {
	e.value = v
	return nil
}

func (e *Element) getValue() js.Value {
	return e.value
}

// TODO : implement enough stuff to remove these two
func (d *Element) GetVariable(varName string) js.Value {
	return d.value.Get(varName)
}

func (d *Element) SetVariable(varName string, value any) {
	d.value.Set(varName, value)
}

func (e *Element) SetInnerHtml(html string) {
	e.value.Set("innerHTML", html)
}
func (e *Element) GetInnerHtml() string {
	return e.value.Get("innerHTML").String()
}

func (e *Element) SetStyle(style, value string) {
	e.value.Get("style").Set(style, value)
}

func (e *Element) GetStyle(style string) string {
	s := e.value.Get("style").Get(style).String()
	return s
}

func (e *Element) SetHidden(h bool) {
	e.value.Set("hidden", h)
}
func (e *Element) GetHidden() bool {
	return e.value.Get("hidden").Bool()
}

// MOUSE LISTENERS
func makeJSMouseFunc(f func(MouseEvent)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		me := MouseEvent{args[0]}
		f(me)
		return nil
	})
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Element) AddMouseDownListener(f func(MouseEvent)) func() {
	jsf := makeJSMouseFunc(f)
	d.value.Call("addEventListener", "mousedown", jsf)
	return func() { jsf.Release() }
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Element) AddMouseUpListener(f func(MouseEvent)) func() {
	jsf := makeJSMouseFunc(f)
	d.value.Call("addEventListener", "mouseup", jsf)
	return func() { jsf.Release() }
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Element) AddMouseClickListener(f func(MouseEvent)) func() {
	jsf := makeJSMouseFunc(f)
	d.value.Call("addEventListener", "click", jsf)
	return func() { jsf.Release() }
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Element) AddMouseDoubleClickListener(f func(MouseEvent)) func() {
	jsf := makeJSMouseFunc(f)
	d.value.Call("addEventListener", "doubleclick", jsf)
	return func() { jsf.Release() }
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Element) AddMouseOverListener(f func(MouseEvent)) func() {
	jsf := makeJSMouseFunc(f)
	d.value.Call("addEventListener", "mouseover", jsf)
	return func() { jsf.Release() }
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Element) AddMouseOutListener(f func(MouseEvent)) func() {
	jsf := makeJSMouseFunc(f)
	d.value.Call("addEventListener", "mouseout", jsf)
	return func() { jsf.Release() }
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Element) AddMouseEnterListener(f func(MouseEvent)) func() {
	jsf := makeJSMouseFunc(f)
	d.value.Call("addEventListener", "mouseenter", jsf)
	return func() { jsf.Release() }
}

// AddMouseMoveListener adds a listener to mouse move and returns a function to stop listening
func (d *Element) AddMouseLeaveListener(f func(MouseEvent)) func() {
	jsf := makeJSMouseFunc(f)
	d.value.Call("addEventListener", "mouseleave", jsf)
	return func() { jsf.Release() }
}
