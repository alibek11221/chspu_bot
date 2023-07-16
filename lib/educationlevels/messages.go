package educationlevels

import (
	"CHGPU_T_BOT/lib/utils"
	"github.com/mymmrac/telego"
)

func SendBachelorMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Бакалавриат</b>"
	menu := getMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}

func SendMastersMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Магистратура</b>"
	menu := getMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}

func SendPostGradMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Аспирантура</b>"
	menu := getMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}

func SendSecondaryMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Среднее профессиональное образование</b>"
	menu := getMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}
