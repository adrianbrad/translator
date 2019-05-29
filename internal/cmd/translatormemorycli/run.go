// +build memory cli

package translatormemorycli

import (
	"translator"
	"translator/internal/cmd/internal"
	"translator/internal/dao"
	"translator/internal/views"
)

func Run() {
	memoryDAO := dao.NewMemoryDAO()

	t := translator.New(memoryDAO)

	cli := views.NewCLI(t)

	internal.RunCLI(cli)
}
