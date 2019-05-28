package views

type translator interface {
	GetTranslation(languageFrom string, textFrom string, languageTo string) (textTo string, err error)
	StoreTranslation(languageFrom string, textFrom string, languageTo string, textTo string) (err error)
}
