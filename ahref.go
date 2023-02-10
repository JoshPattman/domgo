package domgo

import (
	"errors"
	"syscall/js"
)

type Link struct {
	Element
}

func (l *Link) setValue(v js.Value) error {
	if getNodeName(v) != "a" {
		return errors.New("Cannot use that element as a link")
	}
	l.value = v
	return nil
}

func (l *Link) getValue() js.Value {
	return l.value
}

func (l *Link) SetDownload(filename string) {
	l.value.Set("download", filename)
}

func (l *Link) GetDownload() string {
	return l.value.Get("download").String()
}

func (l *Link) SetHref(filename string) {
	l.value.Set("href", filename)
}

func (l *Link) GetHref() string {
	return l.value.Get("href").String()
}
