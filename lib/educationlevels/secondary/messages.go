package secondary

import (
	"CHGPU_T_BOT/lib/domain"
	"CHGPU_T_BOT/lib/utils"
	"github.com/mymmrac/telego"
)

func SendApplicationRulesMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Каждый поступающий обязан внимательно ознакомиться с \n " +
		"Правилами приема в Университете</b>"
	menu := getMenu().ToReplyMarkup()
	return utils.SendDocument(bot, update, text, domain.SecondaryFile, menu)
}
func SendMinimalPointsMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "При приеме на обучение по программам среднего профессионального образования вступительные испытания не проводятся.\n" +
		"Прием осуществляется на основании конкурса аттестатов о среднем общем образовании (11 класс) или аттестатов об основном общем образовании (9 класс)"
	menu := getMenu().ToReplyMarkup()

	return utils.SendDocument(bot, update, text, domain.MinPointsFile, menu)
}

func SendListMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Информация доступна по <a href='https://chspu.ru/abitur/spo_sredn/#abitur_plan_priema'>ссылке</a>"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendIndividualListMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Индивидуальные достижения на специальности СПО не учитываются"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendExceptionMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Особые права и преимущества отсутствуют при поступлении на программы СПО"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendPeriodMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Прием документов начинается с 20 июня 2023 г. \n" +
		"Завершается прием документов: \n" +
		"10 августа 2023 года на очную форму в рамках КЦП (на базе 9 классов). \n" +
		"20 августа 2023 года на заочную форму в рамках КЦП (на базе 9 классов). \n" +
		"10 сентября 2023 года на очную, очно-заочную и заочную формам обучения на места по договорам об оказании " +
		"платных образовательных услуг (на базе 9 и 11 классов)."

	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}
