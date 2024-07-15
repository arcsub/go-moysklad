package moysklad

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go.uber.org/ratelimit"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	Version                      = "v0.0.64"                                // Версия библиотеки
	baseApiURL                   = "https://api.moysklad.ru/api/remap/1.2/" // Базовый адрес API
	headerWebHookDisable         = "X-Lognex-WebHook-Disable"               // Заголовок временного отключения уведомлений через API.
	headerGetContent             = "X-Lognex-Get-Content"                   // Заголовок для получения файла напрямую.
	headerWebHookDisableByPrefix = "X-Lognex-WebHook-DisableByPrefix"       // Заголовок временного отключения
	headerContentDisposition     = "Content-Disposition"                    // Заголовок содержит название файла при `X-Lognex-Get-Content: true`
	MaxPositions                 = 1000                                     // Максимальное число объектов, передаваемых в одном массиве в запросе
	MaxQueriesPerSecond          = 15                                       // Не более 45 запросов за 3 секундный период от аккаунта (45/3)
	MaxQueriesPerUser            = 5                                        // Не более 5 параллельных запросов от одного пользователя
	MaxPrintCount                = 1000                                     // Максимальное количество ценников/термоэтикеток

	//MaxFiles                = 100                           // Максимальное количество файлов
	//MaxImages               = 10                            // Максимальное количество изображений
	//headerRateLimit         = "X-RateLimit-Limit"           // Количество запросов, которые равномерно можно сделать в течение интервала до появления 429 ошибки.
	//headerRateRemaining     = "X-RateLimit-Remaining"       // Число запросов, которые можно отправить до получения 429 ошибки.
	//headerRetryTimeInterval = "X-Lognex-Retry-TimeInterval" // Интервал в миллисекундах, в течение которого можно сделать эти запросы
	//headerRateReset         = "X-Lognex-Reset"              // Время до сброса ограничения в миллисекундах. Равно нулю, если ограничение не установлено.
	//headerRetryAfter        = "X-Lognex-Retry-After"        // Время до сброса ограничения в миллисекундах.

)

type queryLimits struct {
	rl       ratelimit.Limiter // Лимит между запросами
	queryBuf chan struct{}     // Буферизированный канал
}

func (queryLimits *queryLimits) Wait() {
	queryLimits.queryBuf <- struct{}{}
	queryLimits.rl.Take()
}

func (queryLimits *queryLimits) Done() {
	<-queryLimits.queryBuf
}

// Client базовый клиент для взаимодействия с API МойСклад.
type Client struct {
	*resty.Client
	limits   *queryLimits
	clientMu sync.Mutex
}

// NewClient возвращает новый клиент для работы с API МойСклад.
func NewClient() *Client {
	return (&Client{Client: resty.New()}).init()
}

// NewHTTPClient принимает [http.Client] и возвращает новый клиент для работы с API МойСклад.
func NewHTTPClient(httpClient *http.Client) *Client {
	return (&Client{Client: resty.NewWithClient(httpClient)}).init()
}

// NewRestyClient принимает [resty.Client] и возвращает новый клиент для работы с API МойСклад.
func NewRestyClient(restyClient *resty.Client) *Client {
	return (&Client{Client: restyClient}).init()
}

// init инициализирует параметры клиента.
func (client *Client) init() *Client {
	client.setQueryLimits().
		SetBaseURL(baseApiURL).
		SetHeaders(headers()).
		AddRetryCondition(retryCondition)
	return client
}

// headers устанавливает необходимые заголовки.
func headers() map[string]string {
	return map[string]string{
		"Accept":          "application/json;charset=utf-8", // https://dev.moysklad.ru/doc/api/remap/1.2/index.html#error_1062
		"Accept-Encoding": "gzip",                           // Обязательное использование сжатия содержимого ответов
		"User-Agent":      fmt.Sprintf("go-moysklad/%s (https://github.com/arcsub/go-moysklad)", Version),
	}
}

// retryCondition проверяет условия для повторного запроса: 429 ошибка, 500 и выше, или ошибка сети
func retryCondition(r *resty.Response, err error) bool {
	return r.StatusCode() == http.StatusTooManyRequests || r.StatusCode() >= http.StatusInternalServerError || isNetError(err)
}

// isNetError принимает ошибку и возвращает true, если ошибка является одной из перечисленных:
//   - [net.Error]
//   - [net.ErrClosed]
//   - [context.DeadlineExceeded]
func isNetError(err error) bool {
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return true
	}
	return errors.Is(err, net.ErrClosed) || errors.Is(err, context.DeadlineExceeded)
}

// setQueryLimits устанавливает количество запросов за 3-х секундный период и количество параллельных запросов.
func (client *Client) setQueryLimits() *Client {
	client.limits = &queryLimits{
		rl:       ratelimit.New(MaxQueriesPerSecond),
		queryBuf: make(chan struct{}, MaxQueriesPerUser),
	}
	return client
}

// WithTimeout устанавливает необходимый таймаут для http клиента.
func (client *Client) WithTimeout(timeout time.Duration) *Client {
	client.SetTimeout(timeout)
	return client
}

// WithTokenAuth возвращает клиент с авторизацией через Bearer токен.
func (client *Client) WithTokenAuth(token string) *Client {
	client.SetAuthToken(token)
	return client
}

// WithBasicAuth возвращает клиент с авторизацией по паре логин:пароль.
func (client *Client) WithBasicAuth(username, password string) *Client {
	client.SetBasicAuth(username, password)
	return client
}

// WithDisabledWebhookContent устанавливает флаг, который отвечает
// за формирование заголовка временного отключения уведомления вебхуков через API (X-Lognex-WebHook-Disable).
//
// [Подробнее]
//
// [Подробнее]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-cherez-api
func (client *Client) WithDisabledWebhookContent(value bool) *Client {
	client.SetHeader(headerWebHookDisable, strconv.FormatBool(value))
	return client
}

// WithDisabledWebhookByPrefix позволяет указать набор префиксов url-адресов.
//
// Если адрес вебхука содержит один из указанных префиксов, то этот вебхук не будет инициирован по результатам запроса.
//
// [Подробнее]
//
// [Подробнее]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-x-lognex-webhook-disablebyprefix-cherez-api
func (client *Client) WithDisabledWebhookByPrefix(urls ...string) *Client {
	for _, url := range urls {
		client.Header.Add(headerWebHookDisableByPrefix, url)
	}
	return client
}
