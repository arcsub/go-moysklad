package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

// Notification TODO: Общие атрибуты уведомлений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-lenta-uwedomlenij-obschie-atributy-uwedomlenij
type Notification struct {
	Meta        Meta      `json:"meta"`
	Created     Timestamp `json:"created"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	AccountID   uuid.UUID `json:"accountId"`
	ID          uuid.UUID `json:"id"`
	Read        bool      `json:"read"`
}

func (n Notification) String() string {
	return Stringify(n)
}

func (n Notification) MetaType() MetaType {
	return MetaTypeNotification
}

// NotificationService
// Сервис для работы с уведомлениями.
type NotificationService interface {
	GetList(ctx context.Context, params *Params) (*List[Notification], *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Notification, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	MarkAsRead(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	MarkAsReadAll(ctx context.Context) (bool, *resty.Response, error)
}

type notificationService struct {
	Endpoint
	endpointGetList[Notification]
	endpointGetById[Notification]
	endpointDelete
}

func NewNotificationService(client *Client) NotificationService {
	e := NewEndpoint(client, "notification")
	return &notificationService{
		Endpoint:        e,
		endpointGetList: endpointGetList[Notification]{e},
		endpointGetById: endpointGetById[Notification]{e},
		endpointDelete:  endpointDelete{e},
	}
}

// MarkAsRead Отметить Уведомление как прочитанное.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-uwedomlenie-otmetit-uwedomlenie-kak-prochitannoe
func (s *notificationService) MarkAsRead(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/markasread", s.uri, id)
	_, resp, err := NewRequestBuilder[any](s.client, path).Put(ctx, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}

// MarkAsReadAll Отметить все Уведомления как прочитанные.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-uwedomlenie-otmetit-wse-uwedomleniq-kak-prochitannye
func (s *notificationService) MarkAsReadAll(ctx context.Context) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/markasreadall", s.uri)
	_, resp, err := NewRequestBuilder[any](s.client, path).Put(ctx, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}
