package moysklad

import (
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
	Version                      = "v0.0.73"                                // Версия библиотеки
	baseApiURL                   = "https://api.moysklad.ru/api/remap/1.2/" // Базовый адрес API
	ApplicationJson              = "application/json"                       // Тип данных
	headerWebHookDisable         = "X-Lognex-WebHook-Disable"               // Заголовок временного отключения уведомлений через API.
	headerGetContent             = "X-Lognex-Get-Content"                   // Заголовок для получения файла напрямую.
	headerWebHookDisableByPrefix = "X-Lognex-WebHook-DisableByPrefix"       // Заголовок временного отключения
	headerContentDisposition     = "Content-Disposition"                    // Заголовок содержит название файла при `X-Lognex-Get-Content: true`
	MaxPositions                 = 1000                                     // Максимальное число объектов, передаваемых в одном массиве в запросе
	MaxQueriesPerSecond          = 15                                       // Не более 45 запросов за 3 секундный период от аккаунта (45/3)
	MaxQueriesPerUser            = 5                                        // Не более 5 параллельных запросов от одного пользователя
	MaxPrintCount                = 1000                                     // Максимальное количество ценников/термоэтикеток
	headerRateLimit              = "X-RateLimit-Limit"                      // Количество запросов, которые равномерно можно сделать в течение интервала до появления 429 ошибки.
	headerRateRemaining          = "X-RateLimit-Remaining"                  // Число запросов, которые можно отправить до получения 429 ошибки.
	headerRetryTimeInterval      = "X-Lognex-Retry-TimeInterval"            // Интервал в миллисекундах, в течение которого можно сделать эти запросы

	//MaxFiles                = 100                           // Максимальное количество файлов
	//MaxImages               = 10                            // Максимальное количество изображений
	//headerRateReset         = "X-Lognex-Reset"              // Время до сброса ограничения в миллисекундах. Равно нулю, если ограничение не установлено.
	//headerRetryAfter        = "X-Lognex-Retry-After"        // Время до сброса ограничения в миллисекундах.
)

// Client базовый клиент для взаимодействия с API МойСклад.
type Client struct {
	*resty.Client
	limits      *queryLimits
	mu          sync.Mutex
	nextReqTime time.Time // время следующего запроса
}

func New(options ...func(*Client)) *Client {
	client := &Client{
		Client: resty.New(),
		limits: &queryLimits{
			// количество запросов за 3-х секундный период и количество параллельных запросов.
			rl:       ratelimit.New(MaxQueriesPerSecond),
			queryBuf: make(chan struct{}, MaxQueriesPerUser),
		},
	}

	for _, o := range options {
		o(client)
	}

	// устанавливаем базовый URL
	client.SetBaseURL(baseApiURL)

	// устанавливаем необходимые заголовки
	client.Header.Set("Accept", "application/json;charset=utf-8")
	client.Header.Set("Accept-Encoding", "gzip")
	client.Header.Set("User-Agent", fmt.Sprintf("go-moysklad/%s, https://github.com/arcsub/go-moysklad", Version))

	return client
}

// WithTokenAuth устанавливает авторизацию через Bearer токен.
func WithTokenAuth(token string) func(*Client) {
	return func(client *Client) {
		client.SetAuthToken(token)
	}
}

// WithBasicAuth устанавливает авторизацию по логину и паролю.
func WithBasicAuth(username, password string) func(*Client) {
	return func(client *Client) {
		client.SetBasicAuth(username, password)
	}
}

// WithHTTPClient устанавливает заранее инициализированный клиент [http.Client].
func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(client *Client) {
		if httpClient != nil {
			client.Client = resty.NewWithClient(httpClient)
		}
	}
}

// WithRestyClient устанавливает заранее инициализированный клиент [resty.Client].
func WithRestyClient(restyClient *resty.Client) func(*Client) {
	return func(client *Client) {
		if restyClient != nil {
			client.Client = restyClient
		}
	}
}

// WithDisabledWebhookContent устанавливает флаг, который отвечает
// за формирование заголовка временного отключения уведомления вебхуков через API (X-Lognex-WebHook-Disable).
//
// [Подробнее]
//
// [Подробнее]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-cherez-api
func WithDisabledWebhookContent(value bool) func(*Client) {
	return func(client *Client) {
		client.SetHeader(headerWebHookDisable, strconv.FormatBool(value))
	}
}

// WithDisabledWebhookByPrefix позволяет указать набор префиксов url-адресов.
//
// Если адрес вебхука содержит один из указанных префиксов, то этот вебхук не будет инициирован по результатам запроса.
//
// [Подробнее]
//
// [Подробнее]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-x-lognex-webhook-disablebyprefix-cherez-api
func WithDisabledWebhookByPrefix(urls ...string) func(*Client) {
	return func(client *Client) {
		for _, url := range urls {
			client.Header.Add(headerWebHookDisableByPrefix, url)
		}
	}
}

// WithTimeout устанавливает необходимый таймаут для http клиента.
func WithTimeout(timeout time.Duration) func(*Client) {
	return func(client *Client) {
		transport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: timeout,
			}).DialContext,
		}

		client.SetTimeout(timeout)
		client.SetTransport(transport)
	}
}

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
