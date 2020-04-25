package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

// Error implements the error interface and simply returns feedback for
// logging to the rtc log file
type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrNoAccount = Error("User has not created an account yet and must do so before they continue to use this application.")
	ErrNoKey     = Error("User has not created an api key yet and must do so before they continue to use this application.")
)

type FatalErrorCode int

const (
	SetLoggerFail    FatalErrorCode = 1
	ConfigFileInit   FatalErrorCode = 2
	GetRequestFailed FatalErrorCode = 3
	URLParseFailed   FatalErrorCode = 4
	CSVParseFailed   FatalErrorCode = 5
)

// InternalFatalError will report the error code to the user and return such error
func InternalFatalError(fatal FatalErrorCode) {

	errorMessage := fmt.Sprintf("An internal error has occured. Exit code: %d", fatal)
	color.Red.Println(errorMessage)
	os.Exit(int(fatal))
}
