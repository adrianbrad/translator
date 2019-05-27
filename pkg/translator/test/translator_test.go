package test

import (
	"testing"
	"translator/pkg/translator"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const malformedTag = "malformed tag"

func TestTranslator_GetTranslation_ErrorParsingLanguageTo(t *testing.T) {
	tr := translator.New(nil)

	textTo, err := tr.GetTranslation("en", "", malformedTag)

	assert.Empty(t, textTo)
	assert.NotNil(t, err)
	assert.Equal(t, "Error while parsing languageTo: language: tag is not well-formed", err.Error())
}

func TestTranslator_GetTranslation_ErrorParsingLanguageFrom(t *testing.T) {
	tr := translator.New(nil)

	textTo, err := tr.GetTranslation(malformedTag, "", "")

	assert.Empty(t, textTo)
	assert.NotNil(t, err)
	assert.Equal(t, "Error while parsing languageFrom: language: tag is not well-formed", err.Error())
}

func TestTranslator_GetTranslation_Success(t *testing.T) {
	tDAO := &translatorDAOMock{}
	tr := translator.New(tDAO)

	tDAO.On("GetTranslation", mock.Anything).Return("success", nil)

	textTo, err := tr.GetTranslation("ts", "test", "st")

	assert.Nil(t, err)
	assert.Equal(t, "success", textTo)
}

func TestTranslator_StoreTranslation_ErrorParsingLanguageFrom(t *testing.T) {
	tr := translator.New(nil)

	err := tr.StoreTranslation(malformedTag, "", "", "")

	assert.NotNil(t, err)
	assert.Equal(t, "Error while parsing languageFrom: language: tag is not well-formed", err.Error())
}

func TestTranslator_StoreTranslation_ErrorParsingLanguageTo(t *testing.T) {
	tr := translator.New(nil)

	err := tr.StoreTranslation("en", "", malformedTag, "")

	assert.NotNil(t, err)
	assert.Equal(t, "Error while parsing languageTo: language: tag is not well-formed", err.Error())
}

func TestTranslator_StoreTranslation_Success(t *testing.T) {
	tDAO := &translatorDAOMock{}
	tr := translator.New(tDAO)

	tDAO.On("StoreTranslation", mock.Anything).Return(nil)

	err := tr.StoreTranslation("ts", "test", "st", "tset")

	assert.Nil(t, err)
}
