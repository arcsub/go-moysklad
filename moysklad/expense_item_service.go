package moysklad

// ExpenseItemService
// Сервис для работы со статьями расходов.
type ExpenseItemService struct {
	endpointGetList[ExpenseItem]
	endpointCreate[ExpenseItem]
	endpointCreateUpdateDeleteMany[ExpenseItem]
	endpointDelete
	endpointGetById[ExpenseItem]
	endpointUpdate[ExpenseItem]
	endpointRemove
}

func NewExpenseItemService(client *Client) *ExpenseItemService {
	e := NewEndpoint(client, "entity/expenseitem")
	return &ExpenseItemService{
		endpointGetList:                endpointGetList[ExpenseItem]{e},
		endpointCreate:                 endpointCreate[ExpenseItem]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[ExpenseItem]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[ExpenseItem]{e},
		endpointUpdate:                 endpointUpdate[ExpenseItem]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
