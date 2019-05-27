package translator

import (
	"golang.org/x/text/language"
)

type Translator struct {
	TranslationDAO TranslationDAO
}

func New(translationDAO TranslationDAO) *Translator {
	return &Translator{
		TranslationDAO: translationDAO,
	}
}

func (t *Translator) StoreTranslation(
	languageFrom string,
	textFrom string,
	languageTo string,
	textTo string,
) (err error) {
	languageTagFrom, err := language.Parse(languageFrom)
	if err != nil {
		return newLanguageTagParsingError(err.Error(), "languageFrom")
	}

	languageTagTo, err := language.Parse(languageTo)
	if err != nil {
		return newLanguageTagParsingError(err.Error(), "languageTo")
	}

	err = t.TranslationDAO.StoreTranslation(languageTagFrom, textFrom, languageTagTo, textTo)
	return
}

func (t *Translator) GetTranslation(languageFrom string, textFrom string, languageTo string) (textTo string, err error) {
	languageTagFrom, err := language.Parse(languageFrom)
	if err != nil {
		err = newLanguageTagParsingError(err.Error(), "languageFrom")
		return
	}

	languageTagTo, err := language.Parse(languageTo)
	if err != nil {
		err = newLanguageTagParsingError(err.Error(), "languageTo")
		return
	}

	textTo, err = t.TranslationDAO.GetTranslation(languageTagFrom, textFrom, languageTagTo)
	if err != nil {
		err = newLanguageTagParsingError(err.Error(), "languageFrom")
	}

	return
}
