package moysklad

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"go.uber.org/ratelimit"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	Version                  = "v0.0.46"
	baseApiURL               = "https://api.moysklad.ru/api/remap/1.2/"
	headerWebHookDisable     = "X-Lognex-WebHook-Disable" // Заголовок временного отключения уведомлений через API.
	headerGetContent         = "X-Lognex-Get-Content"     // Заголовок для получения файла напрямую.
	headerContentDisposition = "Content-Disposition"      // Заголовок содержит название файла при `X-Lognex-Get-Content: true`
	MaxFiles                 = 100                        // Максимальное количество файлов
	MaxImages                = 10                         // Максимальное количество изображений
	MaxPositions             = 1000                       // Максимальное число объектов, передаваемых в одном массиве в запросе
	MaxQueriesPerSecond      = 15                         // Не более 45 запросов за 3 секундный период от аккаунта (45/3)
	MaxQueriesPerUser        = 5                          // Не более 5 параллельных запросов от одного пользователя

	//headerRateLimit         = "X-RateLimit-Limit"           // Количество запросов, которые равномерно можно сделать в течение интервала до появления 429 ошибки.
	//headerRateRemaining     = "X-RateLimit-Remaining"       // Число запросов, которые можно отправить до получения 429 ошибки.
	//headerRetryTimeInterval = "X-Lognex-Retry-TimeInterval" // Интервал в миллисекундах, в течение которого можно сделать эти запросы
	//headerRateReset         = "X-Lognex-Reset"              // Время до сброса ограничения в миллисекундах. Равно нулю, если ограничение не установлено.
	//headerRetryAfter        = "X-Lognex-Retry-After"        // Время до сброса ограничения в миллисекундах.
	//MaxPrintCount           = 1000                          // TODO: Максимальное количество ценников/термоэтикеток

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
				return r.StatusCode() == http.StatusTooManyRequests ||
					r.StatusCode() >= http.StatusBadGateway
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
	clone := client.copy()
	clone.SetAuthToken(token)
	return clone
}

// WithBasicAuth возвращает клиент с базовой авторизацией логин/пароль.
func (client *Client) WithBasicAuth(username, password string) *Client {
	clone := client.copy()
	clone.SetBasicAuth(username, password)
	return clone
}

// WithDisabledWebhookContent устанавливает флаг, который отвечает
// за формирование заголовка временного отключения уведомления вебхуков через API (X-Lognex-WebHook-Disable).
// Подробнее: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-cherez-api
func (client *Client) WithDisabledWebhookContent(value bool) *Client {
	clone := client.copy()
	client.copy().SetHeader(headerWebHookDisable, strconv.FormatBool(value))
	return clone
}

// copy возвращает копию клиента.
func (client *Client) copy() *Client {
	client.clientMu.Lock()
	defer client.clientMu.Unlock()
	clone := &Client{
		Client: client.Client,
		limits: client.limits,
	}
	return clone
}
