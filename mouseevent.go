package domgo

import "syscall/js"

type MouseEvent struct {
	e js.Value
}

func (m *MouseEvent) Alt() bool {
	return m.e.Get("altKey").Bool()
}

func (m *MouseEvent) Button() int {
	return m.e.Get("button").Int()
}

/*func (m *MouseEvent) GetAltKey() bool {
	return m.e.Get("buttons").
}*/

func (m *MouseEvent) ClientX() float64 {
	return m.e.Get("clientX").Float()
}

func (m *MouseEvent) ClientY() float64 {
	return m.e.Get("clientY").Float()
}

func (m *MouseEvent) Ctrl() bool {
	return m.e.Get("ctrlKey").Bool()
}

func (m *MouseEvent) Meta() bool {
	return m.e.Get("metaKey").Bool()
}

func (m *MouseEvent) MovementX() float64 {
	return m.e.Get("movementX").Float()
}

func (m *MouseEvent) MovementY() float64 {
	return m.e.Get("movementY").Float()
}

func (m *MouseEvent) OffsetX() float64 {
	return m.e.Get("offsetX").Float()
}

func (m *MouseEvent) OffsetY() float64 {
	return m.e.Get("offsetY").Float()
}

func (m *MouseEvent) PageX() float64 {
	return m.e.Get("pageX").Float()
}

func (m *MouseEvent) PageY() float64 {
	return m.e.Get("pageY").Float()
}

/*func (m *MouseEvent) Region() bool {
	return m.e.Get("altKey").Bool()
}

func (m *MouseEvent) GetAltKey() bool {
	return m.e.Get("altKey").Bool()
}*/

func (m *MouseEvent) ScreenX() float64 {
	return m.e.Get("screenX").Float()
}

func (m *MouseEvent) ScreenY() float64 {
	return m.e.Get("screenY").Float()
}

func (m *MouseEvent) Shift() bool {
	return m.e.Get("shift").Bool()
}

/*
func (m *MouseEvent) MozPressure() float64 {
	return m.e.Get("mozPressure").Float()
}

func (m *MouseEvent) GetAltKey() bool {
	return m.e.Get("altKey").Bool()
}

func (m *MouseEvent) GetAltKey() bool {
	return m.e.Get("altKey").Bool()
}*/
