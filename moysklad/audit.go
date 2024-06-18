package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"log"
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
	Diff           Diff           `json:"diff"`
	AdditionalInfo string         `json:"additionalInfo"`
	EventType      AuditEventType `json:"eventType"`
	EntityType     string         `json:"entityType"`
	Name           string         `json:"name"`
	ObjectType     string         `json:"objectType"`
	UID            string         `json:"uid"`
	Source         string         `json:"source"`
	ObjectCount    int            `json:"objectCount"`
	SupportAccess  bool           `json:"supportAccess"`
}

func (auditEvent AuditEvent) String() string {
	return Stringify(auditEvent)
}

func (auditEvent AuditEvent) MetaType() MetaType {
	return MetaTypeAuditEvent
}

// Diff тип поля diff
type Diff map[string]any

// Keys возвращает все ключи (поля) объекта Diff
func (diff Diff) Keys() []string {
	r := make([]string, 0, len(diff))
	for k := range diff {
		r = append(r, k)
	}
	return r
}

// IsExist возвращает true, если fieldName присутствует в объекте Diff
func (diff Diff) IsExist(fieldName string) bool {
	if _, ok := diff[fieldName]; ok {
		return true
	}
	return false
}

type OldNew[T any] struct {
	OldValue T `json:"oldValue"`
	NewValue T `json:"newValue"`
}

type SalePriceElem struct {
	OldValue struct {
		Uom   string
		Value float64
	} `json:"oldValue"`
	NewValue struct {
		Uom   string
		Value float64
	} `json:"newValue"`
}

type AuditPosition struct {
	Assortment struct {
		Meta Meta   `json:"meta"`
		Name string `json:"name"`
	} `json:"assortment"`
	Uom      string  `json:"uom"`
	Quantity float64 `json:"quantity"`
	Reserve  float64 `json:"reserve"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
}

func unmarshallAny[T any](data any) (T, error) {
	var t T
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return t, err
	}

	if err = json.Unmarshal(b, &t); err != nil {
		log.Println(err)
		return t, err
	}

	return t, nil
}

func getUnmarshall[T any](diff Diff, field string) (bool, T) {
	if positions, ok := diff[field]; ok {
		if p, err := unmarshallAny[T](positions); err == nil {
			return true, p
		}
	}
	return false, *new(T)
}

// GetPositions возвращает объект SalePriceElem
func (diff Diff) GetPositions() (bool, []OldNew[AuditPosition]) {
	return getUnmarshall[[]OldNew[AuditPosition]](diff, "positions")
}

// GetSalesPrices возвращает объект SalePriceElem
func (diff Diff) GetSalesPrices() (bool, SalePriceElem) {
	var o SalePriceElem
	if salePrices, ok := diff["salePrices"]; ok {
		var t OldNew[[]any]
		b, err := json.Marshal(salePrices)
		if err != nil {
			log.Println(err)
			return false, o
		}

		if err = json.Unmarshal(b, &t); err != nil {
			log.Println(err)
			return false, o
		}

		o.NewValue.Value = t.NewValue[0].(float64)
		o.NewValue.Uom = t.NewValue[1].(string)
		o.OldValue.Value = t.OldValue[0].(float64)
		o.OldValue.Uom = t.OldValue[1].(string)

		return true, o
	}
	return false, o
}

// GetFieldString возвращает объект OldNew со значениями типа string
func (diff Diff) GetFieldString(fieldName string) (bool, OldNew[string]) {
	return getUnmarshall[OldNew[string]](diff, fieldName)
}

// GetFieldBool возвращает объект OldNew со значениями типа bool
func (diff Diff) GetFieldBool(fieldName string) (bool, OldNew[bool]) {
	return getUnmarshall[OldNew[bool]](diff, fieldName)
}

// GetFieldFloat возвращает объект OldNew со значениями типа float64
func (diff Diff) GetFieldFloat(fieldName string) (bool, OldNew[float64]) {
	return getUnmarshall[OldNew[float64]](diff, fieldName)
}

// GetFieldInt возвращает объект OldNew со значениями типа int
func (diff Diff) GetFieldInt(fieldName string) (bool, OldNew[int]) {
	return getUnmarshall[OldNew[int]](diff, fieldName)
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
	GetContexts(ctx context.Context, params ...*Params) (*List[Audit], *resty.Response, error)

	// GetEvents выполняет запрос на получение Событий по Контексту AuditEvent.
	GetEvents(ctx context.Context, id uuid.UUID) (*List[AuditEvent], *resty.Response, error)

	// GetFilters выполняет запрос на получение Фильтров AuditFilters.
	GetFilters(ctx context.Context) (*AuditFilters, *resty.Response, error)
}

type auditService struct {
	Endpoint
}

func NewAuditService(client *Client) AuditService {
	return &auditService{NewEndpoint(client, "audit")}
}

func (service *auditService) GetContexts(ctx context.Context, params ...*Params) (*List[Audit], *resty.Response, error) {
	return NewRequestBuilder[List[Audit]](service.client, service.uri).SetParams(params...).Get(ctx)
}

func (service *auditService) GetEvents(ctx context.Context, id uuid.UUID) (*List[AuditEvent], *resty.Response, error) {
	path := fmt.Sprintf("audit/%s/events", id)
	return NewRequestBuilder[List[AuditEvent]](service.client, path).Get(ctx)
}

func (service *auditService) GetFilters(ctx context.Context) (*AuditFilters, *resty.Response, error) {
	path := "audit/metadata/filters"
	return NewRequestBuilder[AuditFilters](service.client, path).Get(ctx)
}
