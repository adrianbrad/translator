package translator

import "golang.org/x/text/language"

type TranslationDAO interface {
	StoreTranslation(languageFrom language.Tag, textFrom string, languageTo language.Tag, textTo string) (err error)

	GetTranslation(languageFrom language.Tag, textFrom string, languagTo language.Tag) (text string, err error)
}
