package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

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
