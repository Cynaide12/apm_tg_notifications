package main

import (
	// internal "accept/internal"
	key "accept/internal/middleware"
	menu "accept/internal/menu"
	update "accept/pkg/update"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return menu.Menu(c,b, "Добро пожаловать! Чтобы вам начали приходить уведомления, необходимо пройти регистрацию.")
	})
	// Создаем роутер
	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)    // Логирование запросов
	r.Use(key.ApiKeyMiddleware) // Проверка API-ключа

	// Обработчик для POST-запросов
	r.Post("/send", func(w http.ResponseWriter, r *http.Request) {
		var msgReq update.MessageRequest

		// Декодируем JSON-запрос
		err := json.NewDecoder(r.Body).Decode(&msgReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Отправляем сообщение пользователю
		for _, telegramID := range msgReq.UserID {
			_, err = b.Send(tele.ChatID(telegramID), msgReq.Message)
			if err != nil {
				log.Printf("Failed to send message to %d: %v", telegramID, err)
				continue // Продолжаем отправку, даже если одному пользователю не удалось отправить
			}
		}

		// Отправляем успешный ответ
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message sent successfully"))
	})

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Println("Server started at :8081")
		log.Fatal(http.ListenAndServe("localhost:8081", r))
	}()

	log.Println("Бот запущен...")

	b.Start()
}
