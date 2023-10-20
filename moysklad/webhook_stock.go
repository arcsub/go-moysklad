package moysklad

import (
	"github.com/google/uuid"
)

// WebhookStock Вебхук на изменение остатков.
// Ключевое слово: webhookstock
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuk-na-izmenenie-ostatkow
type WebhookStock struct {
	AccountID         *uuid.UUID    `json:"accountId,omitempty"`         // ID учетной записи
	AuthorApplication *Meta         `json:"authorApplication,omitempty"` // Метаданные Приложения, создавшего вебхук на изменение остатков
	Enabled           *bool         `json:"enabled,omitempty"`           // Флажок состояние веб-хука (включен / отключен)
	StockType         *string       `json:"stockType,omitempty"`         // Тип остатков, изменение которых вызывает вебхук на изменение остатков. Возможные значения: [stock]
	ReportType        WebhookReport `json:"reportType,omitempty"`        // Тип отчета остатков, к которым привязан вебхук на изменение остатков. Возможные значения: [all, bystore]
	ReportUrl         *string       `json:"reportUrl,omitempty"`         // URL на получения данных по изменившейся номенклатуре за указанный период
	ID                *uuid.UUID    `json:"id,omitempty"`                // ID вебхука на изменение остатков
	Meta              *Meta         `json:"meta,omitempty"`              // Метаданные вебхука на изменение остатков
	URL               *string       `json:"url,omitempty"`               // URL, по которому будет происходить обработка вебхука. Допустимая длина до 255 символов
}

func (w WebhookStock) String() string {
	return Stringify(w)
}

func (w WebhookStock) MetaType() MetaType {
	return MetaTypeWebhookStock
}

// WebhookReport Тип отчета остатков, к которым привязан вебхук на изменение остатков
type WebhookReport string

const (
	WebhookReportAll     WebhookReport = "all"
	WebhookReportByStore WebhookReport = "bystore"
)
