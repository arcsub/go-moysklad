package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Counterparty Контрагент.
// Ключевое слово: counterparty
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent
type Counterparty struct {
	AccountID          *uuid.UUID                `json:"accountId,omitempty"`          // ID учетной записи
	Accounts           *MetaArray[AgentAccount]  `json:"accounts,omitempty"`           // Массив счетов Контрагентов
	ActualAddress      *string                   `json:"actualAddress,omitempty"`      // Фактический адрес Контрагента
	ActualAddressFull  *Address                  `json:"actualAddressFull,omitempty"`  // Фактический адрес Контрагента с детализацией по отдельным полям
	Archived           *bool                     `json:"archived,omitempty"`           // Добавлен ли Контрагент в архив
	Attributes         *Attributes               `json:"attributes,omitempty"`         // Массив метаданных доп. полей
	BonusPoints        *int                      `json:"bonusPoints,omitempty"`        // Бонусные баллы по активной бонусной программе
	BonusProgram       *BonusProgram             `json:"bonusProgram,omitempty"`       // Метаданные активной Бонусной программы
	Code               *string                   `json:"code,omitempty"`               // Код Контрагента
	CompanyType        CompanyType               `json:"companyType,omitempty"`        // Тип Контрагента. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	ContactPersons     *MetaArray[ContactPerson] `json:"contactpersons,omitempty"`     // Массив контактных лиц фирмы Контрагента
	Created            *Timestamp                `json:"created,omitempty"`            // Момент создания
	Description        *string                   `json:"description,omitempty"`        // Комментарий к Контрагенту
	DiscountCardNumber *string                   `json:"discountCardNumber,omitempty"` // Номер дисконтной карты Контрагента
	Discounts          *Discounts                `json:"discounts,omitempty"`          // Массив скидок Контрагента. Массив может содержать персональные и накопительные скидки. Персональная скидка выводится, если хотя бы раз изменялся процент скидки для контрагента, значение будет указано в поле personalDiscount
	Email              *string                   `json:"email,omitempty"`              // Адрес электронной почты
	ExternalCode       *string                   `json:"externalCode,omitempty"`       // Внешний код Контрагента
	Fax                *string                   `json:"fax,omitempty"`                // Номер факса
	Files              *Files                    `json:"files,omitempty"`              // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group              *Group                    `json:"group,omitempty"`              // Отдел сотрудника
	ID                 *uuid.UUID                `json:"id,omitempty"`                 // ID сущности
	Meta               *Meta                     `json:"meta,omitempty"`               // Метаданные
	Name               *string                   `json:"name,omitempty"`               // Наименование
	Notes              *MetaArray[Note]          `json:"notes,omitempty"`              // Массив событий Контрагента
	Owner              *Employee                 `json:"owner,omitempty"`              // Владелец (Сотрудник)
	Phone              *string                   `json:"phone,omitempty"`              // Номер городского телефона
	PriceType          *PriceType                `json:"priceType,omitempty"`          // Тип цены Контрагента
	SalesAmount        *Decimal                  `json:"salesAmount,omitempty"`        // Сумма продаж
	Shared             *bool                     `json:"shared,omitempty"`             // Общий доступ
	State              *State                    `json:"state,omitempty"`              // Метаданные Статуса Контрагента
	SyncID             *uuid.UUID                `json:"syncId,omitempty"`             // ID синхронизации
	Tags               *Tags                     `json:"tags,omitempty"`               // Группы контрагента
	Updated            *Timestamp                `json:"updated,omitempty"`            // Момент последнего обновления Контрагента
	BirthDate          *Timestamp                `json:"birthDate,omitempty"`          // Дата рождения [15-08-2023]
	CertificateDate    *Timestamp                `json:"certificateDate,omitempty"`    // Дата свидетельства
	CertificateNumber  *string                   `json:"certificateNumber,omitempty"`  // Номер свидетельства
	INN                *string                   `json:"inn,omitempty"`                // ИНН
	KPP                *string                   `json:"kpp,omitempty"`                // КПП
	LegalAddress       *string                   `json:"legalAddress,omitempty"`       // Юридический адрес
	LegalAddressFull   *Address                  `json:"legalAddressFull,omitempty"`   // Структурированный Юридический адрес
	LegalTitle         *string                   `json:"legalTitle,omitempty"`         // Полное наименование
	OGRN               *string                   `json:"ogrn,omitempty"`               // ОГРН
	OGRNIP             *string                   `json:"ogrnip,omitempty"`             // ОГРНИП
	OKPO               *string                   `json:"okpo,omitempty"`               // ОКПО
	Sex                SexType                   `json:"sex,omitempty"`                // Пол Контрагента [15-08-2023]
}

func (c Counterparty) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c Counterparty) GetMeta() *Meta {
	return c.Meta
}

func (c Counterparty) MetaType() MetaType {
	return MetaTypeCounterparty
}

type DiscountData struct {
	Discount             *Discount `json:"discount,omitempty"`             // Скидка
	PersonalDiscount     *float64  `json:"personalDiscount,omitempty"`     // Значение персональной скидки
	DemandSumCorrection  *float64  `json:"demandSumCorrection,omitempty"`  // Коррекция суммы накоплений по скидке
	AccumulationDiscount *float64  `json:"accumulationDiscount,omitempty"` // Значение накопительной скидки
}

func (d DiscountData) String() string {
	return Stringify(d)
}

// CounterpartySettings Настройки справочника контрагентов
// Ключевое слово: counterpartysettings
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-nastrojki-sprawochnika-kontragentow
type CounterpartySettings struct {
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные
	CreateShared    *bool            `json:"createShared,omitempty"`    // Создавать новые документы с меткой «Общий»
	UniqueCodeRules *UniqueCodeRules `json:"uniqueCodeRules,omitempty"` // Настройки кодов контрагентов
}

func (c CounterpartySettings) String() string {
	return Stringify(c)
}

func (c CounterpartySettings) MetaType() MetaType {
	return MetaTypeCounterpartySettings
}

// Note Событие Контрагента.
// Ключевое слово: note
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-attributy-suschnosti-adres-sobytiq-kontragenta
type Note struct {
	AccountID         *uuid.UUID    `json:"accountId,omitempty"`         // ID учетной записи
	Agent             *Counterparty `json:"agent,omitempty"`             // Метаданные Контрагента
	Author            *Employee     `json:"author,omitempty"`            // Метаданные Сотрудника - создателя события (администратор аккаунта, если автор - приложение)
	AuthorApplication *Application  `json:"authorApplication,omitempty"` // Метаданные Приложения - создателя события
	Created           *Timestamp    `json:"created,omitempty"`           // Момент создания события Контрагента
	Description       *string       `json:"description,omitempty"`       // Текст события Контрагента
	ID                *uuid.UUID    `json:"id,omitempty"`                // ID сущности
	Meta              *Meta         `json:"meta,omitempty"`              // Метаданные
}

func (n Note) String() string {
	return Stringify(n)
}

func (n Note) MetaType() MetaType {
	return MetaTypeNote
}

// SexType Пол Контрагента
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-pol-kontragenta
type SexType string

const (
	SexMale   SexType = "MALE"   // Мужской
	SexFemale SexType = "FEMALE" // Женский
)

type Tags = []string

// CounterpartyService
// Сервис для работы с контрагентами.
type CounterpartyService interface {
	GetList(ctx context.Context, params *Params) (*List[Counterparty], *resty.Response, error)
	Create(ctx context.Context, counterparty *Counterparty, params *Params) (*Counterparty, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, counterpartyList []*Counterparty, params *Params) (*[]Counterparty, *resty.Response, error)
	DeleteMany(ctx context.Context, counterpartyList []*Counterparty) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Counterparty, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, counterparty *Counterparty, params *Params) (*Counterparty, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataCounterparty, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetSettings(ctx context.Context) (*CounterpartySettings, *resty.Response, error)
	UpdateSettings(ctx context.Context, settings *CounterpartySettings) (*CounterpartySettings, *resty.Response, error)
	GetAccounts(ctx context.Context, id *uuid.UUID) (*List[AgentAccount], *resty.Response, error)
	GetAccountByID(ctx context.Context, id *uuid.UUID, accountId *uuid.UUID) (*AgentAccount, *resty.Response, error)
	UpdateAccounts(ctx context.Context, id *uuid.UUID, accounts []*AgentAccount) (*[]AgentAccount, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Counterparty, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetAsync(ctx context.Context, params *Params) (AsyncResultService[List[Counterparty]], *resty.Response, error)
	GetContactPersons(ctx context.Context, id *uuid.UUID, params *Params) (*List[ContactPerson], *resty.Response, error)
	GetContactPersonById(ctx context.Context, id, contactPersonID *uuid.UUID) (*ContactPerson, *resty.Response, error)
	CreateContactPerson(ctx context.Context, id *uuid.UUID, contactPerson *ContactPerson) (*[]ContactPerson, *resty.Response, error)
	UpdateContactPerson(ctx context.Context, id, contactPersonID *uuid.UUID, contactPerson *ContactPerson) (*ContactPerson, *resty.Response, error)
	GetNotes(ctx context.Context, id *uuid.UUID) (*List[Note], *resty.Response, error)
	GetNoteById(ctx context.Context, id, noteID *uuid.UUID) (*Note, *resty.Response, error)
	CreateNote(ctx context.Context, id *uuid.UUID, note *Note) (*[]Note, *resty.Response, error)
	UpdateNote(ctx context.Context, id, noteID *uuid.UUID, note *Note) (*Note, *resty.Response, error)
	DeleteNote(ctx context.Context, id, noteID *uuid.UUID) (bool, *resty.Response, error)
}

type counterpartyService struct {
	Endpoint
	endpointGetList[Counterparty]
	endpointCreate[Counterparty]
	endpointCreateUpdateMany[Counterparty]
	endpointDeleteMany[Counterparty]
	endpointDelete
	endpointGetById[Counterparty]
	endpointUpdate[Counterparty]
	endpointMetadata[MetadataCounterparty]
	endpointAttributes
	endpointSettings[CounterpartySettings]
	endpointAccounts
	endpointSyncID[Counterparty]
	endpointNamedFilter
}

func NewCounterpartyService(client *Client) CounterpartyService {
	e := NewEndpoint(client, "entity/counterparty")
	return &counterpartyService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Counterparty]{e},
		endpointCreate:           endpointCreate[Counterparty]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Counterparty]{e},
		endpointDeleteMany:       endpointDeleteMany[Counterparty]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[Counterparty]{e},
		endpointUpdate:           endpointUpdate[Counterparty]{e},
		endpointMetadata:         endpointMetadata[MetadataCounterparty]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSettings:         endpointSettings[CounterpartySettings]{e},
		endpointAccounts:         endpointAccounts{e},
		endpointSyncID:           endpointSyncID[Counterparty]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}

// GetAsync Запрос на получения списка Контрагентов (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-spisok-kontragentow
func (s *counterpartyService) GetAsync(ctx context.Context, params *Params) (AsyncResultService[List[Counterparty]], *resty.Response, error) {
	return NewRequestBuilder[List[Counterparty]](s.client, s.uri).SetParams(params).Async(ctx)
}

// GetContactPersons Список контактных лиц.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-spisok-kontaktnyh-lic
func (s *counterpartyService) GetContactPersons(ctx context.Context, id *uuid.UUID, params *Params) (*List[ContactPerson], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons", s.uri, id)
	return NewRequestBuilder[List[ContactPerson]](s.client, path).SetParams(params).Get(ctx)
}

// GetContactPersonById Получить контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-kontaktnoe-lico
func (s *counterpartyService) GetContactPersonById(ctx context.Context, id, contactPersonId *uuid.UUID) (*ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons/%s", s.uri, id, contactPersonId)
	return NewRequestBuilder[ContactPerson](s.client, path).Get(ctx)
}

// CreateContactPerson Создать контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-sozdat-kontaktnoe-lico
func (s *counterpartyService) CreateContactPerson(ctx context.Context, id *uuid.UUID, contactPerson *ContactPerson) (*[]ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons", s.uri, id)
	return NewRequestBuilder[[]ContactPerson](s.client, path).Post(ctx, contactPerson)
}

// UpdateContactPerson Изменить контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-izmenit-kontaktnoe-lico
func (s *counterpartyService) UpdateContactPerson(ctx context.Context, id, contactPersonId *uuid.UUID, contactPerson *ContactPerson) (*ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons/%s", s.uri, id, contactPersonId)
	return NewRequestBuilder[ContactPerson](s.client, path).Put(ctx, contactPerson)
}

// GetNotes Список событий.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-spisok-sobytij
func (s *counterpartyService) GetNotes(ctx context.Context, id *uuid.UUID) (*List[Note], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", s.uri, id)
	return NewRequestBuilder[List[Note]](s.client, path).Get(ctx)
}

// GetNoteById Получить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-sobytie
func (s *counterpartyService) GetNoteById(ctx context.Context, id, noteId *uuid.UUID) (*Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", s.uri, id, noteId)
	return NewRequestBuilder[Note](s.client, path).Get(ctx)
}

// CreateNote Добавить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-dobawit-sobytie
func (s *counterpartyService) CreateNote(ctx context.Context, id *uuid.UUID, note *Note) (*[]Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", s.uri, id)
	return NewRequestBuilder[[]Note](s.client, path).Post(ctx, note)
}

// UpdateNote Изменить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-izmenit-sobytie
func (s *counterpartyService) UpdateNote(ctx context.Context, id, noteId *uuid.UUID, note *Note) (*Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", s.uri, id, noteId)
	return NewRequestBuilder[Note](s.client, path).Put(ctx, note)
}

// DeleteNote Удалить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-udalit-sobytie
func (s *counterpartyService) DeleteNote(ctx context.Context, id, noteId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", s.uri, id, noteId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
