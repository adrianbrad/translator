package userstory

import (
	"testing"
	"translator/internal/cmd/translatormemoryweb"
	"net/http"
	"fmt"
	"translator/internal/views"
	"os"
	"bufio"
	"encoding/json"
	"strings"
)

func TestMemWeb(t *testing.T) {
	go translatormemoryweb.Run()
	c := &http.Client{}

	f, err := os.Open("userstory.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
    for {
        line, _, err := rd.ReadLine()
        if err != nil {
            return
        }
		i, err := views.DecodeIntoInput(string(line))
		if err != nil{
			panic(err)
		}
		if i.TextTo == "" {
			req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/%s/%s/%s", i.LanguageFrom, i.TextFrom, i.LanguageTo), nil)
			fmt.Printf("Sending GET-REQUEST{%s(%s)->%s}\n", i.LanguageFrom, i.TextFrom, i.LanguageTo)
			r, _ := c.Do(req)
			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body)
			fmt.Printf("RESPONSE{%s}\n",body["translation"])
		} else {
			body := strings.NewReader(fmt.Sprintf(`{"translation": "%s"}`, i.TextTo))
			req, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/%s/%s/%s", i.LanguageFrom, i.TextFrom, i.LanguageTo), body)
			fmt.Printf("Sending STORE-REQUEST{%s(%s)->%s(%s)}\n", i.LanguageFrom, i.TextFrom, i.LanguageTo, i.TextTo)
			r, _ := c.Do(req)
			fmt.Printf("RESPONSE{%s}\n",r.Status)

		}
	}
}
