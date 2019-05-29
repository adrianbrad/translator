package userstory

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
	"translator/internal/cmd/translatordbweb"
	"translator/internal/views"
)

func TestDBWeb(t *testing.T) {
	go translatordbweb.Run()

	//waiting for db connection
	time.Sleep(1 * time.Second)

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
		if err != nil {
			panic(err)
		}
		if i.TextTo == "" {
			req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/%s/%s/%s", i.LanguageFrom, i.TextFrom, i.LanguageTo), nil)
			fmt.Printf("Sending GET-REQUEST{%s(%s)->%s}\n", i.LanguageFrom, i.TextFrom, i.LanguageTo)
			r, err := c.Do(req)
			if err != nil {
				panic(err)
			}

			var body map[string]string
			json.NewDecoder(r.Body).Decode(&body)
			fmt.Printf("RESPONSE{%s}\n", body["translation"])
		} else {
			body := strings.NewReader(fmt.Sprintf(`{"translation": "%s"}`, i.TextTo))
			req, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/%s/%s/%s", i.LanguageFrom, i.TextFrom, i.LanguageTo), body)
			fmt.Printf("Sending STORE-REQUEST{%s(%s)->%s(%s)}\n", i.LanguageFrom, i.TextFrom, i.LanguageTo, i.TextTo)
			r, err := c.Do(req)
			if err != nil {
				panic(err)
			}
			fmt.Printf("RESPONSE{%s}\n", r.Status)
		}
	}
}
