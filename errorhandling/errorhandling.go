package errorhandling

import (
	"testing"
	"fmt"
)

// set via -ldflags "-X github.com/FactomProject/electiontesting/errorhandling.ErrorMode=debug" on the build line
var ErrorMode string // "" is production, "testing" is running a go test, "debug" is development
var T *testing.T     // Should be set by all tests first

func HandleError(note string) {
	switch ErrorMode {
	case "":
		_, err := fmt.Print(note)
		if err != nil {
			panic(err)
		}
	case "debug":
		panic(note)
	case "testing":
		if (T != nil) {
			T.Error(note)
		} else {
			panic("Unset testing: " + note)
		}
	}
}

func HandleFatal(note string) {
	switch ErrorMode {
	case "":
		_, err := fmt.Print(note)
		if err != nil {
			panic(err)
		}
	case "debug":
		panic(note)
	case "testing":
		if (T != nil) {
			T.Fatal(note)
		} else {
			panic("Unset testing: " + note)
		}
	}
}

func HandleErrorf(format string, a ...interface{}) {
	switch ErrorMode {
	case "":
		_, err := fmt.Printf(format, a ...)
		if err != nil {
			panic(err)
		}
	case "debug":
		panic(fmt.Sprintf(format, a))
	case "testing":
		if (T != nil) {
			T.Errorf(format, a)
		} else {
			panic("Unset testing: " + fmt.Sprintf(format, a))
		}
	}
}

func HandleFatalf(format string, a ...interface{}) {
	switch ErrorMode {
	case "":
		_, err := fmt.Printf(format, a...)
		if err != nil {
			panic(err)
		}
	case "debug":
		panic(fmt.Sprintf(format, a...))
	case "testing":
		if (T != nil) {
			T.Fatalf(format, a...)
		} else {
			panic("Unset testing: " + fmt.Sprintf(format, a))
		}
	}
}