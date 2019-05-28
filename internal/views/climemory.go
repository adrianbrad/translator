package views

import (
	"fmt"
	"regexp"
	"strings"
)

type input struct {
	languageFrom string
	textFrom     string

	languageTo string
	textTo     string
}

func decodeIntoInput(s string) (i input, err error) {
	// (s(s),s(s))
	storeR := regexp.MustCompile(`(\()([a-z]*)(\()([a-z]*)(\))(,)([a-z]*)(\()([a-z]*)(\))(\))`)
	// (s(s),s)
	getR := regexp.MustCompile(`(\()([a-z]*)(\()([a-z]*)(\))(,)([a-z]*)(\))`)

	//beware, black magic
	if storeR.MatchString(s) {
		// s = (s(s),s(s))
		languageFrom := getTextBetweenTwoOpenRoundBrackets(s)
		s = s[len(languageFrom)+1:]
		// s = (s),s(s))

		textFrom := getTextBetweenOpenAndCloseRoundBrackets(s)
		// +3 comes from the open bracket, closed bracket and comma
		s = s[len(textFrom)+3:]
		// s = s(s))

		languageTo := s[:strings.Index(s, "(")]
		s = s[len(languageTo):]
		// s = (s))

		textTo := getTextBetweenOpenAndCloseRoundBrackets(s)

		i = input{
			languageFrom: languageFrom,
			textFrom:     textFrom,
			languageTo:   languageTo,
			textTo:       textTo,
		}
		return
	}

	if getR.MatchString(s) {
		// s = (s(s),s)
		languageFrom := getTextBetweenTwoOpenRoundBrackets(s)
		s = s[len(languageFrom)+1:]
		// s = (s),s)
		textFrom := getTextBetweenOpenAndCloseRoundBrackets(s)
		// +3 comes from the open bracket, closed bracket and comma
		s = s[len(textFrom)+3:]
		// s = s)

		languageTo := s[:strings.Index(s, ")")]

		i = input{
			languageFrom: languageFrom,
			textFrom:     textFrom,
			languageTo:   languageTo,
		}
		return
	}

	err = fmt.Errorf("Invalid input format")
	return
}

func getTextBetweenTwoOpenRoundBrackets(s string) string {
	return s[strings.Index(s, "(")+1 : strings.Index(s[strings.Index(s, "(")+1:], "(")+1]
}

func getTextBetweenOpenAndCloseRoundBrackets(s string) string {
	return s[strings.Index(s, "(")+1 : strings.Index(s[strings.Index(s, "(")+1:], ")")+1]
}

type CLIMemory struct {
	t translator
}

func NewCLIMemory(t translator) *CLIMemory {
	return &CLIMemory{
		t: t,
	}
}

func (c *CLIMemory) Intro() {
	fmt.Println("Hello!")
}

func (c *CLIMemory) Outro() {
	fmt.Println("\nGoodbye!")
}

func (c *CLIMemory) CallToAction() {
	fmt.Println("If you wish to store a translation please provide the input using the following format: (en(cat),ge(katze))")
	fmt.Println("If you wish to get a translation please provide input the using the following format: (en(cat),gb)")
}

func (c *CLIMemory) ResolveAction() {
	var cliInput string
	_, err := fmt.Scanf("%s\n", &cliInput)
	if err != nil {
		fmt.Printf("Error while scanning input: %s\n", err.Error())
		return
	}

	i, err := decodeIntoInput(cliInput)
	if err != nil {
		fmt.Printf("Error while decoding input: %s\n", err.Error())
		return
	}

	// GET translation
	if i.textTo == "" {
		textTo, err := c.t.GetTranslation(i.languageFrom, i.textFrom, i.languageTo)
		if err != nil {
			fmt.Printf("Error while getting translation: %s\n", err.Error())
			return
		}
		fmt.Println(textTo)
	} else
	// STORE translation

	{
		err := c.t.StoreTranslation(i.languageFrom, i.textFrom, i.languageTo, i.textTo)
		if err != nil {
			fmt.Printf("Error while storing translation: %s\n", err.Error())
			return
		}
	}
	return
}
