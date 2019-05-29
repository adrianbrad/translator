package main

import (
	"os"
	"translator/internal/cmd/translatordbcli"
)

func main() {
	translatordbcli.Run(os.Stdin)
}
