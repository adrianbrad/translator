package dao

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestPSQLDao_StoreTranslation_Success(t *testing.T) {
	d := NewPSQLDao("localhost", "5432", "admin", "admin", "translator")

	fmt.Println(d.StoreTranslation(language.German, "catze", language.Spanish, "gato"))
	d.StoreTranslation(language.Spanish, "gato", language.English, "cat")
	d.StoreTranslation(language.English, "cat", language.French, "chat")

	fmt.Println(d.GetTranslation(language.English, "cat", language.German))
}
