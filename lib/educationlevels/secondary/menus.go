package secondary

import "CHGPU_T_BOT/lib/domain"

func getMenu() *domain.KeyboardMenu {
	return domain.NewKeyboardMenu().
		Push(domain.NewKeyboardMenuItem("Назад")).
		Push(domain.NewKeyboardMenuItem(domain.AskSpecialist))
}
