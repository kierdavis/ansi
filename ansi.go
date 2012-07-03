package ansi

import (
	"fmt"
	"io"
	"sync"
)

func init() {
	go func() {
		// On program exit, ensure that we reset the screen.
		defer func() {
			fmt.Print("\x1b[0m")
		}()

		<-make(chan struct{})
	}()
}

var Mutex sync.Mutex
var UseMutex bool

type Color uint

type Attribute struct {
	Attr uint
	FG   Color
	BG   Color
}

const (
	Bold uint = 1 << iota
	Underline
	Blink
	Inverse
)

const (
	None    Color = 0 // This should be black, but the default (0) needs to be "no color"
	Red     Color = 1
	Green   Color = 2
	Yellow  Color = 3
	Blue    Color = 4
	Magenta Color = 5
	Cyan    Color = 6
	White   Color = 7
	Black   Color = 8 // Luckily, color % 8 will turn this 8 into a 0 correctly
)

func SAttrOn(attr Attribute) (s string) {
	if attr.FG != None {
		s += fmt.Sprintf("\x1b[3%dm", attr.FG%8)
	}

	if attr.BG != None {
		s += fmt.Sprintf("\x1b[4%dm", attr.BG%8)
	}

	if attr.Attr&Bold != 0 {
		s += "\x1b[1m"
	}

	if attr.Attr&Underline != 0 {
		s += "\x1b[4m"
	}

	if attr.Attr&Blink != 0 {
		s += "\x1b[5m"
	}

	if attr.Attr&Inverse != 0 {
		s += "\x1b[7m"
	}

	return s
}

func FAttrOn(w io.Writer, attr Attribute) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Fprint(w, SAttrOn(attr))
}

func AttrOn(attr Attribute) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Print(SAttrOn(attr))
}

func SAttrOff(attr Attribute) (s string) {
	if attr.FG != None {
		s += "\x1b[39m"
	}

	if attr.BG != None {
		s += "\x1b[49m"
	}

	if attr.Attr&Bold != 0 {
		s += "\x1b[22m"
	}

	if attr.Attr&Underline != 0 {
		s += "\x1b[24m"
	}

	if attr.Attr&Blink != 0 {
		s += "\x1b[25m"
	}

	if attr.Attr&Inverse != 0 {
		s += "\x1b[27m"
	}

	return s
}

func FAttrOff(w io.Writer, attr Attribute) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Fprint(w, SAttrOff(attr))
}

func AttrOff(attr Attribute) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Print(SAttrOff(attr))
}

func Sprint(attr Attribute, a ...interface{}) (s string) {
	return SAttrOn(attr) + fmt.Sprint(a...) + SAttrOff(attr)
}

func Sprintln(attr Attribute, a ...interface{}) (s string) {
	return SAttrOn(attr) + fmt.Sprintln(a...) + SAttrOff(attr)
}

func Sprintf(attr Attribute, format string, a ...interface{}) (s string) {
	return SAttrOn(attr) + fmt.Sprintf(format, a...) + SAttrOff(attr)
}

func Print(attr Attribute, a ...interface{}) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Print(Sprint(attr, a...))
}

func Println(attr Attribute, a ...interface{}) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Print(Sprintln(attr, a...))
}

func Printf(attr Attribute, format string, a ...interface{}) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Print(Sprintf(attr, format, a...))
}

func Fprint(w io.Writer, attr Attribute, a ...interface{}) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Fprint(w, Sprint(attr, a...))
}

func Fprintln(w io.Writer, attr Attribute, a ...interface{}) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Fprint(w, Sprintln(attr, a...))
}

func Fprintf(w io.Writer, attr Attribute, format string, a ...interface{}) (n int, err error) {
	if UseMutex {
		Mutex.Lock()
		defer Mutex.Unlock()
	}

	return fmt.Fprint(w, Sprintf(attr, format, a...))
}
