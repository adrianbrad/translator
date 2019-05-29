// +build memory cli

package translatormemorycli

import (
	"io"
	"translator"
	"translator/internal/cmd/internal"
	"translator/internal/dao"
	"translator/internal/views"
)

func Run(r io.Reader) {
	memoryDAO := dao.NewMemoryDAO()

	t := translator.New(memoryDAO)

	cli := views.NewCLI(t, r)

	internal.RunCLI(cli)
}
