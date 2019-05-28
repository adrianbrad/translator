package translator

import (
	"strconv"

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

func (t *Translator) StoreTranslation(languageFrom string, textFrom string, languageTo string, textTo string) (err error) {
	languageTags, err := parseLanguageTags(languageFrom, languageTo)
	if err != nil {
		return
	}

	err = t.TranslationDAO.StoreTranslation(languageTags[0], textFrom, languageTags[1], textTo)
	return
}

func (t *Translator) GetTranslation(languageFrom string, textFrom string, languageTo string) (textTo string, err error) {
	languageTags, err := parseLanguageTags(languageFrom, languageTo)
	if err != nil {
		return
	}

	textTo, err = t.TranslationDAO.GetTranslation(languageTags[0], textFrom, languageTags[1])
	if err != nil {
		err = newLanguageTagParsingError(err.Error(), "languageFrom")
	}

	return
}

func parseLanguageTags(stringTags ...string) (languageTags []language.Tag, err error) {
	languageTags = make([]language.Tag, 0, len(stringTags))
	var languageTag language.Tag
	for index, tag := range stringTags {
		languageTag, err = language.Parse(tag)
		if err != nil {
			err = newLanguageTagParsingError(err.Error(), strconv.Itoa(index+1))
			return
		}

		languageTags = append(languageTags, languageTag)
	}
	return
}
