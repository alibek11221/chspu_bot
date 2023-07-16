package post_grad

import (
	"CHGPU_T_BOT/lib/domain"
	"CHGPU_T_BOT/lib/utils"

	"github.com/mymmrac/telego"
)

func SendApplicationRulesMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Каждый поступающий обязан внимательно ознакомиться с \n " +
		"Правилами приема в Университете</b>"
	menu := getMenu().ToReplyMarkup()
	return utils.SendDocument(bot, update, text, domain.PostGradFile, menu)
}
func SendMinimalPointsMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Вступительные испытания в аспирантуру проводятся по следующим дисциплинам: " +
		"\nИностранный язык" +
		"\nСпециальная дисциплина" +
		"\nДля успешной сдачи экзамена необходимо набрать от 40 баллов"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendListMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Информация доступна по <a href='https://chspu.ru/abitur/aspirant/#abitur_plan_priema'>ссылке</a>"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendIndividualListMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Информация о порядке учета индивидуальных достижений поступающих"
	menu := getMenu().ToReplyMarkup()

	return utils.SendDocument(bot, update, text, domain.PostGradIndividualFile, menu)
}

func SendExceptionMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Особые права и преимущества отсутствуют при поступлении на направления аспирантура"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendPeriodMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Прием документов на места в рамках КЦП начинается с 20 июня и завершается 10 августа 23 года.\n" +
		"Прием документов на места в сверх КЦП начинается с 20 июня и завершается 12 сентября 2023 года.\n" +
		"Вступительные испытания в аспирантуру проводятся с 2 октября по 19 августа 2023 года."

	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}
