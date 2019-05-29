// +build db cli

package translatordbcli

import (
	"translator"
	"translator/internal/cmd/internal"
	"translator/internal/dao"
	"translator/internal/views"
)

func Run() {
	dbHost, dbPort, dbUser, dbPass := internal.GetDBCredentialsEnvVar()

	psqlDAO := dao.NewPSQLDao(dbHost, dbPort, dbUser, dbPass, "translator")

	t := translator.New(psqlDAO)

	cli := views.NewCLI(t)

	internal.RunCLI(cli)
}
