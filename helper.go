package domgo

import (
	"strings"
	"syscall/js"
)

func getNodeName(v js.Value) string {
	return strings.ToLower(v.Get("nodeName").String())
}

func ConvertElement(tar, src documentElement) error {
	return src.setValue(tar.getValue())
}
