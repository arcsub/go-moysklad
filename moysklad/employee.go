package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

// Employee Сотрудник.
// Ключевое слово: employee
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik
type Employee struct {
	ID           *uuid.UUID            `json:"id,omitempty"`
	Owner        *Employee             `json:"owner,omitempty"`
	Image        *Image                `json:"image,omitempty"`
	INN          *string               `json:"inn,omitempty"`
	Code         *string               `json:"code,omitempty"`
	Created      *Timestamp            `json:"created,omitempty"`
	Description  *string               `json:"description,omitempty"`
	Email        *string               `json:"email,omitempty"`
	ExternalCode *string               `json:"externalCode,omitempty"`
	FirstName    *string               `json:"firstName,omitempty"`
	FullName     *string               `json:"fullName,omitempty"`
	Group        *Group                `json:"group,omitempty"`
	Updated      *Timestamp            `json:"updated,omitempty"`
	AccountID    *uuid.UUID            `json:"accountId,omitempty"`
	Cashiers     *MetaArray[Cashier]   `json:"cashiers,omitempty"`
	LastName     *string               `json:"lastName,omitempty"`
	Meta         *Meta                 `json:"meta,omitempty"`
	MiddleName   *string               `json:"middleName,omitempty"`
	Name         *string               `json:"name,omitempty"`
	Archived     *bool                 `json:"archived,omitempty"`
	Phone        *string               `json:"phone,omitempty"`
	Position     *string               `json:"position,omitempty"`
	Shared       *bool                 `json:"shared,omitempty"`
	ShortFio     *string               `json:"shortFio,omitempty"`
	UID          *string               `json:"uid,omitempty"`
	Attributes   Slice[AttributeValue] `json:"attributes,omitempty"`
}

func (employee Employee) Clean() *Employee {
	return &Employee{Meta: employee.Meta}
}

func (employee Employee) GetID() uuid.UUID {
	return Deref(employee.ID)
}

func (employee Employee) GetOwner() Employee {
	return Deref(employee.Owner)
}

func (employee Employee) GetImage() Image {
	return Deref(employee.Image)
}

func (employee Employee) GetINN() string {
	return Deref(employee.INN)
}

func (employee Employee) GetCode() string {
	return Deref(employee.Code)
}

func (employee Employee) GetCreated() Timestamp {
	return Deref(employee.Created)
}

func (employee Employee) GetDescription() string {
	return Deref(employee.Description)
}

func (employee Employee) GetEmail() string {
	return Deref(employee.Email)
}

func (employee Employee) GetExternalCode() string {
	return Deref(employee.ExternalCode)
}

func (employee Employee) GetFirstName() string {
	return Deref(employee.FirstName)
}

func (employee Employee) GetFullName() string {
	return Deref(employee.FullName)
}

func (employee Employee) GetGroup() Group {
	return Deref(employee.Group)
}

func (employee Employee) GetUpdated() Timestamp {
	return Deref(employee.Updated)
}

func (employee Employee) GetAccountID() uuid.UUID {
	return Deref(employee.AccountID)
}

func (employee Employee) GetCashiers() MetaArray[Cashier] {
	return Deref(employee.Cashiers)
}

func (employee Employee) GetLastName() string {
	return Deref(employee.LastName)
}

func (employee Employee) GetMeta() Meta {
	return Deref(employee.Meta)
}

func (employee Employee) GetMiddleName() string {
	return Deref(employee.MiddleName)
}

func (employee Employee) GetName() string {
	return Deref(employee.Name)
}

func (employee Employee) GetArchived() bool {
	return Deref(employee.Archived)
}

func (employee Employee) GetPhone() string {
	return Deref(employee.Phone)
}

func (employee Employee) GetPosition() string {
	return Deref(employee.Position)
}

func (employee Employee) GetShared() bool {
	return Deref(employee.Shared)
}

func (employee Employee) GetShortFio() string {
	return Deref(employee.ShortFio)
}

func (employee Employee) GetUID() string {
	return Deref(employee.UID)
}

func (employee Employee) GetAttributes() Slice[AttributeValue] {
	return employee.Attributes
}

func (employee *Employee) SetOwner(owner *Employee) *Employee {
	employee.Owner = owner.Clean()
	return employee
}

func (employee *Employee) SetImage(image *Image) *Employee {
	employee.Image = image
	return employee
}

func (employee *Employee) SetINN(inn string) *Employee {
	employee.INN = &inn
	return employee
}

func (employee *Employee) SetCode(code string) *Employee {
	employee.Code = &code
	return employee
}

func (employee *Employee) SetDescription(description string) *Employee {
	employee.Description = &description
	return employee
}

func (employee *Employee) SetEmail(email string) *Employee {
	employee.Email = &email
	return employee
}

func (employee *Employee) SetExternalCode(externalCode string) *Employee {
	employee.ExternalCode = &externalCode
	return employee
}

func (employee *Employee) SetFirstName(firstName string) *Employee {
	employee.FirstName = &firstName
	return employee
}

func (employee *Employee) SetGroup(group *Group) *Employee {
	employee.Group = group.Clean()
	return employee
}

func (employee *Employee) SetLastName(lastName string) *Employee {
	employee.LastName = &lastName
	return employee
}

func (employee *Employee) SetMeta(meta *Meta) *Employee {
	employee.Meta = meta
	return employee
}

func (employee *Employee) SetMiddleName(middleName string) *Employee {
	employee.MiddleName = &middleName
	return employee
}

func (employee *Employee) SetArchived(archived bool) *Employee {
	employee.Archived = &archived
	return employee
}

func (employee *Employee) SetPhone(phone string) *Employee {
	employee.Phone = &phone
	return employee
}

func (employee *Employee) SetPosition(position string) *Employee {
	employee.Position = &position
	return employee
}

func (employee *Employee) SetShared(shared bool) *Employee {
	employee.Shared = &shared
	return employee
}

func (employee *Employee) SetAttributes(attributes Slice[AttributeValue]) *Employee {
	employee.Attributes = attributes
	return employee
}

func (employee Employee) String() string {
	return Stringify(employee)
}

func (employee Employee) MetaType() MetaType {
	return MetaTypeEmployee
}

// MailActivationRequired структура ответа на запрос активации сотрудника
// Если поле равно true и сотрудник ранее не был активен, это означает, что на указанную у сотрудника почту было выслано письмо со ссылкой на вход для сотрудника.
// Если поле равно false, то можно использовать ранее заданный пароль для данного пользователя
type MailActivationRequired struct {
	MailActivationRequired bool `json:"mailActivationRequired"`
}

// EmployeePermission Права Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-rabota-s-prawami-sotrudnika
type EmployeePermission struct {
	AuthorizedIpNetmask *string  `json:"authorizedIpNetmask,omitempty"`
	AuthorizedIpNetwork *string  `json:"authorizedIpNetwork,omitempty"`
	Email               *string  `json:"email,omitempty"`
	Group               *Group   `json:"group,omitempty"`
	IsActive            *bool    `json:"isActive,omitempty"`
	Login               *string  `json:"login,omitempty"`
	Role                *Role    `json:"role,omitempty"`
	AuthorizedHosts     []string `json:"authorizedHosts,omitempty"`
}

func (employeePermission EmployeePermission) GetAuthorizedIpNetmask() string {
	return Deref(employeePermission.AuthorizedIpNetmask)
}

func (employeePermission EmployeePermission) GetAuthorizedIpNetwork() string {
	return Deref(employeePermission.AuthorizedIpNetwork)
}

func (employeePermission EmployeePermission) GetEmail() string {
	return Deref(employeePermission.Email)
}

func (employeePermission EmployeePermission) GetGroup() Group {
	return Deref(employeePermission.Group)
}

func (employeePermission EmployeePermission) GetIsActive() bool {
	return Deref(employeePermission.IsActive)
}

func (employeePermission EmployeePermission) GetLogin() string {
	return Deref(employeePermission.Login)
}

func (employeePermission EmployeePermission) GetRole() Role {
	return Deref(employeePermission.Role)
}

func (employeePermission EmployeePermission) GetAuthorizedHosts() []string {
	return employeePermission.AuthorizedHosts
}

func (employeePermission *EmployeePermission) SetAuthorizedIpNetmask(authorizedIpNetmask string) *EmployeePermission {
	employeePermission.AuthorizedIpNetmask = &authorizedIpNetmask
	return employeePermission
}

func (employeePermission *EmployeePermission) SetAuthorizedIpNetwork(authorizedIpNetwork string) *EmployeePermission {
	employeePermission.AuthorizedIpNetwork = &authorizedIpNetwork
	return employeePermission
}

func (employeePermission *EmployeePermission) SetEmail(email string) *EmployeePermission {
	employeePermission.Email = &email
	return employeePermission
}

func (employeePermission *EmployeePermission) SetGroup(group *Group) *EmployeePermission {
	employeePermission.Group = group.Clean()
	return employeePermission
}

func (employeePermission *EmployeePermission) SetIsActive(isActive bool) *EmployeePermission {
	employeePermission.IsActive = &isActive
	return employeePermission
}

func (employeePermission *EmployeePermission) SetLogin(login string) *EmployeePermission {
	employeePermission.Login = &login
	return employeePermission
}

func (employeePermission *EmployeePermission) SetRole(role *Role) *EmployeePermission {
	employeePermission.Role = role.Clean()
	return employeePermission
}

func (employeePermission *EmployeePermission) SetAuthorizedHosts(authorizedHosts []string) *EmployeePermission {
	employeePermission.AuthorizedHosts = authorizedHosts
	return employeePermission
}

func (employeePermission EmployeePermission) String() string {
	return Stringify(employeePermission)
}

// EmployeeService
// Сервис для работы с сотрудниками.
type EmployeeService interface {
	GetList(ctx context.Context, params *Params) (*List[Employee], *resty.Response, error)
	Create(ctx context.Context, employee *Employee, params *Params) (*Employee, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, employeeList Slice[Employee], params *Params) (*Slice[Employee], *resty.Response, error)
	DeleteMany(ctx context.Context, employeeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList Slice[Attribute]) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Employee, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, employee *Employee, params *Params) (*Employee, *resty.Response, error)
	GetPermissions(ctx context.Context, id uuid.UUID) (*EmployeePermission, *resty.Response, error)
	UpdatePermissions(ctx context.Context, id uuid.UUID, permissions *EmployeePermission) (*EmployeePermission, *resty.Response, error)
	Activate(ctx context.Context, id uuid.UUID, permissions *EmployeePermission) (*MailActivationRequired, *resty.Response, error)
	Deactivate(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	ResetPassword(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

type employeeService struct {
	Endpoint
	endpointGetList[Employee]
	endpointCreate[Employee]
	endpointCreateUpdateMany[Employee]
	endpointDeleteMany[Employee]
	endpointDelete
	endpointMetadata[MetaAttributesSharedWrapper]
	endpointAttributes
	endpointGetByID[Employee]
	endpointUpdate[Employee]
}

func NewEmployeeService(client *Client) EmployeeService {
	e := NewEndpoint(client, "entity/employee")
	return &employeeService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Employee]{e},
		endpointCreate:           endpointCreate[Employee]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Employee]{e},
		endpointDeleteMany:       endpointDeleteMany[Employee]{e},
		endpointDelete:           endpointDelete{e},
		endpointMetadata:         endpointMetadata[MetaAttributesSharedWrapper]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointGetByID:          endpointGetByID[Employee]{e},
		endpointUpdate:           endpointUpdate[Employee]{e},
	}
}

// GetPermissions Запрос на получение информации о правах Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-poluchit-informaciu-o-prawah-sotrudnika
func (service *employeeService) GetPermissions(ctx context.Context, id uuid.UUID) (*EmployeePermission, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/security", service.uri, id)
	return NewRequestBuilder[EmployeePermission](service.client, path).Get(ctx)
}

// UpdatePermissions Запрос на изменение информации о правах Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-izmenit-informaciu-o-prawah-sotrudnika
func (service *employeeService) UpdatePermissions(ctx context.Context, id uuid.UUID, permissions *EmployeePermission) (*EmployeePermission, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/security", service.uri, id)
	return NewRequestBuilder[EmployeePermission](service.client, path).Put(ctx, permissions)
}

// Activate Активация Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-aktiwaciq-sotrudnika
func (service *employeeService) Activate(ctx context.Context, id uuid.UUID, permissions *EmployeePermission) (*MailActivationRequired, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/access/activate", service.uri, id)
	return NewRequestBuilder[MailActivationRequired](service.client, path).Put(ctx, permissions)
}

// Deactivate Деактивация Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-deaktiwaciq-sotrudnika
func (service *employeeService) Deactivate(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/access/deactivate", service.uri, id)
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusNoContent, resp, nil
}

// ResetPassword Сброс пароля Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-sbros-parolq-sotrudnika
func (service *employeeService) ResetPassword(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/access/resetpassword", id)
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusNoContent, resp, nil
}
