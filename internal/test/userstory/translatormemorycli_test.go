// +build memory cli

package userstory

import (
	"os"
	"testing"
	"translator/internal/cmd/translatormemorycli"
)

func TestMemCLI(t *testing.T) {
	f, err := os.Open("userstory.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	a := &printRead{f}
	translatormemorycli.Run(a)
}
