package domgo

import (
	"errors"
	"strconv"
	"syscall/js"
)

type Image struct {
	Element
}

func (i *Image) setValue(v js.Value) error {
	if getNodeName(v) != "img" {
		return errors.New("Cannot use that element as an image")
	}
	i.value = v
	return nil
}

func (i *Image) getValue() js.Value {
	return i.value
}

func (i *Image) SetSource(source string) {
	i.value.Set("src", source)
}
func (i *Image) GetSource() string {
	return i.value.Get("src").String()
}

func (i *Image) SetAlt(alt string) {
	i.value.Set("alt", alt)
}
func (i *Image) GetAlt() string {
	return i.value.Get("alt").String()
}

func (i *Image) SetHeight(h int) {
	i.value.Set("height", strconv.Itoa(h))
}
func (i *Image) GetHeight() int {
	v, _ := strconv.Atoi(i.value.Get("height").String())
	return v
}

func (i *Image) SetWidth(h int) {
	i.value.Set("width", strconv.Itoa(h))
}
func (i *Image) GetWidth() int {
	v, _ := strconv.Atoi(i.value.Get("width").String())
	return v
}
