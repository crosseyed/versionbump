package internal

import (
	"fmt"
	"os"
)

const (
	Default = iota // Default
	Testing        // Testing
)

// Assigned in main
var App AppHelper

// Helper for exits and errors.
// mode can be one of Default or Testing.
type AppHelper struct {
	mode int
}

// Exit code wrapper.
// Used for unit testing.
func (a *AppHelper) Exit(code int) {
	for {
		switch a.mode {
		case Default:
			os.Exit(code)
		case Testing:
			panic(fmt.Sprintf("Exit(%d)", code))
		default:
			// Unknown mode switching to default
			a.mode = Default
		}
	}
}

// Check if an error was generated and run a callback.
// If err is nil the function returns.
//
// Print a standard error and exit:
//  a := AppHelper{}
//  err := errors.New("More cowbell")
//  a.Errors(err, nil)
//
// Custom code on error:
//  errcnt := 0
//  efunc := func(err error) {
//    e := fmt.Errorf("Got error: %v", err)
//    log.Printf("%v", e)
//    errcnt = errcnt + 1
//  }
//  err := SomeFunction()
//  a.Errors(err, efunc)
//  log.Printf("Error count: %d", errcnt)
func (a *AppHelper) Errors(err error, call func(err error)) {
	if err == nil {
		return
	}
	if call != nil {
		call(err)
	} else {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", err)
		a.Exit(-1)
	}
}
