package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Audit Контексты Аудита.
// Ключевое слово: audit
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-kontexty
type Audit struct {
	Events        MetaWrapper      `json:"events"`
	Meta          Meta             `json:"meta"`
	Moment        Timestamp        `json:"moment"`
	EntityType    string           `json:"entityType"`
	EventType     AuditEventType   `json:"eventType"`
	Info          AuditContextInfo `json:"info"`
	ObjectType    string           `json:"objectType"`
	Source        string           `json:"source"`
	UID           string           `json:"uid"`
	ObjectCount   int              `json:"objectCount"`
	ID            uuid.UUID        `json:"id"`
	SupportAccess bool             `json:"supportAccess"`
}

func (audit Audit) String() string {
	return Stringify(audit)
}

func (audit Audit) MetaType() MetaType {
	return MetaTypeAudit
}

type AuditContextInfo struct {
	AdditionalInfo string `json:"additionalInfo"`
}

// AuditEvent Событие аудита.
// Ключевое слово: auditevent
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq
type AuditEvent struct {
	Entity         MetaWrapper    `json:"entity"`
	Audit          MetaWrapper    `json:"audit"`
	Moment         Timestamp      `json:"moment"`
	Diff           any            `json:"diff"`
	EntityType     string         `json:"entityType"`
	EventType      AuditEventType `json:"eventType"`
	AdditionalInfo string         `json:"additionalInfo"`
	Name           string         `json:"name"`
	ObjectType     string         `json:"objectType"`
	Source         string         `json:"source"`
	UID            string         `json:"uid"`
	ObjectCount    int            `json:"objectCount"`
	SupportAccess  bool           `json:"supportAccess"`
}

func (auditEvent AuditEvent) String() string {
	return Stringify(auditEvent)
}

func (auditEvent AuditEvent) MetaType() MetaType {
	return MetaTypeAuditEvent
}

// TODO: поле diff в событии аудита может иметь разный тип данных: {},[] и ""

// AuditEventDiff Формат поля diff
//
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq-format-polq-diff
type AuditEventDiff struct {
	OldValue        any    `json:"oldValue,omitempty"`
	NewValue        any    `json:"newValue,omitempty"`
	Account         string `json:"account,omitempty"`
	Country         string `json:"country,omitempty"`
	TemplateName    string `json:"templateName,omitempty"`
	PublicationHref string `json:"publicationHref,omitempty"`
	SenderEmail     string `json:"senderEmail,omitempty"`
	TargetEmail     string `json:"targetEmail,omitempty"`
	SubjectEmail    string `json:"subjectEmail,omitempty"`
	Text            string `json:"text,omitempty"`
	AttributeName   string `json:"attributeName,omitempty"`
}

func (auditEventDiff AuditEventDiff) String() string {
	return Stringify(auditEventDiff)
}

// AuditEventType Действие События
type AuditEventType string

const (
	AuditEventTypeRegistration          AuditEventType = "registration"          // Регистрация
	AuditEventTypeBulkOperation         AuditEventType = "bulkoperation"         // Массовая операция
	AuditEventTypeClosePublication      AuditEventType = "closepublication"      // Удаление публикации
	AuditEventTypeCreate                AuditEventType = "create"                // Создание сущностей
	AuditEventTypeDelete                AuditEventType = "delete"                // Удаление сущностей
	AuditEventTypeOpenPublication       AuditEventType = "openpublication"       // Создание публикации
	AuditEventTypePrint                 AuditEventType = "print"                 // Печать документа
	AuditEventTypePutToArchive          AuditEventType = "puttoarchive"          // Помещение в архив
	AuditEventTypePutToRecycleBin       AuditEventType = "puttorecyclebin"       // Помещение в корзину
	AuditEventTypeReplaceToken          AuditEventType = "replacetoken"          // Смена токена для Точки продаж
	AuditEventTypeRestoreFromArchive    AuditEventType = "restorefromarchive"    // Извлечение из архива
	AuditEventTypeRestoreFromRecycleBin AuditEventType = "restorefromrecyclebin" // Извлечение из корзины
	AuditEventTypeSendEmailFromEntity   AuditEventType = "sendemailfromentity"   // Отправка письма
	AuditEventTypeUpdate                AuditEventType = "update"                // Изменение сущностей
)

// AuditFilters Фильтры
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-fil-try
type AuditFilters struct {
	EventType  Slice[string] `json:"eventType"`  // Действия, по которым могут быть отфильтрованы сущности аудита
	Source     Slice[string] `json:"source"`     // Типы действий, по которым могут быть отфильтрованы сущности аудита
	EntityType Slice[string] `json:"entityType"` // Названия сущностей, по которым могут быть отфильтрованы сущности аудита
}

func (auditFilters AuditFilters) String() string {
	return Stringify(auditFilters)
}

// AuditService
// Сервис для работы с аудитом.
type AuditService interface {
	// GetContexts выполняет запрос на получение Контекстов Audit.
	//
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-kontexty
	GetContexts(ctx context.Context, params *Params) (*List[Audit], *resty.Response, error)

	// GetEvents выполняет запрос на получение Событий по Контексту AuditEvent.
	//
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-sobytiq-po-kontextu
	GetEvents(ctx context.Context, id uuid.UUID) (*List[AuditEvent], *resty.Response, error)

	// GetFilters выполняет запрос на получение Фильтров AuditFilters.
	//
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-poluchit-fil-try
	GetFilters(ctx context.Context) (*AuditFilters, *resty.Response, error)
}

type auditService struct {
	Endpoint
}

func NewAuditService(client *Client) AuditService {
	return &auditService{NewEndpoint(client, "audit")}
}

func (service *auditService) GetContexts(ctx context.Context, params *Params) (*List[Audit], *resty.Response, error) {
	return NewRequestBuilder[List[Audit]](service.client, service.uri).SetParams(params).Get(ctx)
}

func (service *auditService) GetEvents(ctx context.Context, id uuid.UUID) (*List[AuditEvent], *resty.Response, error) {
	path := fmt.Sprintf("audit/%s/events", id)
	return NewRequestBuilder[List[AuditEvent]](service.client, path).Get(ctx)
}

func (service *auditService) GetFilters(ctx context.Context) (*AuditFilters, *resty.Response, error) {
	path := "audit/metadata/filters"
	return NewRequestBuilder[AuditFilters](service.client, path).Get(ctx)
}
