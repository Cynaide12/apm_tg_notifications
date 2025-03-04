package internal

import (
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func HandleText(c tele.Context, b *tele.Bot) error {
	c.Send("Введите код")
	b.Handle(tele.OnText, func(c tele.Context) error {
		userID := c.Sender().ID
		code := c.Text()

		if len(code) != 6 {
			return c.Send("Код неверный, попробуйте еще раз.")
		}

		// Отправка запроса на сайт
		responseMessage, responseCode, err := SendCodeToServer(code, int(userID))
		if err != nil {
			return c.Send("Произошла ошибка при отправке кода. Попробуйте еще раз.")
		}

		switch responseCode {
		case 200:
			b.Handle(tele.OnText, func(c tele.Context) error {
				return nil
			})
            
			return c.Send("Аккаунт успешно привязан к системе APM. Теперь вы начнете получать уведомления!")
		case 401:
			c.Send("Неверный код. Попробуйте еще раз.")
		default:
			c.Send(fmt.Sprintf("Неизвестный ответ от сервера: %s", responseMessage))
		}
		return nil
	})
	return nil
}
