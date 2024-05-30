package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// BonusProgramService
// Сервис для работы с бонусными программами.
type BonusProgramService interface {
	GetList(ctx context.Context, params *Params) (*List[BonusProgram], *resty.Response, error)
	Create(ctx context.Context, bonusProgram *BonusProgram, params *Params) (*BonusProgram, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, bonusProgram *BonusProgram, params *Params) (*BonusProgram, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*BonusProgram, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteMany(ctx context.Context, bonusProgramList []*BonusProgram) (*DeleteManyResponse, *resty.Response, error)
}

func NewBonusProgramService(client *Client) BonusProgramService {
	e := NewEndpoint(client, "entity/bonusprogram")
	return newMainService[BonusProgram, any, any, any](e)
}
