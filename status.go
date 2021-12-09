package holiday

import (
	"fmt"

	"github.com/mgutz/ansi"
)

type DoctorItemStatus int

const (
	PASS DoctorItemStatus = iota
	WARN
	FAIL
)

var (
	green  = ansi.ColorCode("green+h")
	yellow = ansi.ColorCode("yellow+h")
	red    = ansi.ColorCode("red+h")
	white  = ansi.ColorCode("white+h")
	reset  = ansi.ColorCode("reset")
)

func SprintStatusMessage(status DoctorItemStatus, indent string, message string) string {
	var statusStr, colorPrefix string

	switch status {
	case PASS:
		statusStr = "✓"
		colorPrefix = green
		break
	case WARN:
		statusStr = "!"
		colorPrefix = yellow
		break
	case FAIL:
		statusStr = "✗"
		colorPrefix = red
		break
	}

	return fmt.Sprintf("%s%s[%s]%s %s", indent, colorPrefix, statusStr, reset, message)
}
