package masters

import (
	"CHGPU_T_BOT/lib/domain"
	"CHGPU_T_BOT/lib/utils"

	"github.com/mymmrac/telego"
)

func SendApplicationRulesMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Каждый поступающий обязан внимательно ознакомиться с \n " +
		"Правилами приема в Университете</b>"
	menu := getMenu().ToReplyMarkup()
	return utils.SendDocument(bot, update, text, domain.BachelorAndMastersFile, menu)
}
func SendMinimalPointsMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Ознакомьтесь с минимальными баллами ЕГЭ и внутренних вступительных испытаний университета, " +
		"а также со сроками обучения в вузе"
	menu := getMenu().ToReplyMarkup()

	return utils.SendDocument(bot, update, text, domain.MinPointsFile, menu)
}

func SendListMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Информация доступна по <a href='https://chspu.ru/abitur/bachelor/#abitur_plan_priema'>ссылке</a>"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendIndividualListMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Максимальное количество баллов за индивидуальные достижения—10 баллов \n" +
		"Полный перечень учета достижений смотрите по <a href='https://chspu.ru/abitur/bachelor/#abitur_priemIndivid'>ссылке</a>"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendExceptionMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Особые права и преимущества отсутствуют при поступлении на направления магистратуры"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendPeriodMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Прием документов начинается с 20 июня 2023 г. \n" +
		"Завершается прием документов: \n" +
		"10 августа 2023 г. завершается прием документов на очную и заочную формы обучения на бюджетные места. \n" +
		"20 августа 2023 г. завершается прием документов на очную, очно-заочную и заочную формы обучения на платные места."

	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}
