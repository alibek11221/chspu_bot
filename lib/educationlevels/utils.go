package educationlevels

import (
	"CHGPU_T_BOT/lib/domain"
	"CHGPU_T_BOT/lib/educationlevels/bachelor"
	"CHGPU_T_BOT/lib/educationlevels/masters"
	"CHGPU_T_BOT/lib/educationlevels/post_grad"
	"CHGPU_T_BOT/lib/educationlevels/secondary"
	"CHGPU_T_BOT/lib/mainmenu"
	"CHGPU_T_BOT/lib/state"

	"github.com/mymmrac/telego"
)

func EntranceRules(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text, ok := state.GetState(update.Message.Chat.ID)
	if !ok {
		return mainmenu.SendErrorMsg(bot, update)
	}
	if text == domain.Bachelor {
		return bachelor.SendApplicationRulesMessage(bot, update)
	}
	if text == domain.Masters {
		return masters.SendApplicationRulesMessage(bot, update)
	}
	if text == domain.PostGrad {
		return post_grad.SendApplicationRulesMessage(bot, update)
	}
	if text == domain.Secondary {
		return secondary.SendApplicationRulesMessage(bot, update)
	}
	return mainmenu.SendErrorMsg(bot, update)
}

func MinimalPoints(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text, ok := state.GetState(update.Message.Chat.ID)
	if !ok {
		return mainmenu.SendErrorMsg(bot, update)
	}
	if text == domain.Bachelor {
		return bachelor.SendMinimalPointsMessage(bot, update)
	}
	if text == domain.Masters {
		return masters.SendMinimalPointsMessage(bot, update)
	}
	if text == domain.PostGrad {
		return post_grad.SendMinimalPointsMessage(bot, update)
	}
	if text == domain.Secondary {
		return secondary.SendMinimalPointsMessage(bot, update)
	}
	return mainmenu.SendErrorMsg(bot, update)
}

func List(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text, ok := state.GetState(update.Message.Chat.ID)
	if !ok {
		return mainmenu.SendErrorMsg(bot, update)
	}
	if text == domain.Bachelor {
		return bachelor.SendListMessage(bot, update)
	}
	if text == domain.Masters {
		return masters.SendListMessage(bot, update)
	}
	if text == domain.PostGrad {
		return post_grad.SendListMessage(bot, update)
	}
	if text == domain.Secondary {
		return secondary.SendListMessage(bot, update)
	}
	return mainmenu.SendErrorMsg(bot, update)
}
func IndividualList(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text, ok := state.GetState(update.Message.Chat.ID)
	if !ok {
		return mainmenu.SendErrorMsg(bot, update)
	}
	if text == domain.Bachelor {
		return bachelor.SendIndividualListMessage(bot, update)
	}
	if text == domain.Masters {
		return masters.SendIndividualListMessage(bot, update)
	}
	if text == domain.PostGrad {
		return post_grad.SendIndividualListMessage(bot, update)
	}
	if text == domain.Secondary {
		return secondary.SendIndividualListMessage(bot, update)
	}
	return mainmenu.SendErrorMsg(bot, update)
}

func Exception(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text, ok := state.GetState(update.Message.Chat.ID)
	if !ok {
		return mainmenu.SendErrorMsg(bot, update)
	}
	if text == domain.Bachelor {
		return bachelor.SendExceptionMessage(bot, update)
	}
	if text == domain.Masters {
		return masters.SendExceptionMessage(bot, update)
	}
	if text == domain.PostGrad {
		return post_grad.SendExceptionMessage(bot, update)
	}
	if text == domain.Secondary {
		return secondary.SendExceptionMessage(bot, update)
	}
	return mainmenu.SendErrorMsg(bot, update)
}
func Periods(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text, ok := state.GetState(update.Message.Chat.ID)
	if !ok {
		return mainmenu.SendErrorMsg(bot, update)
	}
	if text == domain.Bachelor {
		return bachelor.SendPeriodMessage(bot, update)
	}
	if text == domain.Masters {
		return masters.SendPeriodMessage(bot, update)
	}
	if text == domain.PostGrad {
		return post_grad.SendPeriodMessage(bot, update)
	}
	if text == domain.Secondary {
		return secondary.SendPeriodMessage(bot, update)
	}
	return mainmenu.SendErrorMsg(bot, update)
}
func HandleBackMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text, ok := state.GetState(update.Message.Chat.ID)
	if !ok {
		return mainmenu.SendErrorMsg(bot, update)
	}
	if text == domain.Bachelor {
		return SendBachelorMessage(bot, update)
	}
	if text == domain.Masters {
		return SendMastersMessage(bot, update)
	}
	if text == domain.PostGrad {
		return SendPostGradMessage(bot, update)
	}
	if text == domain.Secondary {
		return SendSecondaryMessage(bot, update)
	}
	return mainmenu.SendErrorMsg(bot, update)
}
