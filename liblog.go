package liblog

import (
	"fmt"
	"strings"
	"time"

	"github.com/victorgeel/libutils"
	"github.com/buger/goterm"
)

var (
	Colors = map[string]string{
		"B1": "\033[30;1m", "B2": "\033[30;2m",
		"R1": "\033[31;1m", "R2": "\033[31;2m",
		"G1": "\033[32;1m", "G2": "\033[32;2m",
		"Y1": "\033[33;1m", "Y2": "\033[33;2m",
		"B3": "\033[34;1m", "B4": "\033[34;2m",
		"M1": "\033[35;1m", "M2": "\033[35;2m",
		"C1": "\033[36;1m", "C2": "\033[36;2m",
		"W1": "\033[37;1m", "W2": "\033[37;2m",
		"CC": "\033[0m",
		"BG_BLACK":   "\033[0;40m",
		"BG_RED":     "\033[0;41m",
		"BG_GREEN":   "\033[0;42m",
		"BG_YELLOW":  "\033[0;43m",
		"BG_BLUE":    "\033[0;44m",
		"BG_MAGENTA": "\033[0;45m",
		"BG_CYAN":    "\033[0;46m",
		"BG_WHITE":   "\033[0;47m",
	}
)

func LimitMessageLength(message string, slice int) (string, string) {
	terminal_width := goterm.Width() - slice
	messages := []string{message, ""}

	if len(message) > terminal_width {
		messages[0] = message[:terminal_width]
		messages[1] = message[terminal_width:]
	}

	return messages[0], messages[1]
}

func Log(message string, prefix string) {
	fmt.Printf("%s%s%s%s%s", "\r", "\033[K", message, Colors["CC"], prefix)
}

func LogColor(message string, color string) {
	messages := strings.Split(message, "\n")

	for _, value := range messages {
		Log(color+value, "\n")
	}
}

func Header(messages []string, color string) {
	libutils.ClearScreen()

	LogColor(strings.Join(messages, "\n")+"\n", color)
}

func LogInfo(message string, info string, color string) {
	datetime := time.Now()
	LogColor(
		fmt.Sprintf("[%.2d:%.2d:%.2d]%[5]s %[4]s::%[5]s %[6]s%[7]s%[5]s %[4]s::%[5]s %[6]s%[8]s",
			datetime.Hour(), datetime.Minute(), datetime.Second(),
			Colors["B1"], Colors["CC"], color,
			info, message),
		color,
	)
}

func LogInfoSplit(message string, slice int, info string, color string) {
	var data string
	var i = 0

	for len(message) != 0 {
		data, message = LimitMessageLength(message, slice)
		if i == 0 {
			LogInfo(data, info, color)
		} else {
			LogColor(strings.Repeat(" ", slice)+data, color)
		}

		i++
	}
}

func LogKeyboardInterrupt() {
	LogInfo(
		"Keyboard Interrupt\n\n"+
			"|   Ctrl-C again if not exiting automaticly\n"+
			"|   Please wait...\n|\n",
		"INFO", Colors["R1"],
	)
}

func LogException(err error, info string) {
	LogInfo(fmt.Sprintf("Exception:\n\n|   %v\n|\n", err), info, Colors["R1"])
}

func LogReplace(message string, color string) {
	message, data := LimitMessageLength(message, 4)
	if len(data) != 0 {
		message = message + "..."
	}
	Log(color+message, "\r")
}
