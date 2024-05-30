package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// ContextEmployeeService
// Сервис для работы с контекстом сотрудника.
type ContextEmployeeService interface {
	Get(ctx context.Context, params *Params) (*ContextEmployee, *resty.Response, error)
}

func NewContextEmployeeService(client *Client) ContextEmployeeService {
	e := NewEndpoint(client, "context/employee")
	return newMainService[ContextEmployee, any, any, any](e)
}
