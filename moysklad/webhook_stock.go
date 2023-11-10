package moysklad

import (
	"github.com/google/uuid"
)

// WebhookStock Вебхук на изменение остатков.
// Ключевое слово: webhookstock
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuk-na-izmenenie-ostatkow
type WebhookStock struct {
	AccountID         *uuid.UUID    `json:"accountId,omitempty"`
	AuthorApplication *Meta         `json:"authorApplication,omitempty"`
	Enabled           *bool         `json:"enabled,omitempty"`
	StockType         *string       `json:"stockType,omitempty"`
	ReportUrl         *string       `json:"reportUrl,omitempty"`
	ID                *uuid.UUID    `json:"id,omitempty"`
	Meta              *Meta         `json:"meta,omitempty"`
	URL               *string       `json:"url,omitempty"`
	ReportType        WebhookReport `json:"reportType,omitempty"`
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
