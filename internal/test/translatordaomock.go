package test

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/text/language"
)

type translatorDAOMock struct {
	mock.Mock
}

func (t *translatorDAOMock) StoreTranslation(languageFrom language.Tag, textFrom string, languageTo language.Tag, textTo string) (err error) {
	args := t.Called()
	return args.Error(0)
}

func (t *translatorDAOMock) GetTranslation(languageFrom language.Tag, textFrom string, languagTo language.Tag) (text string, err error) {
	args := t.Called()
	err = args.Error(1)
	text = args.Get(0).(string)
	return
}
