package main

import (
	"os"
	"translator/internal/cmd/translatormemorycli"
)

func main() {
	translatormemorycli.Run(os.Stdin)
}
