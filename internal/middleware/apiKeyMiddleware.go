package middleware

import (
	"net/http"
	"os"
)

func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем API-ключ из заголовка
		apiKey := r.Header.Get("Authorization")

		// Проверяем, совпадает ли ключ
		if apiKey != os.Getenv("MY_API_KEY") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Если ключ верный, передаем запрос дальше
		next.ServeHTTP(w, r)
	})
}