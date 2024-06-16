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
	Name               *string                   `json:"name,omitempty"`
	OKPO               *string                   `json:"okpo,omitempty"`
	ActualAddress      *string                   `json:"actualAddress,omitempty"`
	AccountID          *uuid.UUID                `json:"accountId,omitempty"`
	Archived           *bool                     `json:"archived,omitempty"`
	Notes              *MetaArray[Note]          `json:"notes,omitempty"`
	BonusPoints        *int                      `json:"bonusPoints,omitempty"`
	BonusProgram       *NullValue[BonusProgram]  `json:"bonusProgram,omitempty"`
	Code               *string                   `json:"code,omitempty"`
	OGRNIP             *string                   `json:"ogrnip,omitempty"`
	ContactPersons     *MetaArray[ContactPerson] `json:"contactpersons,omitempty"`
	Created            *Timestamp                `json:"created,omitempty"`
	Description        *string                   `json:"description,omitempty"`
	DiscountCardNumber *string                   `json:"discountCardNumber,omitempty"`
	Discounts          *DiscountsData            `json:"discounts,omitempty"`
	Email              *string                   `json:"email,omitempty"`
	ExternalCode       *string                   `json:"externalCode,omitempty"`
	Owner              *Employee                 `json:"owner,omitempty"`
	Files              *MetaArray[File]          `json:"files,omitempty"`
	Group              *Group                    `json:"group,omitempty"`
	ID                 *uuid.UUID                `json:"id,omitempty"`
	Meta               *Meta                     `json:"meta,omitempty"`
	ActualAddressFull  *Address                  `json:"actualAddressFull,omitempty"`
	Accounts           *MetaArray[AgentAccount]  `json:"accounts,omitempty"`
	Fax                *string                   `json:"fax,omitempty"`
	Phone              *string                   `json:"phone,omitempty"`
	PriceType          *PriceType                `json:"priceType,omitempty"`
	SalesAmount        *float64                  `json:"salesAmount,omitempty"`
	Shared             *bool                     `json:"shared,omitempty"`
	State              *NullValue[State]         `json:"state,omitempty"`
	SyncID             *uuid.UUID                `json:"syncId,omitempty"`
	Tags               Slice[string]             `json:"tags,omitempty"`
	Updated            *Timestamp                `json:"updated,omitempty"`
	BirthDate          *Timestamp                `json:"birthDate,omitempty"`
	CertificateDate    *Timestamp                `json:"certificateDate,omitempty"`
	CertificateNumber  *string                   `json:"certificateNumber,omitempty"`
	INN                *string                   `json:"inn,omitempty"`
	KPP                *string                   `json:"kpp,omitempty"`
	LegalAddress       *string                   `json:"legalAddress,omitempty"`
	LegalAddressFull   *Address                  `json:"legalAddressFull,omitempty"`
	LegalTitle         *string                   `json:"legalTitle,omitempty"`
	OGRN               *string                   `json:"ogrn,omitempty"`
	CompanyType        CompanyType               `json:"companyType,omitempty"`
	Sex                Sex                       `json:"sex,omitempty"`
	Attributes         Slice[AttributeValue]     `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (counterparty Counterparty) Clean() *Counterparty {
	return &Counterparty{Meta: counterparty.Meta}
}

func (counterparty Counterparty) GetName() string {
	return Deref(counterparty.Name)
}

func (counterparty Counterparty) GetOKPO() string {
	return Deref(counterparty.OKPO)
}

func (counterparty Counterparty) GetActualAddress() string {
	return Deref(counterparty.ActualAddress)
}

func (counterparty Counterparty) GetAccountID() uuid.UUID {
	return Deref(counterparty.AccountID)
}

func (counterparty Counterparty) GetArchived() bool {
	return Deref(counterparty.Archived)
}

func (counterparty Counterparty) GetNotes() MetaArray[Note] {
	return Deref(counterparty.Notes)
}

func (counterparty Counterparty) GetBonusPoints() int {
	return Deref(counterparty.BonusPoints)
}

func (counterparty Counterparty) GetBonusProgram() BonusProgram {
	return counterparty.BonusProgram.Get()
}

func (counterparty Counterparty) GetCode() string {
	return Deref(counterparty.Code)
}

func (counterparty Counterparty) GetOGRNIP() string {
	return Deref(counterparty.OGRNIP)
}

func (counterparty Counterparty) GetContactPersons() MetaArray[ContactPerson] {
	return Deref(counterparty.ContactPersons)
}

func (counterparty Counterparty) GetCreated() Timestamp {
	return Deref(counterparty.Created)
}

func (counterparty Counterparty) GetDescription() string {
	return Deref(counterparty.Description)
}

func (counterparty Counterparty) GetDiscountCardNumber() string {
	return Deref(counterparty.DiscountCardNumber)
}

func (counterparty Counterparty) GetDiscounts() DiscountsData {
	return Deref(counterparty.Discounts)
}

func (counterparty Counterparty) GetEmail() string {
	return Deref(counterparty.Email)
}

func (counterparty Counterparty) GetExternalCode() string {
	return Deref(counterparty.ExternalCode)
}

func (counterparty Counterparty) GetOwner() Employee {
	return Deref(counterparty.Owner)
}

func (counterparty Counterparty) GetFiles() MetaArray[File] {
	return Deref(counterparty.Files)
}

func (counterparty Counterparty) GetGroup() Group {
	return Deref(counterparty.Group)
}

func (counterparty Counterparty) GetID() uuid.UUID {
	return Deref(counterparty.ID)
}

func (counterparty Counterparty) GetMeta() Meta {
	return Deref(counterparty.Meta)
}

func (counterparty Counterparty) GetActualAddressFull() Address {
	return Deref(counterparty.ActualAddressFull)
}

func (counterparty Counterparty) GetAccounts() MetaArray[AgentAccount] {
	return Deref(counterparty.Accounts)
}

func (counterparty Counterparty) GetFax() string {
	return Deref(counterparty.Fax)
}

func (counterparty Counterparty) GetPhone() string {
	return Deref(counterparty.Phone)
}

func (counterparty Counterparty) GetPriceType() PriceType {
	return Deref(counterparty.PriceType)
}

func (counterparty Counterparty) GetSalesAmount() float64 {
	return Deref(counterparty.SalesAmount)
}

func (counterparty Counterparty) GetShared() bool {
	return Deref(counterparty.Shared)
}

func (counterparty Counterparty) GetState() State {
	return counterparty.State.Get()
}

func (counterparty Counterparty) GetSyncID() uuid.UUID {
	return Deref(counterparty.SyncID)
}

func (counterparty Counterparty) GetTags() Slice[string] {
	return counterparty.Tags
}

func (counterparty Counterparty) GetUpdated() Timestamp {
	return Deref(counterparty.Updated)
}

func (counterparty Counterparty) GetBirthDate() Timestamp {
	return Deref(counterparty.BirthDate)
}

func (counterparty Counterparty) GetCertificateDate() Timestamp {
	return Deref(counterparty.CertificateDate)
}

func (counterparty Counterparty) GetCertificateNumber() string {
	return Deref(counterparty.CertificateNumber)
}

func (counterparty Counterparty) GetINN() string {
	return Deref(counterparty.INN)
}

func (counterparty Counterparty) GetKPP() string {
	return Deref(counterparty.KPP)
}

func (counterparty Counterparty) GetLegalAddress() string {
	return Deref(counterparty.LegalAddress)
}

func (counterparty Counterparty) GetLegalAddressFull() Address {
	return Deref(counterparty.LegalAddressFull)
}

func (counterparty Counterparty) GetLegalTitle() string {
	return Deref(counterparty.LegalTitle)
}

func (counterparty Counterparty) GetOGRN() string {
	return Deref(counterparty.OGRN)
}

func (counterparty Counterparty) GetCompanyType() CompanyType {
	return counterparty.CompanyType
}

func (counterparty Counterparty) GetSex() Sex {
	return counterparty.Sex
}

func (counterparty Counterparty) GetAttributes() Slice[AttributeValue] {
	return counterparty.Attributes
}

func (counterparty *Counterparty) SetName(name string) *Counterparty {
	counterparty.Name = &name
	return counterparty
}

func (counterparty *Counterparty) SetOKPO(okpo string) *Counterparty {
	counterparty.OKPO = &okpo
	return counterparty
}

func (counterparty *Counterparty) SetActualAddress(actualAddress string) *Counterparty {
	counterparty.ActualAddress = &actualAddress
	return counterparty
}

func (counterparty *Counterparty) SetArchived(archived bool) *Counterparty {
	counterparty.Archived = &archived
	return counterparty
}

func (counterparty *Counterparty) SetNotes(notes Slice[Note]) *Counterparty {
	counterparty.Notes = NewMetaArrayRows(notes)
	return counterparty
}

func (counterparty *Counterparty) SetBonusProgram(bonusProgram *BonusProgram) *Counterparty {
	counterparty.BonusProgram = NewNullValueWith(bonusProgram.Clean())
	return counterparty
}

func (counterparty *Counterparty) SetNullBonusProgram() *Counterparty {
	counterparty.BonusProgram = NewNullValue[BonusProgram]()
	return counterparty
}

func (counterparty *Counterparty) SetCode(code string) *Counterparty {
	counterparty.Code = &code
	return counterparty
}

func (counterparty *Counterparty) SetOGRNIP(ogrnip string) *Counterparty {
	counterparty.OGRNIP = &ogrnip
	return counterparty
}

func (counterparty *Counterparty) SetContactPersons(contactPersons Slice[ContactPerson]) *Counterparty {
	counterparty.ContactPersons = NewMetaArrayRows(contactPersons)
	return counterparty
}

func (counterparty *Counterparty) SetCreated(created *Timestamp) *Counterparty {
	counterparty.Created = created
	return counterparty
}

func (counterparty *Counterparty) SetDescription(description string) *Counterparty {
	counterparty.Description = &description
	return counterparty
}

func (counterparty *Counterparty) SetDiscountCardNumber(discountCardNumber string) *Counterparty {
	counterparty.DiscountCardNumber = &discountCardNumber
	return counterparty
}

func (counterparty *Counterparty) SetDiscounts(discounts *DiscountsData) *Counterparty {
	counterparty.Discounts = discounts
	return counterparty
}

func (counterparty *Counterparty) SetEmail(email string) *Counterparty {
	counterparty.Email = &email
	return counterparty
}

func (counterparty *Counterparty) SetExternalCode(externalCode string) *Counterparty {
	counterparty.ExternalCode = &externalCode
	return counterparty
}

func (counterparty *Counterparty) SetOwner(owner *Employee) *Counterparty {
	counterparty.Owner = owner.Clean()
	return counterparty
}

func (counterparty *Counterparty) SetFiles(files Slice[File]) *Counterparty {
	counterparty.Files = NewMetaArrayRows(files)
	return counterparty
}

func (counterparty *Counterparty) SetGroup(group *Group) *Counterparty {
	counterparty.Group = group.Clean()
	return counterparty
}

func (counterparty *Counterparty) SetMeta(meta *Meta) *Counterparty {
	counterparty.Meta = meta
	return counterparty
}

func (counterparty *Counterparty) SetActualAddressFull(actualAddressFull *Address) *Counterparty {
	counterparty.ActualAddressFull = actualAddressFull
	return counterparty
}

func (counterparty *Counterparty) SetAccounts(accounts Slice[AgentAccount]) *Counterparty {
	counterparty.Accounts = NewMetaArrayRows(accounts)
	return counterparty
}

func (counterparty *Counterparty) SetFax(fax string) *Counterparty {
	counterparty.Fax = &fax
	return counterparty
}

func (counterparty *Counterparty) SetPhone(phone string) *Counterparty {
	counterparty.Phone = &phone
	return counterparty
}

func (counterparty *Counterparty) SetPriceType(priceType *PriceType) *Counterparty {
	counterparty.PriceType = priceType
	return counterparty
}

func (counterparty *Counterparty) SetShared(shared bool) *Counterparty {
	counterparty.Shared = &shared
	return counterparty
}

func (counterparty *Counterparty) SetState(state *State) *Counterparty {
	counterparty.State = NewNullValueWith(state.Clean())
	return counterparty
}

func (counterparty *Counterparty) SetNullState() *Counterparty {
	counterparty.State = NewNullValue[State]()
	return counterparty
}

func (counterparty *Counterparty) SetSyncID(syncID uuid.UUID) *Counterparty {
	counterparty.SyncID = &syncID
	return counterparty
}

func (counterparty *Counterparty) SetTags(tags Slice[string]) *Counterparty {
	counterparty.Tags = tags
	return counterparty
}

func (counterparty *Counterparty) SetBirthDate(birthDate *Timestamp) *Counterparty {
	counterparty.BirthDate = birthDate
	return counterparty
}

func (counterparty *Counterparty) SetCertificateDate(certificateDate *Timestamp) *Counterparty {
	counterparty.CertificateDate = certificateDate
	return counterparty
}

func (counterparty *Counterparty) SetCertificateNumber(certificateNumber string) *Counterparty {
	counterparty.CertificateNumber = &certificateNumber
	return counterparty
}

func (counterparty *Counterparty) SetINN(inn string) *Counterparty {
	counterparty.INN = &inn
	return counterparty
}

func (counterparty *Counterparty) SetKPP(kpp string) *Counterparty {
	counterparty.KPP = &kpp
	return counterparty
}

func (counterparty *Counterparty) SetLegalAddress(legalAddress string) *Counterparty {
	counterparty.LegalAddress = &legalAddress
	return counterparty
}

func (counterparty *Counterparty) SetLegalAddressFull(legalAddressFull *Address) *Counterparty {
	counterparty.LegalAddressFull = legalAddressFull
	return counterparty
}

func (counterparty *Counterparty) SetLegalTitle(legalTitle string) *Counterparty {
	counterparty.LegalTitle = &legalTitle
	return counterparty
}

func (counterparty *Counterparty) SetOGRN(ogrn string) *Counterparty {
	counterparty.OGRN = &ogrn
	return counterparty
}

func (counterparty *Counterparty) SetCompanyType(companyType CompanyType) *Counterparty {
	counterparty.CompanyType = companyType
	return counterparty
}

func (counterparty *Counterparty) SetSex(sex Sex) *Counterparty {
	counterparty.Sex = sex
	return counterparty
}

func (counterparty *Counterparty) SetAttributes(attributes Slice[AttributeValue]) *Counterparty {
	counterparty.Attributes = attributes
	return counterparty
}

func (counterparty Counterparty) String() string {
	return Stringify(counterparty)
}

func (counterparty Counterparty) MetaType() MetaType {
	return MetaTypeCounterparty
}

type DiscountsData struct {
	Discount             MetaWrapper `json:"discount,omitempty"`             // Скидка
	PersonalDiscount     float64     `json:"personalDiscount,omitempty"`     // Значение персональной скидки
	DemandSumCorrection  float64     `json:"demandSumCorrection,omitempty"`  // Коррекция суммы накоплений по скидке
	AccumulationDiscount float64     `json:"accumulationDiscount,omitempty"` // Значение накопительной скидки
}

func (discountsData DiscountsData) String() string {
	return Stringify(discountsData)
}

// CounterpartySettings Настройки справочника контрагентов
// Ключевое слово: counterpartysettings
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-nastrojki-sprawochnika-kontragentow
type CounterpartySettings struct {
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные
	CreateShared    *bool            `json:"createShared,omitempty"`    // Создавать новые документы с меткой «Общий»
	UniqueCodeRules *UniqueCodeRules `json:"uniqueCodeRules,omitempty"` // Настройки кодов контрагентов
}

func (counterpartySettings CounterpartySettings) GetMeta() Meta {
	return Deref(counterpartySettings.Meta)
}

func (counterpartySettings CounterpartySettings) GetCreateShared() bool {
	return Deref(counterpartySettings.CreateShared)
}

func (counterpartySettings CounterpartySettings) GetUniqueCodeRules() UniqueCodeRules {
	return Deref(counterpartySettings.UniqueCodeRules)
}

func (counterpartySettings *CounterpartySettings) SetCreateShared(createShared bool) *CounterpartySettings {
	counterpartySettings.CreateShared = &createShared
	return counterpartySettings
}

func (counterpartySettings *CounterpartySettings) SetUniqueCodeRules(uniqueCodeRules *UniqueCodeRules) *CounterpartySettings {
	counterpartySettings.UniqueCodeRules = uniqueCodeRules
	return counterpartySettings
}

func (counterpartySettings CounterpartySettings) String() string {
	return Stringify(counterpartySettings)
}

func (counterpartySettings CounterpartySettings) MetaType() MetaType {
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

func (note Note) GetAccountID() uuid.UUID {
	return Deref(note.AccountID)
}

func (note Note) GetAgent() Counterparty {
	return Deref(note.Agent)
}

func (note Note) GetAuthor() Employee {
	return Deref(note.Author)
}

func (note Note) GetAuthorApplication() Application {
	return Deref(note.AuthorApplication)
}

func (note Note) GetCreated() Timestamp {
	return Deref(note.Created)
}

func (note Note) GetDescription() string {
	return Deref(note.Description)
}

func (note Note) GetID() uuid.UUID {
	return Deref(note.ID)
}

func (note Note) GetMeta() Meta {
	return Deref(note.Meta)
}

func (note *Note) SetDescription(description string) *Note {
	note.Description = &description
	return note
}

func (note *Note) SetMeta(meta *Meta) *Note {
	note.Meta = meta
	return note
}

func (note Note) String() string {
	return Stringify(note)
}

func (note Note) MetaType() MetaType {
	return MetaTypeNote
}

// Sex Пол Контрагента
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-pol-kontragenta
type Sex string

const (
	SexMale   Sex = "MALE"   // Мужской
	SexFemale Sex = "FEMALE" // Женский
)

// CounterpartyService
// Сервис для работы с контрагентами.
type CounterpartyService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Counterparty], *resty.Response, error)
	Create(ctx context.Context, counterparty *Counterparty, params ...*Params) (*Counterparty, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, counterpartyList Slice[Counterparty], params ...*Params) (*Slice[Counterparty], *resty.Response, error)
	DeleteMany(ctx context.Context, counterpartyList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Counterparty, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, counterparty *Counterparty, params ...*Params) (*Counterparty, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaTagsWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList Slice[Attribute]) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	GetSettings(ctx context.Context) (*CounterpartySettings, *resty.Response, error)
	UpdateSettings(ctx context.Context, settings *CounterpartySettings) (*CounterpartySettings, *resty.Response, error)
	GetAccounts(ctx context.Context, id uuid.UUID) (*List[AgentAccount], *resty.Response, error)
	GetAccountByID(ctx context.Context, id uuid.UUID, accountId uuid.UUID) (*AgentAccount, *resty.Response, error)
	UpdateAccounts(ctx context.Context, id uuid.UUID, accounts Slice[AgentAccount]) (*MetaArray[AgentAccount], *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Counterparty, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[Counterparty]], *resty.Response, error)
	GetContactPersons(ctx context.Context, id uuid.UUID, params ...*Params) (*List[ContactPerson], *resty.Response, error)
	GetContactPersonById(ctx context.Context, id, contactPersonID uuid.UUID) (*ContactPerson, *resty.Response, error)
	CreateContactPerson(ctx context.Context, id uuid.UUID, contactPerson *ContactPerson) (*Slice[ContactPerson], *resty.Response, error)
	UpdateContactPerson(ctx context.Context, id, contactPersonID uuid.UUID, contactPerson *ContactPerson) (*ContactPerson, *resty.Response, error)
	GetNotes(ctx context.Context, id uuid.UUID) (*List[Note], *resty.Response, error)
	GetNoteByID(ctx context.Context, id, noteID uuid.UUID) (*Note, *resty.Response, error)
	CreateNote(ctx context.Context, id uuid.UUID, note *Note) (*MetaArray[Note], *resty.Response, error)
	UpdateNote(ctx context.Context, id, noteID uuid.UUID, note *Note) (*Note, *resty.Response, error)
	DeleteNote(ctx context.Context, id, noteID uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFiles(ctx context.Context, id uuid.UUID, files Slice[File]) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id uuid.UUID, files []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
}

type counterpartyService struct {
	Endpoint
	endpointGetList[Counterparty]
	endpointCreate[Counterparty]
	endpointCreateUpdateMany[Counterparty]
	endpointDeleteMany[Counterparty]
	endpointDelete
	endpointGetByID[Counterparty]
	endpointUpdate[Counterparty]
	endpointMetadata[MetaTagsWrapper]
	endpointAttributes
	endpointSettings[CounterpartySettings]
	endpointAccounts
	endpointSyncID[Counterparty]
	endpointNamedFilter
	endpointStates
	endpointFiles
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
		endpointGetByID:          endpointGetByID[Counterparty]{e},
		endpointUpdate:           endpointUpdate[Counterparty]{e},
		endpointMetadata:         endpointMetadata[MetaTagsWrapper]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSettings:         endpointSettings[CounterpartySettings]{e},
		endpointAccounts:         endpointAccounts{e},
		endpointSyncID:           endpointSyncID[Counterparty]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
	}
}

// GetAsync Запрос на получения списка Контрагентов (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-spisok-kontragentow
func (service *counterpartyService) GetAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[Counterparty]], *resty.Response, error) {
	return NewRequestBuilder[List[Counterparty]](service.client, service.uri).SetParams(params...).Async(ctx)
}

// GetContactPersons Список контактных лиц.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-spisok-kontaktnyh-lic
func (service *counterpartyService) GetContactPersons(ctx context.Context, id uuid.UUID, params ...*Params) (*List[ContactPerson], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons", service.uri, id)
	return NewRequestBuilder[List[ContactPerson]](service.client, path).SetParams(params...).Get(ctx)
}

// GetContactPersonById Получить контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-kontaktnoe-lico
func (service *counterpartyService) GetContactPersonById(ctx context.Context, id, contactPersonID uuid.UUID) (*ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons/%s", service.uri, id, contactPersonID)
	return NewRequestBuilder[ContactPerson](service.client, path).Get(ctx)
}

// CreateContactPerson Создать контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-sozdat-kontaktnoe-lico
func (service *counterpartyService) CreateContactPerson(ctx context.Context, id uuid.UUID, contactPerson *ContactPerson) (*Slice[ContactPerson], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons", service.uri, id)
	return NewRequestBuilder[Slice[ContactPerson]](service.client, path).Post(ctx, contactPerson)
}

// UpdateContactPerson Изменить контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-izmenit-kontaktnoe-lico
func (service *counterpartyService) UpdateContactPerson(ctx context.Context, id, contactPersonID uuid.UUID, contactPerson *ContactPerson) (*ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons/%s", service.uri, id, contactPersonID)
	return NewRequestBuilder[ContactPerson](service.client, path).Put(ctx, contactPerson)
}

// GetNotes Список событий.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-spisok-sobytij
func (service *counterpartyService) GetNotes(ctx context.Context, id uuid.UUID) (*List[Note], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, id)
	return NewRequestBuilder[List[Note]](service.client, path).Get(ctx)
}

// GetNoteByID Получить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-sobytie
func (service *counterpartyService) GetNoteByID(ctx context.Context, id, noteID uuid.UUID) (*Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, id, noteID)
	return NewRequestBuilder[Note](service.client, path).Get(ctx)
}

// CreateNote Добавить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-dobawit-sobytie
func (service *counterpartyService) CreateNote(ctx context.Context, id uuid.UUID, note *Note) (*MetaArray[Note], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, id)
	return NewRequestBuilder[MetaArray[Note]](service.client, path).Post(ctx, note)
}

// UpdateNote Изменить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-izmenit-sobytie
func (service *counterpartyService) UpdateNote(ctx context.Context, id, noteID uuid.UUID, note *Note) (*Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, id, noteID)
	return NewRequestBuilder[Note](service.client, path).Put(ctx, note)
}

// DeleteNote Удалить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-udalit-sobytie
func (service *counterpartyService) DeleteNote(ctx context.Context, id, noteID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, id, noteID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}
