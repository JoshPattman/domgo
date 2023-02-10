package domgo

import (
	"errors"
	"image"
	"syscall/js"
	"time"
)

type BufferedCanvas struct {
	Element
	ctx     *Context
	imgData js.Value
	buf     js.Value
}

func (c *BufferedCanvas) setValue(v js.Value) error {
	if getNodeName(v) != "canvas" {
		return errors.New("Cannot use that element as a canvas")
	}
	c.value = v
	c.ctx = &Context{v.Call("getContext", "2d")}
	c.setupImageBuffer()
	return nil
}

func (c *BufferedCanvas) setupImageBuffer() {
	width, height := c.GetWidth(), c.GetHeight()
	c.imgData = c.ctx.value.Call("createImageData", width, height)
	c.buf = js.Global().Get("Uint8Array").New(width * height * 4)
}

func (c *BufferedCanvas) Draw(img *image.RGBA) error {
	width, height := c.GetWidth(), c.GetHeight()
	if len(img.Pix) != width*height*4 {
		return errors.New("Image size does not match canvas size")
	}
	js.CopyBytesToJS(c.buf, img.Pix)
	c.imgData.Get("data").Call("set", c.buf)
	c.ctx.value.Call("putImageData", c.imgData, 0, 0)
	time.Sleep(time.Second / 10000)
	return nil
}

func (c *BufferedCanvas) getValue() js.Value {
	return c.value
}

func (c *BufferedCanvas) SetWidth(w int) {
	c.value.Set("width", w)
	c.setupImageBuffer()
}

func (c *BufferedCanvas) SetHeight(h int) {
	c.value.Set("height", h)
	c.setupImageBuffer()
}

func (c *BufferedCanvas) GetWidth() int {
	return c.value.Get("width").Int()
}

func (c *BufferedCanvas) GetHeight() int {
	return c.value.Get("height").Int()
}
