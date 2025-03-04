package menu

import (
	tele "gopkg.in/telebot.v4"
	internal "accept/internal"
)
func Menu(c tele.Context, b *tele.Bot,text string) error {
	menu := &tele.ReplyMarkup{}
	btnRegistration := menu.Data("Ввести код", "input_code")

	menu.Inline(
		menu.Row(btnRegistration),
	)
	b.Handle(&btnRegistration, func(c tele.Context) error {
		return internal.HandleText(c, b)
	})

	err := c.Send(text, menu)
    if err != nil {
        return err
    }

    return nil
}
