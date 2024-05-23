package messagesHandlers

import (
	"github.com/Monska85/telegram-gateway/lib/logHelper"
	"github.com/Monska85/telegram-gateway/lib/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var logger = logHelper.GetInstance()

func RouteNewMessage(update tgbotapi.Update) {
	logger.Out(utils.LogInfo, "Routing new message", "update", update)
}
