package moysklad

// UserSettingsService
// Сервис для работы с настройками пользователей.
type UserSettingsService struct {
	endpointGetOne[UserSettings]
	endpointUpdate[UserSettings]
}

func NewContextUserSettingsService(client *Client) *UserSettingsService {
	e := NewEndpoint(client, "context/usersettings")
	return &UserSettingsService{
		endpointGetOne: endpointGetOne[UserSettings]{e},
		endpointUpdate: endpointUpdate[UserSettings]{e},
	}
}
