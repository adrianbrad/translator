package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestPSQLDao_UserStory_Success(t *testing.T) {
	d := NewPSQLDao("localhost", "5432", "admin", "admin", "translator")

	d.db.Exec(`TRUNCATE translations_text CASCADE`)
	d.db.Exec(`TRUNCATE translations CASCADE`)
	d.db.Exec(`ALTER SEQUENCE translations_translation_id_seq RESTART WITH 1;`)

	d.StoreTranslation(language.German, "catze", language.Spanish, "gato")
	d.StoreTranslation(language.Spanish, "gato", language.English, "cat")
	d.StoreTranslation(language.English, "cat", language.French, "chat")

	text, err := d.GetTranslation(language.English, "cat", language.German)
	assert.Nil(t, err)
	assert.Equal(t, "catze", text)

	text, err = d.GetTranslation(language.German, "catze", language.French)
	assert.Nil(t, err)
	assert.Equal(t, "chat", text)

	d.StoreTranslation(language.English, "brad", language.Spanish, "bradique")
	d.StoreTranslation(language.English, "brad", language.Italian, "bradini")
	d.StoreTranslation(language.English, "brad", language.German, "bradther")

	text, err = d.GetTranslation(language.German, "bradther", language.Spanish)
	assert.Nil(t, err)
	assert.Equal(t, "bradique", text)
}

func TestPSQLDao_StoreTranslation_Success(t *testing.T) {
	d := NewPSQLDao("localhost", "5432", "admin", "admin", "translator")

	d.db.Exec(`TRUNCATE translations_text CASCADE`)
	d.db.Exec(`TRUNCATE translations CASCADE`)
	d.db.Exec(`ALTER SEQUENCE translations_translation_id_seq RESTART WITH 1;`)

	d.StoreTranslation(language.German, "catze", language.Spanish, "gato")
	var id int
	err := d.db.QueryRow(`
SELECT translation_id
FROM translations
	`).Scan(&id)
	assert.Nil(t, err)
	assert.Equal(t, 1, id)

	r, err := d.db.Query(`
SELECT translation_id, language_tag, text
FROM translations_text
ORDER BY language_tag;
		`)
	assert.Nil(t, err)
	var count, tid int
	var tag, text string

	for r.Next() {
		count++
		r.Scan(&tid, &tag, &text)
		if count == 1 {
			assert.Equal(t, 1, tid)
			assert.Equal(t, "de", tag)
			assert.Equal(t, "catze", text)
		} else {
			assert.Equal(t, 1, tid)
			assert.Equal(t, "es", tag)
			assert.Equal(t, "gato", text)
		}
	}
	assert.Equal(t, 2, count)
}

func TestPSQLDao_StoreTranslation_Duplicate(t *testing.T) {
	d := NewPSQLDao("localhost", "5432", "admin", "admin", "translator")

	d.db.Exec(`TRUNCATE translations_text CASCADE`)
	d.db.Exec(`TRUNCATE translations CASCADE`)
	d.db.Exec(`ALTER SEQUENCE translations_translation_id_seq RESTART WITH 1;`)
	err := d.StoreTranslation(language.German, "catze", language.Spanish, "gato")
	assert.Nil(t, err)
	err = d.StoreTranslation(language.German, "catze", language.Spanish, "gato")
	assert.Nil(t, err)
}

func TestPSQLDao_GetTranslation_Success(t *testing.T) {
	d := NewPSQLDao("localhost", "5432", "admin", "admin", "translator")

	d.db.Exec(`TRUNCATE translations_text CASCADE`)
	d.db.Exec(`TRUNCATE translations CASCADE`)
	d.db.Exec(`ALTER SEQUENCE translations_translation_id_seq RESTART WITH 1;`)

	d.StoreTranslation(language.German, "catze", language.Spanish, "gato")

	word, err := d.GetTranslation(language.German, "catze", language.Spanish)
	assert.Nil(t, err)
	assert.Equal(t, "gato", word)
}

func TestPSQLDao_GetTranslation_NotFound(t *testing.T) {
	d := NewPSQLDao("localhost", "5432", "admin", "admin", "translator")

	d.db.Exec(`TRUNCATE translations_text CASCADE`)
	d.db.Exec(`TRUNCATE translations CASCADE`)
	d.db.Exec(`ALTER SEQUENCE translations_translation_id_seq RESTART WITH 1;`)

	word, err := d.GetTranslation(language.German, "not found", language.Spanish)
	assert.Equal(t, "Word not found", err.Error())
	assert.Empty(t, word)
}
