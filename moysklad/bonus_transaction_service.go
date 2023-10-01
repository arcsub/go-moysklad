package moysklad

// BonusTransactionService
// Сервис для работы с бонусными операциями.
type BonusTransactionService struct {
	endpointGetList[BonusTransaction]
	endpointCreate[BonusTransaction]
	endpointCreateUpdateDeleteMany[BonusTransaction]
	endpointDelete
	endpointDeleteMany[BonusTransaction]
	endpointGetById[BonusTransaction]
	endpointUpdate[BonusTransaction]
}

func NewBonusTransactionService(client *Client) *BonusTransactionService {
	e := NewEndpoint(client, "entity/bonustransaction")
	return &BonusTransactionService{
		endpointGetList:                endpointGetList[BonusTransaction]{e},
		endpointCreate:                 endpointCreate[BonusTransaction]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[BonusTransaction]{e},
		endpointDelete:                 endpointDelete{e},
		endpointDeleteMany:             endpointDeleteMany[BonusTransaction]{e},
		endpointGetById:                endpointGetById[BonusTransaction]{e},
		endpointUpdate:                 endpointUpdate[BonusTransaction]{e},
	}
}
