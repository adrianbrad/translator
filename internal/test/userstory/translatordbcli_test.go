package userstory

import (
	"testing"
	"os"
	"translator/internal/cmd/translatordbcli"
)
func TestDBCLI(t *testing.T) {
	f, err := os.Open("userstory.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	a := &printRead{f}
	translatordbcli.Run(a)
}