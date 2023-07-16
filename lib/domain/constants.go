package domain

import (
	"github.com/mymmrac/telego"
)

type MessageHandler func(bot *telego.Bot, update telego.Update) (*telego.Message, error)

const (
	Bachelor  string = "Бакалавриат"
	Masters   string = "Магистратура"
	PostGrad  string = "Аспирантура"
	Secondary string = "Среднее профессиональное образование"
)

const (
	BachelorAndMastersFile = "http://bot.priem-chspu.ru/bachelor.pdf"
	MinPointsFile          = "http://bot.priem-chspu.ru/min_points.pdf"
	PostGradFile           = "http://bot.priem-chspu.ru/asp_rules.pdf"
	PostGradIndividualFile = "http://bot.priem-chspu.ru/indiv_asp.pdf"
	SecondaryFile          = "http://bot.priem-chspu.ru/spo_rules.pdf"
)

const (
	ApplicationTypes = "Способы подачи документов"
	ApplicationDocs  = "Документы для поступления"
	Contacts         = "Контакты"
	EducationLevels  = "Подробно об уровнях образования и направлениях подготовки ➡"
	BachelorMenu     = "Бакалавриат"
	MastersMenu      = "Магистратура"
	PostGradMenu     = "Аспирантура"
	SecondaryMenu    = "Среднее профессиональное образование"
	EntranceRule     = "Правила приема"
	MinPoints        = "Минимальные баллы"
	List             = "Перечень направлений подготовки и количество мест для приема"
	IndividualList   = "Перечень индивидуальных достижений"
	Exception        = "Особые права и преимущества"
	Periods          = "Сроки подачи документов"
	BackMessage      = "Назад"
	BackToLevels     = "Вернуться назад ⬅"
	AskSpecialist    = "Задать вопрос специалисту"
)
