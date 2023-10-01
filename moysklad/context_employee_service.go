package moysklad

// ContextEmployeeService
// Сервис для работы с контекстом сотрудника.
type ContextEmployeeService struct {
	endpointGetOne[ContextEmployee]
}

func NewContextEmployeeService(client *Client) *ContextEmployeeService {
	e := NewEndpoint(client, "context/employee")
	return &ContextEmployeeService{
		endpointGetOne: endpointGetOne[ContextEmployee]{e},
	}
}
