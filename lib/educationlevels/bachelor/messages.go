package bachelor

import (
	"CHGPU_T_BOT/lib/domain"
	"CHGPU_T_BOT/lib/utils"

	"github.com/mymmrac/telego"
)

func SendApplicationRulesMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Каждый поступающий обязан внимательно ознакомиться с\n" +
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
	text := "Максимальное количество баллов за индивидуальные достижения—10 баллов\n" +
		"Полный перечень учета достижений смотрите по <a href='https://chspu.ru/abitur/bachelor/#abitur_priemIndivid'>ссылке</a>"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendExceptionMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Право на прием на обучение за счет бюджетных ассигнований в пределах особой квоты имеют:\n" +
		"дети-инвалиды;\n" +
		"инвалиды I и II групп;\n" +
		"инвалиды с детства;\n" +
		"инвалиды вследствие военной травмы или заболевания, полученных в период прохождения военной службы;\n" +
		"дети-сироты и дети, оставшиеся без попечения родителей;\n" +
		"Лица из числа детей-сирот и детей, оставшихся без попечения родителей.\n\n" +
		"Подробная информация по <a href='https://chspu.ru/abitur/bachelor/#abitur_priemOsob'>ссылке</a>"
	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendPeriodMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Прием документов начинается с 20 июня 2023 г. \n" +
		"Завершается прием документов по очной, очно-заочной и заочной формам обучения на бюджетные места: \n" +
		"18 июля 2023 года от лиц, поступающих на обучение по результатам вступительных испытаний, проводимых университетом самостоятельно; \n" +
		"18 июля 2023 года от лиц, поступающих на обучение по результатам дополнительных вступительных испытаний творческой и (или) профессиональной направленности; \n" +
		"25 июля 2023 года завершается прием документов от лиц, поступающих на основании ЕГЭ. \n" +
		"Подробная информация по <a href='https://chspu.ru/abitur/bachelor/#abitur_podat-dokumenty'>ссылке</a> "

	menu := getMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}
