package mainmenu

import (
	"CHGPU_T_BOT/lib/state"
	"CHGPU_T_BOT/lib/utils"

	"github.com/mymmrac/telego"
)

func SendInitMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Что умеет этот бот?</b>" +
		"\nПривет, будущий студент Чеченского государственного педагогического университета!" +
		"\nЭтот бот поможет тебе разобраться в особенностях поступления в наш Университет." +
		"\nЕсли останутся вопросы, то обратись к специалистам нашей Приемной комиссии."
	menu := getInitMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}
func SendWelcomeBackMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "Начнем сначала 😀"
	menu := getInitMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}
func SendApplicationTypesMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Способы подачи документов:</b>"
	menu := getApplicationTypesMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}

func SendApplicationDocsMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Документы для поступления:</b>\n" +
		"<b>КОЛЛЕДЖ:</b>\n" +
		"1️⃣ Копия паспорта с оригиналом\n" +
		"2️⃣ Копия паспорта одного из родителей или законного представителя несовершеннолетнего\n" +
		"3️⃣ Копия аттестата с оригиналом\n" +
		"4️⃣ Копия СНИЛС с оригиналом\n" +
		"5️⃣ 3х4 6 фотографий\n" +
		"6️⃣ Адрес электронной почты\n" +
		"7️⃣ Медсправка формы 086-у\n" +
		"<b>БАКАЛАВРИАТ:</b>\n" +
		"1️⃣ Копия паспорта с оригиналом\n" +
		"2️⃣ Копия паспорта одного из родителей или законного представителя несовершеннолетнего\n" +
		"3️⃣ Копия аттестата с оригиналом\n" +
		"4️⃣ Копия СНИЛС с оригиналом\n" +
		"5️⃣ 3х4 6 фотографий\n" +
		"6️⃣ Адрес электронной почты\n" +
		"7️⃣ Медсправка формы 086-у\n" +
		"<b>МАГИСТРАТУРА:</b>\n" +
		"1️⃣ Копия паспорта с оригиналом \n" +
		"2️⃣ Копия диплома бакалавра или специалиста с оригиналом \n" +
		"3️⃣ Копия СНИЛС с оригиналом \n" +
		"4️⃣ 3х4 6 фотографий \n" +
		"5️⃣ Адрес электронной почты \n" +
		"<b>АСПИРАНТУРА:</b>\n" +
		"1️⃣ Копия паспорта с оригиналом \n" +
		"2️⃣ Копия диплома специалиста или магистра с оригиналом \n" +
		"3️⃣ Копия СНИЛС с оригиналом \n" +
		"4️⃣ 3х4 6 фотографий \n" +
		"5️⃣ Адрес электронной почты"

	menu := getApplicationDocsMenu().ToReplyMarkup()

	return utils.SendTextMessage(bot, update, text, menu)
}

func SendEducationLevelsMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Выберите уровень образования</b>"
	state.RemoveFromState(update.Message.Chat.ID)
	menu := getEducationLevelsMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}

func SendContactsMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Контакты:</b> \n\n" +
		"📞<a href='tel:8(928)001-05-02'>8 (928) 001-05-02</a> \n" +
		"📞<a href='tel:8(928)887-05-92'>8 (928) 887-05-92</a> \n" +
		"📧<a href='mailto:pk@chspu.ru'>pk@chspu.ru</a> \n" +
		"📬<a href='https://yandex.ru/maps/-/CCUDvZFrpD'>г. Грозный, ул. Субры Кишиевой, 33</a> \n" +
		"🕧Понедельник-пятница с 09:00 до 18:00 (перерыв с 12:30 до 13:30) \n" +
		"🕧Суббота с 09:00 до 14:00 (без перерыва) \n" +
		"Воскресенье (выходной день)"
	menu := getContactsMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}

func SendEGECalculatorMessage(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	text := "<b>Калькулятор ЕГЭ:</b> \n\n" +
		"Подберите себе направление для поступления по предметам ЕГЭ.\n" +
		"Калькулятор доступен <a href='https://chspu.ru/priemnaja-kampanija/'>по ссылке </a>"
	menu := getEGECalculatorMenu().ToReplyMarkup()
	return utils.SendTextMessage(bot, update, text, menu)
}

func SendErrorMsg(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	msg := "Возможно, данное сообщение не соответствует выбранному вами разделу или оно не содержит нужной информации." +
		"Выберите пункт из меню и попробуйте снова"
	menu := getInitMenu().ToReplyMarkup()
	state.RemoveFromState(update.Message.Chat.ID)
	return utils.SendTextMessage(bot, update, msg, menu)
}
