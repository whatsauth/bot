package gcf

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/whatsauth/bot"
)

func init() {
	functions.HTTP("WebHook", bot.PostWithDB)
}
