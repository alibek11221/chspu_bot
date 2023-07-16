package mainmenu

import (
	"CHGPU_T_BOT/lib/domain"
)

func getInitMenu() *domain.KeyboardMenu {
	return domain.NewKeyboardMenu().
		Push(domain.NewKeyboardMenuItem(domain.ApplicationTypes)).
		Push(domain.NewKeyboardMenuItem(domain.ApplicationDocs)).
		Push(domain.NewKeyboardMenuItem(domain.Contacts)).
		Push(domain.NewKeyboardMenuItem(domain.EducationLevels))
}

func getApplicationTypesMenu() *domain.InlineMenu {
	return domain.NewInlineMenu().
		Push(domain.NewInlineMenuItem("🚩 Личная подача г.Грозный, ул.Субры Кишиевой, д.33, ЧГПУ").WithURL("https://yandex.ru/maps/-/CCUDvZFrpD")).
		Push(domain.NewInlineMenuItem("🚩 Посредством информационной системы ЧГПУ").WithURL("http://188.0.166.207/login")).
		Push(domain.NewInlineMenuItem("🚩 Посредством единого портала государственных услуг").WithURL("https://www.gosuslugi.ru/vuzonline")).
		Push(domain.NewInlineMenuItem("🚩 Почтой россии На адрес: 364037, г.Грозный, ул.Субры Кишиевой, д.33, в приемную комиссию ЧГПУ").WithURL("https://yandex.ru/maps/-/CCUDvZFrpD"))
}

func getApplicationDocsMenu() *domain.InlineMenu {
	return domain.NewInlineMenu()
}

func getContactsMenu() *domain.InlineMenu {
	return domain.NewInlineMenu()
}

func getEGECalculatorMenu() *domain.InlineMenu {
	return domain.NewInlineMenu()
}

func getBackInlineButton() *domain.InlineMenuItem {
	return domain.NewInlineMenuItem("Назад").WithCallback("back")
}

func getEducationLevelsMenu() *domain.KeyboardMenu {
	return domain.NewKeyboardMenu().
		Push(domain.NewKeyboardMenuItem(domain.BachelorMenu)).
		Push(domain.NewKeyboardMenuItem(domain.MastersMenu)).
		Push(domain.NewKeyboardMenuItem(domain.PostGradMenu)).
		Push(domain.NewKeyboardMenuItem(domain.SecondaryMenu)).
		Push(domain.NewKeyboardMenuItem(domain.BackToLevels)).
		Push(domain.NewKeyboardMenuItem(domain.AskSpecialist))
}
