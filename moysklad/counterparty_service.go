package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
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
	endpointSyncID[Counterparty]
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
		endpointSyncID:                 endpointSyncID[Counterparty]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}

// GetAsync Запрос на получения списка Контрагентов (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-spisok-kontragentow
func (s *CounterpartyService) GetAsync(ctx context.Context, params *Params) (*AsyncResultService[List[Counterparty]], *resty.Response, error) {
	return NewRequestBuilder[List[Counterparty]](s.client, s.uri).SetParams(params).Async(ctx)
}

// GetContactPersons Список контактных лиц.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-spisok-kontaktnyh-lic
func (s *CounterpartyService) GetContactPersons(ctx context.Context, id *uuid.UUID, params *Params) (*List[ContactPerson], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons", s.uri, id)
	return NewRequestBuilder[List[ContactPerson]](s.client, path).SetParams(params).Get(ctx)
}

// GetContactPersonById Получить контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-kontaktnoe-lico
func (s *CounterpartyService) GetContactPersonById(ctx context.Context, id, contactPersonId *uuid.UUID) (*ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons/%s", s.uri, id, contactPersonId)
	return NewRequestBuilder[ContactPerson](s.client, path).Get(ctx)
}

// CreateContactPerson Создать контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-sozdat-kontaktnoe-lico
func (s *CounterpartyService) CreateContactPerson(ctx context.Context, id *uuid.UUID, contactPerson *ContactPerson) (*[]ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons", s.uri, id)
	return NewRequestBuilder[[]ContactPerson](s.client, path).Post(ctx, contactPerson)
}

// UpdateContactPerson Изменить контактное лицо.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-izmenit-kontaktnoe-lico
func (s *CounterpartyService) UpdateContactPerson(ctx context.Context, id, contactPersonId *uuid.UUID, contactPerson *ContactPerson) (*ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons/%s", s.uri, id, contactPersonId)
	return NewRequestBuilder[ContactPerson](s.client, path).Put(ctx, contactPerson)
}

// GetNotes Список событий.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-spisok-sobytij
func (s *CounterpartyService) GetNotes(ctx context.Context, id *uuid.UUID) (*List[Note], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", s.uri, id)
	return NewRequestBuilder[List[Note]](s.client, path).Get(ctx)
}

// GetNoteById Получить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-poluchit-sobytie
func (s *CounterpartyService) GetNoteById(ctx context.Context, id, noteId *uuid.UUID) (*Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", s.uri, id, noteId)
	return NewRequestBuilder[Note](s.client, path).Get(ctx)
}

// CreateNote Добавить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-dobawit-sobytie
func (s *CounterpartyService) CreateNote(ctx context.Context, id *uuid.UUID, note *Note) (*[]Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", s.uri, id)
	return NewRequestBuilder[[]Note](s.client, path).Post(ctx, note)
}

// UpdateNote Изменить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-izmenit-sobytie
func (s *CounterpartyService) UpdateNote(ctx context.Context, id, noteId *uuid.UUID, note *Note) (*Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", s.uri, id, noteId)
	return NewRequestBuilder[Note](s.client, path).Put(ctx, note)
}

// DeleteNote Удалить событие.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-udalit-sobytie
func (s *CounterpartyService) DeleteNote(ctx context.Context, id, noteId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", s.uri, id, noteId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
