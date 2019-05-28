package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
)

type Web struct {
	t translator
}

func NewWebView(t translator) *Web {
	return &Web{
		t: t,
	}
}

func (wi *Web) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//remove the first slash
	p := r.URL.Path[1:]
	// remove trailing slash if present
	p = path.Clean(p)
	urlParameters := strings.Split(p, "/")

	if len(urlParameters) != 3 {
		http.Error(w,
			"Url should have the following format: /{languageFrom}/{text}/{languageTo}",
			http.StatusBadRequest,
		)
		return
	}

	switch r.Method {
	case http.MethodGet:
		textTo, err := wi.t.GetTranslation(urlParameters[0], urlParameters[1], urlParameters[2])
		if err != nil {
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"translation":"%s"}`, textTo)
		return

	case http.MethodPost:
		var body map[string]string
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		err = wi.t.StoreTranslation(urlParameters[0], urlParameters[1], urlParameters[2], body["translation"])
		if err != nil {
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
