package domgo

import "syscall/js"

type KeyEvent struct {
	e js.Value
}

func (k *KeyEvent) Key() string {
	return k.e.Get("key").String()
}
func (k *KeyEvent) Location() int {
	return k.e.Get("location").Int()
}

func (k *KeyEvent) Code() string {
	return k.e.Get("code").String()
}

func (k *KeyEvent) Shift() bool {
	return k.e.Get("shift").Bool()
}

func (k *KeyEvent) Ctrl() bool {
	return k.e.Get("ctrlKey").Bool()
}

func (k *KeyEvent) Meta() bool {
	return k.e.Get("metaKey").Bool()
}

func (k *KeyEvent) Alt() bool {
	return k.e.Get("altKey").Bool()
}
func (k *KeyEvent) Repeat() bool {
	return k.e.Get("repeat").Bool()
}
