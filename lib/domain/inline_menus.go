package domain

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type InlineMenuItem struct {
	Label    string
	URL      string
	Callback string
}

func NewInlineMenuItem(label string) *InlineMenuItem {
	return &InlineMenuItem{
		Label: label,
	}
}

func (im *InlineMenuItem) WithURL(url string) *InlineMenuItem {
	im.URL = url
	return im
}
func (im *InlineMenuItem) WithCallback(callback string) *InlineMenuItem {
	im.Callback = callback
	return im
}

type InlineMenu struct {
	Items []*InlineMenuItem
}

func NewInlineMenu() *InlineMenu {
	return &InlineMenu{}
}

func (m *InlineMenu) Push(item *InlineMenuItem) *InlineMenu {
	m.Items = append(m.Items, item)
	return m
}

func (m *InlineMenu) ToReplyMarkup() telego.ReplyMarkup {
	rows := make([][]telego.InlineKeyboardButton, 0, len(m.Items))
	for _, v := range m.Items {
		button := tu.InlineKeyboardButton(v.Label)
		if len(v.URL) > 0 {
			button = button.WithURL(v.URL)
		}
		if len(v.Callback) > 0 {
			button = button.WithCallbackData(v.Callback)
		}
		rows = append(rows, tu.InlineKeyboardRow(button))
	}
	return tu.InlineKeyboard().WithInlineKeyboard(rows...)
}
