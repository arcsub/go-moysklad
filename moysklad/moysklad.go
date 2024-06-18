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
	Version                      = "v0.0.61"
	baseApiURL                   = "https://api.moysklad.ru/api/remap/1.2/"
	headerWebHookDisable         = "X-Lognex-WebHook-Disable"         // Заголовок временного отключения уведомлений через API.
	headerGetContent             = "X-Lognex-Get-Content"             // Заголовок для получения файла напрямую.
	headerWebHookDisableByPrefix = "X-Lognex-WebHook-DisableByPrefix" // Заголовок временного отключения
	headerContentDisposition     = "Content-Disposition"              // Заголовок содержит название файла при `X-Lognex-Get-Content: true`
	MaxPositions                 = 1000                               // Максимальное число объектов, передаваемых в одном массиве в запросе
	MaxQueriesPerSecond          = 15                                 // Не более 45 запросов за 3 секундный период от аккаунта (45/3)
	MaxQueriesPerUser            = 5                                  // Не более 5 параллельных запросов от одного пользователя

	//MaxFiles                = 100                           // Максимальное количество файлов
	//MaxImages               = 10                            // Максимальное количество изображений
	//headerRateLimit         = "X-RateLimit-Limit"           // Количество запросов, которые равномерно можно сделать в течение интервала до появления 429 ошибки.
	//headerRateRemaining     = "X-RateLimit-Remaining"       // Число запросов, которые можно отправить до получения 429 ошибки.
	//headerRetryTimeInterval = "X-Lognex-Retry-TimeInterval" // Интервал в миллисекундах, в течение которого можно сделать эти запросы
	//headerRateReset         = "X-Lognex-Reset"              // Время до сброса ограничения в миллисекундах. Равно нулю, если ограничение не установлено.
	//headerRetryAfter        = "X-Lognex-Retry-After"        // Время до сброса ограничения в миллисекундах.
	//MaxPrintCount           = 1000                          // Максимальное количество ценников/термоэтикеток

)

func getUserAgent() string {
	return fmt.Sprintf("go-moysklad/%s (https://github.com/arcsub/go-moysklad)", Version)
}

type queryLimits struct {
	rl       ratelimit.Limiter
	queryBuf chan struct{}
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

// Возвращает true, если ошибка является одной из перечисленных:
// net.Error, net.ErrClosed, context.DeadlineExceeded
func isNE(err error) bool {
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return true
	}
	return errors.Is(err, net.ErrClosed) || errors.Is(err, context.DeadlineExceeded)
}

func (client *Client) init() *Client {
	client.setQueryLimits().
		SetBaseURL(baseApiURL).
		SetHeaders(map[string]string{
			"Accept":          "application/json;charset=utf-8", // https://dev.moysklad.ru/doc/api/remap/1.2/index.html#error_1062
			"Accept-Encoding": "gzip",                           // Обязательное использование сжатия содержимого ответов
			"User-Agent":      getUserAgent(),
		}).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() == http.StatusTooManyRequests || // 429
					r.StatusCode() >= http.StatusInternalServerError || // 500+
					isNE(err)
			},
		)

	return client
}

// NewClient возвращает новый клиент для работы с API МойСклад.
// Данный клиент не имеет встроенных сервисов.
// Его необходимо передавать при создании каждого нового экземпляра сервиса.
func NewClient() *Client {
	return (&Client{Client: resty.New()}).init()
}

// NewHTTPClient принимает *http.Client и возвращает новый клиент для работы с API МойСклад.
// Данный клиент не имеет встроенных сервисов.
// Его необходимо передавать при создании каждого нового экземпляра сервиса.
func NewHTTPClient(httpClient *http.Client) *Client {
	return (&Client{Client: resty.NewWithClient(httpClient)}).init()
}

// NewRestyClient принимает *resty.Client и возвращает новый клиент для работы с API МойСклад.
// Данный клиент не имеет встроенных сервисов.
// Его необходимо передавать при создании каждого нового экземпляра сервиса.
func NewRestyClient(restyClient *resty.Client) *Client {
	return (&Client{Client: restyClient}).init()
}

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

// WithTokenAuth возвращает клиент с авторизацией через токен.
func (client *Client) WithTokenAuth(token string) *Client {
	client.SetAuthToken(token)
	return client
}

// WithBasicAuth возвращает клиент с базовой авторизацией логин/пароль.
func (client *Client) WithBasicAuth(username, password string) *Client {
	client.SetBasicAuth(username, password)
	return client
}

// WithDisabledWebhookContent устанавливает флаг, который отвечает
// за формирование заголовка временного отключения уведомления вебхуков через API (X-Lognex-WebHook-Disable).
// Подробнее: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-cherez-api
func (client *Client) WithDisabledWebhookContent(value bool) *Client {
	client.SetHeader(headerWebHookDisable, strconv.FormatBool(value))
	return client
}

// WithDisabledWebhookByPrefix позволяет указать набор префиксов url-адресов.
// Если адрес вебхука содержит один из указанных префиксов, то этот вебхук не будет инициирован по результатам запроса.
// Подробнее: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-x-lognex-webhook-disablebyprefix-cherez-api
func (client *Client) WithDisabledWebhookByPrefix(urls ...string) *Client {
	for _, url := range urls {
		client.Header.Add(headerWebHookDisableByPrefix, url)
	}
	return client
}
