package specialist

import (
	"CHGPU_T_BOT/lib/utils"

	"github.com/mymmrac/telego"
)

func SendSpecialistMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Задайте вопрос специалисту</b>"
	menu := getMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}
