package dao

import (
	"fmt"

	"golang.org/x/text/language"
)

type noEntryForTheProvidedLanguageError struct {
	providedLanguage language.Tag
}

func newNoEntryForTheProvidedLanguageError(providedLanguage language.Tag) *noEntryForTheProvidedLanguageError {
	return &noEntryForTheProvidedLanguageError{
		providedLanguage: providedLanguage,
	}
}

func (e *noEntryForTheProvidedLanguageError) Error() string {
	return fmt.Sprintf("No entry found for the %s language", e.providedLanguage)
}

type noEntryForTheProvidedWordInProvidedLanguage struct {
	providedWord     string
	providedLanguage language.Tag
}

func newNoEntryForTheProvidedWordInProvidedLanguage(providedWord string, providedLanguage language.Tag) *noEntryForTheProvidedWordInProvidedLanguage {
	return &noEntryForTheProvidedWordInProvidedLanguage{
		providedWord:     providedWord,
		providedLanguage: providedLanguage,
	}
}

func (e *noEntryForTheProvidedWordInProvidedLanguage) Error() string {
	return fmt.Sprintf("No entry found for the %s word in provided language %s", e.providedWord, e.providedLanguage)
}

type noEntryForTheProvidedWordInRequestedLanguage struct {
	providedWord      string
	requestedLanguage language.Tag
}

func newNoEntryForTheProvidedWordInRequestedLanguage(providedWord string, requestedLanguage language.Tag) *noEntryForTheProvidedWordInRequestedLanguage {
	return &noEntryForTheProvidedWordInRequestedLanguage{
		providedWord:      providedWord,
		requestedLanguage: requestedLanguage,
	}
}

func (e *noEntryForTheProvidedWordInRequestedLanguage) Error() string {
	return fmt.Sprintf("No entry found for the %s word in requested language %s", e.providedWord, e.requestedLanguage)
}
