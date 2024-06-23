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
//
// Код сущности: audit
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-kontexty
type Audit struct {
	Events        MetaWrapper    `json:"events"`        // Список метаданных Событий аудита
	Meta          Meta           `json:"meta"`          // Метаданные сущности Контекста
	Moment        Timestamp      `json:"moment"`        // Дата изменения
	EntityType    MetaType       `json:"entityType"`    // Название сущности (поле присутствует, только если оно одинаково у всех Событий в рамках данного Контекста)
	EventType     AuditEventType `json:"eventType"`     // Действие Событий (поле присутствует, только если оно одинаково у всех Событий в рамках данного Контекста)
	Info          AuditInfo      `json:"info"`          // Краткое описание
	ObjectType    MetaType       `json:"objectType"`    // Тип сущностей, с которыми связанно данное изменение. Поле присутствует только для entityType = entitysettings или statesettings или templatesettings
	Source        string         `json:"source"`        // Тип изменения
	UID           string         `json:"uid"`           // Логин Сотрудника
	ObjectCount   int            `json:"objectCount"`   // количество измененных объектов
	ID            uuid.UUID      `json:"id"`            // ID Контекста
	SupportAccess bool           `json:"supportAccess"` // Был ли доступ произведен поддержкой от имени пользователя. Флаг отсутствует, если значение false
}

// String реализует интерфейс [fmt.Stringer].
func (audit Audit) String() string {
	return Stringify(audit)
}

// MetaType возвращает тип сущности.
func (Audit) MetaType() MetaType {
	return MetaTypeAudit
}

// AuditInfo Краткое описание.
type AuditInfo struct {
	AdditionalInfo string `json:"additionalInfo"` // Содержание краткого описания
}

// AuditEvent Событие аудита.
//
// Код сущности: auditevent
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq
type AuditEvent struct {
	Entity         MetaWrapper    `json:"entity"`         // Метаданные сущности. Не будет выводиться только для товаров, услуг, модификаций, комплектов удаленных до 20.08.2017
	Audit          MetaWrapper    `json:"audit"`          // Метаданные контекста
	Moment         Timestamp      `json:"moment"`         // Время создания события
	Diff           Diff           `json:"diff"`           // Изменения, произошедшие в Событии
	AdditionalInfo string         `json:"additionalInfo"` // Дополнительная информация о Событии
	EventType      AuditEventType `json:"eventType"`      // Действие События
	EntityType     MetaType       `json:"entityType"`     // Название сущности
	Name           string         `json:"name"`           // Имя сущности
	ObjectType     MetaType       `json:"objectType"`     // Тип сущностей, с которыми связанно данное изменение. Поле присутствует только для entityType = entitysettings или statesettings или templatesettings
	UID            string         `json:"uid"`            // Логин Сотрудника
	Source         string         `json:"source"`         // Тип изменения
	ObjectCount    int            `json:"objectCount"`    // количество измененных объектов
	SupportAccess  bool           `json:"supportAccess"`  // Был ли доступ произведен поддержкой от имени пользователя. Флаг отсутствует, если значение false
}

// String реализует интерфейс [fmt.Stringer].
func (auditEvent AuditEvent) String() string {
	return Stringify(auditEvent)
}

// MetaType возвращает тип сущности.
func (AuditEvent) MetaType() MetaType {
	return MetaTypeAuditEvent
}

// Diff формат поля diff.
type Diff map[string]any

// Keys возвращает все ключи объекта Diff, по которым можно получить данные.
func (diff Diff) Keys() []string {
	r := make([]string, 0, len(diff))
	for k := range diff {
		r = append(r, k)
	}
	return r
}

// IsExist возвращает true, если fieldName (ключ) присутствует в объекте Diff
func (diff Diff) IsExist(fieldName string) bool {
	if _, ok := diff[fieldName]; ok {
		return true
	}
	return false
}

// OldNew представляет значения до изменения и после изменения.
//
// T может быть любым, ориентируясь на тип сущности EntityType события аудита.
//
// Поля типа T должны иметь тег json с названием соответствующего поля.
type OldNew[T any] struct {
	OldValue T `json:"oldValue"` // Значение до изменения
	NewValue T `json:"newValue"` // Значение после изменения
}

// SalePriceElem представляет объект для отображения изменений цен продажи.
type SalePriceElem struct {
	OldValue struct {
		Uom   string
		Value float64
	} `json:"oldValue"` // Значение до изменения
	NewValue struct {
		Uom   string
		Value float64
	} `json:"newValue"` // Значение после изменения
}

// AuditPosition представляет объект для отображения изменений позиций документов.
type AuditPosition struct {
	Assortment struct {
		Meta Meta   `json:"meta"`
		Name string `json:"name"`
	} `json:"assortment"` // Метаданные позиции
	Uom      string  `json:"uom"`      // Единица измерения
	Quantity float64 `json:"quantity"` // Количество
	Reserve  float64 `json:"reserve"`  // Резерв
	Price    float64 `json:"price"`    // Стоимость позиции в документе
	Discount float64 `json:"discount"` // Скидка позиции в документе
}

// getFieldAndUnmarshall достаёт значение из поля field объекта Diff и пытается привести к типу T.
//
// Возвращает true и T в случае успеха.
func getFieldAndUnmarshall[T any](diff Diff, field string) (bool, T) {
	if positions, ok := diff[field]; ok {
		if p, err := UnmarshallAny[T](positions); err == nil {
			return true, p
		}
	}
	return false, *new(T)
}

// GetPositions возвращает true и позиции изменённого документа, если такие присутствуют в объекте Diff.
func (diff Diff) GetPositions() (bool, []OldNew[AuditPosition]) {
	return getFieldAndUnmarshall[[]OldNew[AuditPosition]](diff, "positions")
}

// GetSalesPrices возвращает true и объект SalePriceElem, если в объекте Diff присутствует поле salePrices.
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

// GetFieldString возвращает true и объект OldNew со значениями типа string, поле fieldName присутствует в объекте Diff.
func (diff Diff) GetFieldString(fieldName string) (bool, OldNew[string]) {
	return getFieldAndUnmarshall[OldNew[string]](diff, fieldName)
}

// GetFieldBool возвращает true и объект OldNew со значениями типа bool, поле fieldName присутствует в объекте Diff.
func (diff Diff) GetFieldBool(fieldName string) (bool, OldNew[bool]) {
	return getFieldAndUnmarshall[OldNew[bool]](diff, fieldName)
}

// GetFieldFloat возвращает true и объект OldNew со значениями типа float64, поле fieldName присутствует в объекте Diff.
func (diff Diff) GetFieldFloat(fieldName string) (bool, OldNew[float64]) {
	return getFieldAndUnmarshall[OldNew[float64]](diff, fieldName)
}

// GetFieldInt возвращает true и объект OldNew со значениями типа int, поле fieldName присутствует в объекте Diff.
func (diff Diff) GetFieldInt(fieldName string) (bool, OldNew[int]) {
	return getFieldAndUnmarshall[OldNew[int]](diff, fieldName)
}

// AuditEventType Действие События.
//
// Возможные значения:
//   - AuditEventRegistration          – Регистрация
//   - AuditEventBulkOperation         – Массовая операция
//   - AuditEventClosePublication      – Удаление публикации
//   - AuditEventCreate                – Создание сущностей
//   - AuditEventDelete                – Удаление сущностей
//   - AuditEventOpenPublication       – Создание публикации
//   - AuditEventPrint                 – Печать документа
//   - AuditEventPutToArchive          – Помещение в архив
//   - AuditEventPutToRecycleBin       – Помещение в корзину
//   - AuditEventReplaceToken          – Смена токена для Точки продаж
//   - AuditEventRestoreFromArchive    – Извлечение из архива
//   - AuditEventRestoreFromRecycleBin – Извлечение из корзины
//   - AuditEventSendEmailFromEntity   – Отправка письма
//   - AuditEventUpdate                – Изменение сущностей
type AuditEventType string

const (
	AuditEventRegistration          AuditEventType = "registration"          // Регистрация
	AuditEventBulkOperation         AuditEventType = "bulkoperation"         // Массовая операция
	AuditEventClosePublication      AuditEventType = "closepublication"      // Удаление публикации
	AuditEventCreate                AuditEventType = "create"                // Создание сущностей
	AuditEventDelete                AuditEventType = "delete"                // Удаление сущностей
	AuditEventOpenPublication       AuditEventType = "openpublication"       // Создание публикации
	AuditEventPrint                 AuditEventType = "print"                 // Печать документа
	AuditEventPutToArchive          AuditEventType = "puttoarchive"          // Помещение в архив
	AuditEventPutToRecycleBin       AuditEventType = "puttorecyclebin"       // Помещение в корзину
	AuditEventReplaceToken          AuditEventType = "replacetoken"          // Смена токена для Точки продаж
	AuditEventRestoreFromArchive    AuditEventType = "restorefromarchive"    // Извлечение из архива
	AuditEventRestoreFromRecycleBin AuditEventType = "restorefromrecyclebin" // Извлечение из корзины
	AuditEventSendEmailFromEntity   AuditEventType = "sendemailfromentity"   // Отправка письма
	AuditEventUpdate                AuditEventType = "update"                // Изменение сущностей
)

// AuditFilters Фильтры аудита.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-fil-try
type AuditFilters struct {
	EventType  Slice[string] `json:"eventType"`  // Действия, по которым могут быть отфильтрованы сущности аудита
	Source     Slice[string] `json:"source"`     // Типы действий, по которым могут быть отфильтрованы сущности аудита
	EntityType Slice[string] `json:"entityType"` // Названия сущностей, по которым могут быть отфильтрованы сущности аудита
}

// String реализует интерфейс [fmt.Stringer].
func (auditFilters AuditFilters) String() string {
	return Stringify(auditFilters)
}

// AuditService Сервис для работы с аудитом.
type AuditService interface {
	// GetContexts выполняет запрос на получение Контекстов Аудита.
	// Принимает контекст context.Context и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetContexts(ctx context.Context, params ...*Params) (*List[Audit], *resty.Response, error)

	// GetEvents выполняет запрос на получение Событий по Контексту AuditEvent.
	// Принимает контекст context.Context и ID контекста Аудита.
	// Возвращает объект List.
	GetEvents(ctx context.Context, id uuid.UUID) (*List[AuditEvent], *resty.Response, error)

	// GetFilters выполняет запрос на получение Фильтров Аудита.
	// Принимает контекст context.Context.
	// Возвращает объект Фильтры аудита.
	GetFilters(ctx context.Context) (*AuditFilters, *resty.Response, error)
}

type auditService struct {
	Endpoint
}

// NewAuditService возвращает сервис для работы с аудитом.
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
	return NewRequestBuilder[AuditFilters](service.client, "audit/metadata/filters").Get(ctx)
}
