// +build db cli

package translatordbcli

import (
	"io"
	"translator"
	"translator/internal/cmd/internal"
	"translator/internal/dao"
	"translator/internal/views"
)

func Run(r io.Reader) {
	dbHost, dbPort, dbUser, dbPass := internal.GetDBCredentialsEnvVar()

	psqlDAO := dao.NewPSQLDao(dbHost, dbPort, dbUser, dbPass, "translator")

	t := translator.New(psqlDAO)

	cli := views.NewCLI(t, r)

	internal.RunCLI(cli)
}
