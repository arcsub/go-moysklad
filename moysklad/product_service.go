package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductService
// Сервис для работы с товарами.
type ProductService interface {
	GetList(ctx context.Context, params *Params) (*List[Product], *resty.Response, error)
	Create(ctx context.Context, product *Product, params *Params) (*Product, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productList []*Product, params *Params) (*[]Product, *resty.Response, error)
	DeleteMany(ctx context.Context, productList []*Product) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeShared, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Product, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, product *Product, params *Params) (*Product, *resty.Response, error)
	GetImages(ctx context.Context, id *uuid.UUID) (*MetaArray[Image], *resty.Response, error)
	CreateImage(ctx context.Context, id *uuid.UUID, image *Image) (*[]*Image, *resty.Response, error)
	UpdateImages(ctx context.Context, id *uuid.UUID, images []*Image) (*[]Image, *resty.Response, error)
	DeleteImage(ctx context.Context, id *uuid.UUID, imageId *uuid.UUID) (bool, *resty.Response, error)
	DeleteImages(ctx context.Context, id *uuid.UUID, images []*Image) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Product, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetAudit(ctx context.Context, id *uuid.UUID, params *Params) (*List[AuditEvent], *resty.Response, error)
	PrintLabel(ctx context.Context, id *uuid.UUID, PrintLabelArg *PrintLabelArg) (*PrintFile, *resty.Response, error)
}

func NewProductService(client *Client) ProductService {
	e := NewEndpoint(client, "entity/product")
	return newMainService[Product, any, MetadataAttributeShared, any](e)
}
