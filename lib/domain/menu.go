package domain

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type Menu interface {
	ToReplyMarkup() telego.ReplyMarkup
}

type KeyboardMenuItem struct {
	Text string
}

func NewKeyboardMenuItem(text string) *KeyboardMenuItem {
	return &KeyboardMenuItem{
		Text: text,
	}
}

type KeyboardMenu struct {
	Items []*KeyboardMenuItem
}

func NewKeyboardMenu() *KeyboardMenu {
	return &KeyboardMenu{}
}

func (m *KeyboardMenu) Push(item *KeyboardMenuItem) *KeyboardMenu {
	m.Items = append(m.Items, item)
	return m
}

func (m *KeyboardMenu) ToReplyMarkup() telego.ReplyMarkup {
	rows := make([][]telego.KeyboardButton, 0, len(m.Items))
	for _, v := range m.Items {
		rows = append(rows, tu.KeyboardRow(tu.KeyboardButton(v.Text)))
	}
	return tu.Keyboard(rows...).
		WithSelective().
		WithResizeKeyboard().
		WithIsPersistent().
		WithInputFieldPlaceholder("Выберите пункт")
}
