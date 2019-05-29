// +build memory web

package translatormemoryweb

import (
	"translator"
	"translator/internal/cmd/internal"
	"translator/internal/dao"
	"translator/internal/views"
)

func Run() {
	memoryDAO := dao.NewMemoryDAO()

	t := translator.New(memoryDAO)

	webView := views.NewWebView(t)

	internal.RunWeb(webView)
}
