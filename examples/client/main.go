package main

import (
	ms "github.com/arcsub/go-moysklad/moysklad"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

const (
	MSUsername = "MS_USERNAME" // Имя пользователя
	MSPassword = "MS_USERNAME" // Пароль пользователя
	MSToken    = "MS_TOKEN"    // Токен API
)

func main() {

	// Получить стандартный клиент для работы с API МойСклад
	client := ms.NewClient()

	// Получить клиент cо своим http.Client
	httpClient := http.DefaultClient
	_ = ms.NewHTTPClient(httpClient)

	// Получить клиент cо своим resty.Client
	restyClient := resty.New()
	_ = ms.NewRestyClient(restyClient)

	// Клиент с авторизацией по паре логин:пароль
	client.WithBasicAuth(MSUsername, MSPassword)

	// Клиент с авторизацией по Bearer токену
	client.WithTokenAuth(MSToken)

	// Установить таймаут для запросов
	client.WithTimeout(5 * time.Minute)

	// Отключить уведомление вебхуков на клиенте
	client.WithDisabledWebhookContent(true)

	// Выборочно отключить часть уведомлений вебхуков
	// Подробнее: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-x-lognex-webhook-disablebyprefix-cherez-api
	{
		disabledURLs := []string{
			"https://abc.ru/wh1",
			"https://abc.ru/wh2",
		}

		client.WithDisabledWebhookByPrefix(disabledURLs...)
	}
}
