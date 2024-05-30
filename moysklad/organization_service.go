package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// OrganizationService
// Сервис для работы с юридическими лицами.
type OrganizationService interface {
	GetList(ctx context.Context, params *Params) (*List[Organization], *resty.Response, error)
	Create(ctx context.Context, organization *Organization, params *Params) (*Organization, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, organizationList []*Organization, params *Params) (*[]Organization, *resty.Response, error)
	DeleteMany(ctx context.Context, organizationList []*Organization) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Organization, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, organization *Organization, params *Params) (*Organization, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeShared, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetAccounts(ctx context.Context, id *uuid.UUID) (*List[AgentAccount], *resty.Response, error)
	GetAccountByID(ctx context.Context, id *uuid.UUID, accountId *uuid.UUID) (*AgentAccount, *resty.Response, error)
	UpdateAccounts(ctx context.Context, id *uuid.UUID, accounts []*AgentAccount) (*[]AgentAccount, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Organization, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
}

func NewOrganizationService(client *Client) OrganizationService {
	e := NewEndpoint(client, "entity/organization")
	return newMainService[Organization, any, MetadataAttributeShared, any](e)
}
