package translator

import "fmt"

type languageTagParsingError struct {
	err           string // error description
	parameterName string
}

func newLanguageTagParsingError(err string, parameterName string) *languageTagParsingError {
	return &languageTagParsingError{
		err:           err,
		parameterName: parameterName,
	}
}

// satisfying error interface
func (e *languageTagParsingError) Error() string {
	return fmt.Sprintf("Error while parsing language tag %s: %s", e.parameterName, e.err)
}
