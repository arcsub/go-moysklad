package moysklad

// BonusProgramService
// Сервис для работы с бонусными программами.
type BonusProgramService struct {
	endpointGetList[BonusProgram]
	endpointCreate[BonusProgram]
	endpointUpdate[BonusProgram]
	endpointGetById[BonusProgram]
	endpointDelete
	endpointDeleteMany[BonusProgram]
}

func NewBonusProgramService(client *Client) *BonusProgramService {
	e := NewEndpoint(client, "entity/bonusprogram")
	return &BonusProgramService{
		endpointGetList:    endpointGetList[BonusProgram]{e},
		endpointCreate:     endpointCreate[BonusProgram]{e},
		endpointUpdate:     endpointUpdate[BonusProgram]{e},
		endpointGetById:    endpointGetById[BonusProgram]{e},
		endpointDelete:     endpointDelete{e},
		endpointDeleteMany: endpointDeleteMany[BonusProgram]{e},
	}
}
