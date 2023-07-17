package utils

import (
	"CHGPU_T_BOT/lib/domain"
	"CHGPU_T_BOT/lib/state"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func SendTextMessage(bot *telego.Bot, update telego.Update, text string, menu telego.ReplyMarkup) (*telego.Message, error) {
	msg, err := bot.SendMessage(
		tu.Message(tu.ID(update.Message.Chat.ID), text).
			WithParseMode("html").
			WithReplyMarkup(menu).
			WithDisableWebPagePreview().
			WithProtectContent(),
	)
	switch msg.Text {
	case domain.Bachelor, domain.Masters, domain.PostGrad, domain.Secondary, domain.EduLevels:
		state.SetState(update.Message.Chat.ID, msg.Text)
	}

	return msg, err
}

func SendDocument(bot *telego.Bot, update telego.Update, text string, fileURL string, menu telego.ReplyMarkup) (*telego.Message, error) {
	document := tu.Document(
		tu.ID(update.Message.Chat.ID),
		tu.FileFromURL(fileURL),
	).WithCaption(text).WithReplyMarkup(menu).WithParseMode("html")
	msg, err := bot.SendDocument(document)
	return msg, err
}
