package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// CounterpartyService
// Сервис для работы с контрагентами.
type CounterpartyService struct {
	Endpoint
	endpointGetList[Counterparty]
	endpointCreate[Counterparty]
	endpointCreateUpdateDeleteMany[Counterparty]
	endpointDelete
	endpointGetById[Counterparty]
	endpointUpdate[Counterparty]
	endpointMetadata[MetadataCounterparty]
	endpointAttributes
	endpointSettings[CounterpartySettings]
	endpointAccounts
	endpointSyncId[Counterparty]
	endpointNamedFilter
}

func NewCounterpartyService(client *Client) *CounterpartyService {
	e := NewEndpoint(client, "entity/counterparty")
	return &CounterpartyService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[Counterparty]{e},
		endpointCreate:                 endpointCreate[Counterparty]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Counterparty]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Counterparty]{e},
		endpointUpdate:                 endpointUpdate[Counterparty]{e},
		endpointMetadata:               endpointMetadata[MetadataCounterparty]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointSettings:               endpointSettings[CounterpartySettings]{e},
		endpointAccounts:               endpointAccounts{e},
		endpointSyncId:                 endpointSyncId[Counterparty]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}

// GetAsync Запрос на получения списка Контрагентов (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-spisok-kontragentow
func (s *CounterpartyService) GetAsync(ctx context.Context, params *Params) (*AsyncResultService[List[Counterparty]], *Response, error) {
	return NewRequestBuilder[List[Counterparty]](s.Endpoint, ctx).WithParams(params).Async()
}

// GetContactPersons Список контактных лиц.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-spisok-kontaktnyh-lic
func (s *CounterpartyService) GetContactPersons(ctx context.Context, id *uuid.UUID, params *Params) (*List[ContactPerson], *Response, error) {
	path := fmt.Sprintf("%s/contactpersons", id)
	return NewRequestBuilder[List[ContactPerson]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

// GetContactPersonById Получить контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-kontaktnoe-lico
func (s *CounterpartyService) GetContactPersonById(ctx context.Context, id, contactPersonId uuid.UUID) (*ContactPerson, *Response, error) {
	path := fmt.Sprintf("%s/contactpersons/%s", id, contactPersonId)
	return NewRequestBuilder[ContactPerson](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateContactPerson Создать контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-sozdat-kontaktnoe-lico
func (s *CounterpartyService) CreateContactPerson(ctx context.Context, id *uuid.UUID, contactPerson *ContactPerson) (*Slice[ContactPerson], *Response, error) {
	path := fmt.Sprintf("%s/contactpersons", id)
	return NewRequestBuilder[Slice[ContactPerson]](s.Endpoint, ctx).WithPath(path).WithBody(contactPerson).Post()
}

// UpdateContactPerson Изменить контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-izmenit-kontaktnoe-lico
func (s *CounterpartyService) UpdateContactPerson(ctx context.Context, id, contactPersonId uuid.UUID, contactPerson *ContactPerson) (*ContactPerson, *Response, error) {
	path := fmt.Sprintf("%s/contactpersons/%s", id, contactPersonId)
	return NewRequestBuilder[ContactPerson](s.Endpoint, ctx).WithPath(path).WithBody(contactPerson).Put()
}

// GetNotes Список событий.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-spisok-sobytij
func (s *CounterpartyService) GetNotes(ctx context.Context, id *uuid.UUID) (*List[Note], *Response, error) {
	path := fmt.Sprintf("%s/notes", id)
	return NewRequestBuilder[List[Note]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetNoteById Получить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-sobytie
func (s *CounterpartyService) GetNoteById(ctx context.Context, id, noteId uuid.UUID) (*Note, *Response, error) {
	path := fmt.Sprintf("%s/notes/%s", id, noteId)
	return NewRequestBuilder[Note](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateNote Добавить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-dobawit-sobytie
func (s *CounterpartyService) CreateNote(ctx context.Context, id *uuid.UUID, note *Note) (*Slice[Note], *Response, error) {
	path := fmt.Sprintf("%s/notes", id)
	return NewRequestBuilder[Slice[Note]](s.Endpoint, ctx).WithPath(path).WithBody(note).Post()
}

// UpdateNote Изменить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-izmenit-sobytie
func (s *CounterpartyService) UpdateNote(ctx context.Context, id, noteId uuid.UUID, note *Note) (*Note, *Response, error) {
	path := fmt.Sprintf("%s/notes/%s", id, noteId)
	return NewRequestBuilder[Note](s.Endpoint, ctx).WithPath(path).WithBody(note).Put()
}

// DeleteNote Удалить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-udalit-sobytie
func (s *CounterpartyService) DeleteNote(ctx context.Context, id, noteId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/notes/%s", id, noteId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}
