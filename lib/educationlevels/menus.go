package educationlevels

import (
	"CHGPU_T_BOT/lib/domain"
)

func getMenu() *domain.KeyboardMenu {
	return domain.NewKeyboardMenu().
		Push(domain.NewKeyboardMenuItem(domain.EntranceRule)).
		Push(domain.NewKeyboardMenuItem(domain.MinPoints)).
		Push(domain.NewKeyboardMenuItem(domain.List)).
		Push(domain.NewKeyboardMenuItem(domain.IndividualList)).
		Push(domain.NewKeyboardMenuItem(domain.Exception)).
		Push(domain.NewKeyboardMenuItem(domain.Periods)).
		Push(domain.NewKeyboardMenuItem(domain.BackToLevels)).
		Push(domain.NewKeyboardMenuItem(domain.AskSpecialist))
}
