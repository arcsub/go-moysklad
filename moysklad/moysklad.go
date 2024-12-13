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
	Version                      = "v0.0.75"                                // Версия библиотеки
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

// Config конфигурация клиента.
// Обязательно указывать либо токен (в приоритете), либо логин и пароль.
//
// # Пример:
//
//	client := moysklad.New(moysklad.Config{
//		Token:                  "MS_TOKEN_HERE",
//		DisabledWebhookContent: true,
//	})
type Config struct {
	// Устанавливает авторизацию через Bearer токен (в приоритете).
	Token string

	// Устанавливает авторизацию по логину и паролю.
	Username, Password string

	// Устанавливает заранее инициализированный клиент [resty.Client] (в приоритете).
	RestyClient *resty.Client

	// Устанавливает заранее инициализированный клиент [http.Client].
	HTTPClient *http.Client

	// Устанавливает флаг, который отвечает за формирование заголовка временного отключения уведомления вебхуков
	// через API (X-Lognex-WebHook-Disable).
	//
	// [Подробнее]
	//
	// [Подробнее]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-cherez-api
	DisabledWebhookContent bool

	// Набор префиксов url-адресов.
	//
	// Если адрес вебхука содержит один из указанных префиксов, то этот вебхук не будет инициирован по результатам запроса.
	//
	// [Подробнее]
	//
	// [Подробнее]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-x-lognex-webhook-disablebyprefix-cherez-api
	DisabledWebhookByPrefix []string
}

// apply применяет конфигурацию к клиенту.
func (config Config) apply(client *Client) {
	if config.HTTPClient != nil {
		client.Client = resty.NewWithClient(config.HTTPClient)
	}

	if config.RestyClient != nil {
		client.Client = config.RestyClient
	} else {
		client.Client = resty.New()
	}

	if config.Username != "" && config.Password != "" {
		client.SetBasicAuth(config.Username, config.Password)
	}

	if config.Token != "" {
		client.SetAuthToken(config.Token)
	}

	if config.DisabledWebhookContent {
		client.SetHeader(headerWebHookDisable, strconv.FormatBool(true))
	}

	if len(config.DisabledWebhookByPrefix) > 0 {
		for _, prefix := range config.DisabledWebhookByPrefix {
			client.Header.Add(headerWebHookDisableByPrefix, prefix)
		}
	}

	// устанавливаем базовый URL
	client.SetBaseURL(baseApiURL)

	// устанавливаем необходимые заголовки
	client.Header.Set("Accept", "application/json;charset=utf-8")
	client.Header.Set("Accept-Encoding", "gzip")
	client.Header.Set("User-Agent", fmt.Sprintf("go-moysklad/%s, https://github.com/arcsub/go-moysklad", Version))
}

// New создает новый экземпляр клиента.
//
// Принимает аргумент [Config], в котором необходимо указывать либо токен (в приоритете), либо логин и пароль.
//
// # Пример:
//
//	client := moysklad.New(moysklad.Config{
//		Token:                  "MS_TOKEN_HERE",
//		DisabledWebhookContent: true,
//	})
func New(config Config) *Client {
	client := &Client{
		limits: &queryLimits{
			// количество запросов за 3-х секундный период и количество параллельных запросов.
			rl:       ratelimit.New(MaxQueriesPerSecond),
			queryBuf: make(chan struct{}, MaxQueriesPerUser),
		},
	}

	config.apply(client)

	return client
}

// queryLimits используется для ограничения количества параллельных запросов
// и минимального интервала между ними.
type queryLimits struct {
	rl       ratelimit.Limiter // Лимитатор для контроля частоты запросов
	queryBuf chan struct{}     // Буфер для управления параллельными запросами
}

func (queryLimits *queryLimits) Wait() {
	queryLimits.queryBuf <- struct{}{}
	queryLimits.rl.Take()
}

func (queryLimits *queryLimits) Done() {
	<-queryLimits.queryBuf
}
