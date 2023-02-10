package domgo

import (
	"fmt"
	"image/color"
)

func StyleColor(col color.RGBA) string {
	return fmt.Sprintf("#%02x%02x%02x%02x", col.R, col.G, col.B, col.A)
}
