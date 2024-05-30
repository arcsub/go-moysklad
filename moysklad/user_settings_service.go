package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// UserSettingsService
// Сервис для работы с настройками пользователей.
type UserSettingsService interface {
	Get(ctx context.Context, params *Params) (*UserSettings, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, userSettings *UserSettings, params *Params) (*UserSettings, *resty.Response, error)
}

func NewContextUserSettingsService(client *Client) UserSettingsService {
	e := NewEndpoint(client, "context/usersettings")
	return newMainService[UserSettings, any, any, any](e)
}
