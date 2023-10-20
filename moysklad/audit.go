package moysklad

import (
	"github.com/google/uuid"
)

// Audit Контексты Аудита.
// Ключевое слово: audit
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-kontexty
type Audit struct {
	EntityType    string           `json:"entityType"`    // Название сущности (поле присутствует, только если оно одинаково у всех Событий в рамках данного Контекста)
	EventType     AuditEventType   `json:"eventType"`     // Действие Событий (поле присутствует, только если оно одинаково у всех Событий в рамках данного Контекста)
	Events        MetaWrapper      `json:"events"`        // Список метаданных Событий аудита
	ID            uuid.UUID        `json:"id"`            // ID Контекста
	Info          AuditContextInfo `json:"info"`          // Краткое описание
	Meta          Meta             `json:"meta"`          // Метаданные сущности Контекста
	Moment        Timestamp        `json:"moment"`        // Дата изменения
	ObjectCount   int              `json:"objectCount"`   // Количество измененных объектов
	ObjectType    string           `json:"objectType"`    // Тип сущностей, с которыми связанно данное изменение. Поле присутствует только для entityType = entitysettings или statesettings или templatesettings
	Source        string           `json:"source"`        // Тип изменения
	SupportAccess bool             `json:"supportAccess"` // Был ли доступ произведен поддержкой от имени пользователя. Флаг отсутствует, если значение false
	UID           string           `json:"uid"`           // Логин Сотрудника
}

func (a Audit) String() string {
	return Stringify(a)
}

func (a Audit) MetaType() MetaType {
	return MetaTypeAudit
}

type AuditContextInfo struct {
	AdditionalInfo string `json:"additionalInfo"`
}

// AuditEvent Событие аудита.
// Ключевое слово: auditevent
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq
type AuditEvent struct {
	AdditionalInfo string         `json:"additionalInfo"` // Дополнительная информация о Событии
	Audit          MetaWrapper    `json:"audit"`          // Метаданные контекста
	Diff           any            `json:"diff"`           // TODO Изменения, произошедшие в Событии, в специальном формате diff, описанном в разделе Формат поля diff
	Entity         MetaWrapper    `json:"entity"`         // Метаданные сущности. Не будет выводиться только для товаров, услуг, модификаций, комплектов удаленных до 20.08.2017
	EntityType     string         `json:"entityType"`     // Название сущности
	EventType      AuditEventType `json:"eventType"`      // Действие События
	Moment         Timestamp      `json:"moment"`         // Время создания события
	Name           string         `json:"name"`           // Имя сущности
	ObjectCount    int            `json:"objectCount"`    // Количество измененных объектов
	ObjectType     string         `json:"objectType"`     // Тип сущностей, с которыми связанно данное изменение. Поле присутствует только для entityType = entitysettings или statesettings или templatesettings
	Source         string         `json:"source"`         // Тип изменения
	SupportAccess  bool           `json:"supportAccess"`  // Был ли доступ произведен поддержкой от имени пользователя. Флаг отсутствует, если значение false
	UID            string         `json:"uid"`            // Логин Сотрудника
}

func (a AuditEvent) String() string {
	return Stringify(a)
}

func (a AuditEvent) MetaType() MetaType {
	return MetaTypeAuditEvent
}

// TODO: поле diff в событии аудита может иметь разный тип данных: {},[] и ""

// AuditEventDiff Формат поля diff
//
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq-format-polq-diff
type AuditEventDiff struct {
	// Событие регистрации
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq-sobytie-registracii
	Account string `json:"account,omitempty"` // Название аккаунта
	Country string `json:"country,omitempty"` // Конфигурация аккаунта (страна)

	// События публикации документов
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq-sobytiq-publikacii-dokumentow
	TemplateName    string `json:"templateName,omitempty"`    // Название шаблона
	PublicationHref string `json:"publicationHref,omitempty"` // Ссылка на публикацию

	// События отправки писем
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq-sobytiq-otprawki-pisem
	SenderEmail  string `json:"senderEmail,omitempty"`  // Почта отправителя письма
	TargetEmail  string `json:"targetEmail,omitempty"`  // Почта получателя письма
	SubjectEmail string `json:"subjectEmail,omitempty"` // Тема письма
	Text         string `json:"text,omitempty"`         // Текст письма

	// События удаления сущностей
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq-sobytiq-udaleniq-suschnostej
	AttributeName string `json:"attributeName,omitempty"` // Название атрибута сущности
	OldValue      any    `json:"oldValue,omitempty"`      // Значение атрибута до удаления

	// События обновления сущностей, перемещения/восстановления из корзины, перемещение/восстановление из архива
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/audit/#audit-audit-sobytiq-sobytiq-obnowleniq-suschnostej-peremescheniq-wosstanowleniq-iz-korziny-peremeschenie-wosstanowlenie-iz-arhiwa
	NewValue any `json:"newValue,omitempty"` // Значение атрибута после обновления
}

func (a AuditEventDiff) String() string {
	return Stringify(a)
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

func (a AuditFilters) String() string {
	return Stringify(a)
}
