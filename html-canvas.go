package domgo

import (
	"errors"
	"math"
	"syscall/js"
)

type HtmlCanvas struct {
	Element
	Context2D *Context
}

func (c *HtmlCanvas) SetWidth(w int) {
	c.value.Set("width", w)
}

func (c *HtmlCanvas) SetHeight(h int) {
	c.value.Set("height", h)
}

func (c *HtmlCanvas) GetWidth() int {
	return c.value.Get("width").Int()
}

func (c *HtmlCanvas) GetHeight() int {
	return c.value.Get("height").Int()
}

func (c *HtmlCanvas) setValue(v js.Value) error {
	if getNodeName(v) != "canvas" {
		return errors.New("Cannot use that element as a canvas")
	}
	c.value = v
	c.Context2D = &Context{v.Call("getContext", "2d")}
	return nil
}

func (c *HtmlCanvas) getValue() js.Value {
	return c.value
}

type Context struct {
	value js.Value
}

func (c *Context) SetFillStyle(style string) {
	c.value.Set("fillStyle", style)
}

func (c *Context) Translate(x, y int) {
	c.value.Call("translate", x, y)
}

func (c *Context) Rotate(radians float64) {
	c.value.Call("rotate", radians)
}

func (c *Context) RotateDegrees(degrees float64) {
	c.Rotate(degrees * math.Pi / 180.0)
}

func (c *Context) BeginPath() {
	c.value.Call("beginPath")
}

func (c *Context) MoveTo(x, y int) {
	c.value.Call("moveTo", x, y)
}

func (c *Context) LineTo(x, y int) {
	c.value.Call("lineTo", x, y)
}

func (c *Context) Fill() {
	c.value.Call("fill")
}

func (c *Context) ClearRect(x1, y1, w, h int) {
	c.value.Call("clearRect", x1, y1, w, h)
}

func (c *Context) FillRect(x1, y1, w, h int) {
	c.value.Call("fillRect", x1, y1, w, h)
}
