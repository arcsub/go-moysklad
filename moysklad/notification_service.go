package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

// NotificationService
// Сервис для работы с уведомлениями.
type NotificationService struct {
	Endpoint
	endpointGetList[Notification]
	endpointGetById[Notification]
	endpointDelete
}

func NewNotificationService(client *Client) *NotificationService {
	e := NewEndpoint(client, "notification")
	return &NotificationService{
		Endpoint:        e,
		endpointGetList: endpointGetList[Notification]{e},
		endpointGetById: endpointGetById[Notification]{e},
		endpointDelete:  endpointDelete{e},
	}
}

// MarkAsRead Отметить Уведомление как прочитанное.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-uwedomlenie-otmetit-uwedomlenie-kak-prochitannoe
func (s *NotificationService) MarkAsRead(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/markasread", s.uri, id)
	resp, err := s.client.client.R().SetContext(ctx).Put(path)
	return resp.StatusCode() == http.StatusOK, resp, err
}

// MarkAsReadAll Отметить все Уведомления как прочитанные.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-uwedomlenie-otmetit-wse-uwedomleniq-kak-prochitannye
func (s *NotificationService) MarkAsReadAll(ctx context.Context) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/markasreadall", s.uri)
	resp, err := s.client.client.R().SetContext(ctx).Put(path)
	return resp.StatusCode() == http.StatusOK, resp, err
}
