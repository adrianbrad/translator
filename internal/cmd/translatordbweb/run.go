// +build db web

package translatordbweb

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

	webView := views.NewWebView(t)

	internal.RunWeb(webView)
}
