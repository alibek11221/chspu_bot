package specialist

import (
	"CHGPU_T_BOT/lib/domain"
)

func getMenu() domain.Menu {
	return domain.NewInlineMenu().
		Push(domain.NewInlineMenuItem("Джебир").WithURL("https://t.me/+79280010502")).
		Push(domain.NewInlineMenuItem("Алихан").WithURL("https://t.me/+79288870592"))
}
