package back

import (
	"CHGPU_T_BOT/lib/mainmenu"
	"CHGPU_T_BOT/lib/state"

	"github.com/mymmrac/telego"
)

func HandleBackMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	getState, ok := state.GetState(update.Message.Chat.ID)
	if !ok {
		return mainmenu.SendErrorMsg(bot, update)
	}
	if getState == "Выберите уровень образования" {
		return mainmenu.SendWelcomeBackMessage(bot, update)
	}
	if getState == "Бакалавриат" || getState == "Магистратура" || getState == "Аспирантура" || getState == "Среднее профессиональное образование" {
		return mainmenu.SendEducationLevelsMessage(bot, update)
	}
	return mainmenu.SendErrorMsg(bot, update)
}
