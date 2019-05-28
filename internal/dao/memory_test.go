package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

// remove the locks from `StoreTranslation` if you want to get a data race warning
// this could happen while serving over http as the `ServeHTTP` method is run in a separate go routine for every request
func TestMemoryDAO_RaceCondition(t *testing.T) {
	mDAO := NewMemoryDAO()

	for i := 1; i <= 100; i++ {
		go func() {
			mDAO.StoreTranslation(language.German, "catze", language.Spanish, "gato")
		}()
	}
}

func TestMemoryDAO_StoreTranslation_Success(t *testing.T) {
	mDAO := NewMemoryDAO()

	mDAO.StoreTranslation(language.German, "catze", language.Spanish, "gato")
	mDAO.StoreTranslation(language.Spanish, "gato", language.English, "cat")
	mDAO.StoreTranslation(language.English, "cat", language.French, "chat")

	//total number of languages (german, spanish, english, french)
	assert.Equal(t, len(mDAO.translations), 4)

	assert.Equal(t, mDAO.translations[language.German]["catze"], mDAO.translations[language.Spanish]["gato"])
	assert.Equal(t, mDAO.translations[language.German]["catze"], mDAO.translations[language.English]["cat"])
	assert.Equal(t, mDAO.translations[language.German]["catze"], mDAO.translations[language.French]["chat"])

	assert.Equal(t, mDAO.translations[language.Spanish]["gato"], mDAO.translations[language.German]["catze"])
	assert.Equal(t, mDAO.translations[language.Spanish]["gato"], mDAO.translations[language.English]["cat"])
	assert.Equal(t, mDAO.translations[language.Spanish]["gato"], mDAO.translations[language.French]["chat"])

	assert.Equal(t, mDAO.translations[language.English]["cat"], mDAO.translations[language.German]["catze"])
	assert.Equal(t, mDAO.translations[language.English]["cat"], mDAO.translations[language.Spanish]["gato"])
	assert.Equal(t, mDAO.translations[language.English]["cat"], mDAO.translations[language.French]["chat"])

	assert.Equal(t, mDAO.translations[language.French]["chat"], mDAO.translations[language.German]["catze"])
	assert.Equal(t, mDAO.translations[language.French]["chat"], mDAO.translations[language.Spanish]["gato"])
	assert.Equal(t, mDAO.translations[language.French]["chat"], mDAO.translations[language.English]["cat"])

	mDAO.StoreTranslation(language.English, "brad", language.Spanish, "bradique")
	mDAO.StoreTranslation(language.English, "brad", language.Italian, "bradini")
	mDAO.StoreTranslation(language.English, "brad", language.German, "bradther")

	//total number of languages (german, spanish, english, french, italian)
	assert.Equal(t, len(mDAO.translations), 5)

	assert.Equal(t, mDAO.translations[language.English]["brad"], mDAO.translations[language.German]["bradther"])
	assert.Equal(t, mDAO.translations[language.English]["brad"], mDAO.translations[language.Spanish]["bradique"])
	assert.Equal(t, mDAO.translations[language.English]["brad"], mDAO.translations[language.Italian]["bradini"])
}

func TestMemoryDAO_GetTranslation_ProvidedLanguageNotFound(t *testing.T) {
	mDAO := NewMemoryDAO()

	wordTo, err := mDAO.GetTranslation(language.English, "", language.Tag{})
	assert.Empty(t, wordTo)

	assert.NotNil(t, err)
	assert.Equal(t, err, newNoEntryForTheProvidedLanguageError(language.English))
}

func TestMemoryDAO_GetTranslation_ProvidedWordNotFoundInProvidedLanguage(t *testing.T) {
	mDAO := NewMemoryDAO()
	mDAO.translations[language.English] = nil

	wordTo, err := mDAO.GetTranslation(language.English, "test", language.Tag{})
	assert.Empty(t, wordTo)

	assert.NotNil(t, err)
	assert.Equal(t, err, newNoEntryForTheProvidedWordInProvidedLanguage("test", language.English))
}

func TestMemoryDAO_GetTranslation_ProvidedWordNotFoundInRequestedLanguage(t *testing.T) {
	mDAO := NewMemoryDAO()
	mDAO.translations[language.English] = map[string]translationHolder{
		"test": nil,
	}

	wordTo, err := mDAO.GetTranslation(language.English, "test", language.Spanish)
	assert.Empty(t, wordTo)

	assert.NotNil(t, err)
	assert.Equal(t, err, newNoEntryForTheProvidedWordInRequestedLanguage("test", language.Spanish))
}
