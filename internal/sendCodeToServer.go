package internal

import (
	// request "accept/pkg/requests"
	update "accept/pkg/update"
	// "bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	// "os"
)



// func SendCodeToServer(code string, userID int) (string, int, error) {
//     url := fmt.Sprintf("http://kstapm-02.kst-energo.ru/api/v1/users/check_tg_code?code=%s&tgid=%v",code , userID) // Замените на ваш URL

//     // Создаем структуру для запроса
//     requestData := request.PostJsonMessage{
//         Code:   code,
//         UserID: userID,
//     }

//     // Кодируем структуру в JSON
//     jsonData, err := json.Marshal(requestData)
//     if err != nil {
//         return "", 0, err
//     }

//     // Создаем новый HTTP-запрос
//     req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
//     if err != nil {
//         return "", 0, err
//     }

//     req.Header.Set("Content-Type", "application/json") 
//     req.Header.Set("Authorization", "Bearer "+ os.Getenv("API_KEY_SITE"))

//     client := &http.Client{}
//     resp, err := client.Do(req)
//     if err != nil {
//         return "", 0, err
//     }
//     defer resp.Body.Close()

//     body, err := io.ReadAll(resp.Body)
//     if err != nil {
//         return "", 0, err
//     }

//     var registrationResponse *update.JsonAnswer
//     if err := json.Unmarshal(body, &registrationResponse); err != nil {
//         return "", resp.StatusCode, err
//     }

//     log.Printf("Response: %s", body)

//     return registrationResponse.Message, resp.StatusCode, nil
// }
func SendCodeToServer(code string, userID int) (string, int, error) {
    // Формируем URL с параметрами
    url := fmt.Sprintf("http://kstapm-02.kst-energo.ru/api/v1/users/check_tg_code?code=%s&tgid=%v", code, userID)

    // Создаем новый HTTP-запрос
    req, err := http.NewRequest("GET", url, nil) // Тело запроса не нужно для GET
    if err != nil {
        return "", 0, err
    }

    // Устанавливаем заголовки
    // req.Header.Set("Authorization", "Bearer "+os.Getenv("API_KEY_SITE"))

    // Выполняем запрос
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", 0, err
    }
    defer resp.Body.Close()

    // Читаем ответ
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", 0, err
    }

    // Парсим ответ
    var registrationResponse *update.JsonAnswer
    if err := json.Unmarshal(body, &registrationResponse); err != nil {
        return "", resp.StatusCode, err
    }

    log.Printf("Response: %s", body)

    return registrationResponse.Message, resp.StatusCode, nil
}
