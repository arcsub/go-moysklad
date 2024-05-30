package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductFolderService
// Сервис для работы с группами товаров.
type ProductFolderService interface {
	GetList(ctx context.Context, params *Params) (*List[ProductFolder], *resty.Response, error)
	Create(ctx context.Context, productFolder *ProductFolder, params *Params) (*ProductFolder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productFolderList []*ProductFolder, params *Params) (*[]ProductFolder, *resty.Response, error)
	DeleteMany(ctx context.Context, productFolderList []*ProductFolder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttribute, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProductFolder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, productFolder *ProductFolder, params *Params) (*ProductFolder, *resty.Response, error)
}

func NewProductFolderService(client *Client) ProductFolderService {
	e := NewEndpoint(client, "entity/productfolder")
	return newMainService[ProductFolder, any, MetadataAttribute, any](e)
}
