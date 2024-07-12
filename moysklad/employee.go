package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// Employee Сотрудник.
//
// Код сущности: employee
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik
type Employee struct {
	ID           *uuid.UUID          `json:"id,omitempty"`           // ID Сотрудника
	Owner        *Employee           `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Image        *NullValue[Image]   `json:"image,omitempty"`        // Фотография сотрудника
	INN          *string             `json:"inn,omitempty"`          // ИНН сотрудника (в формате ИНН физического лица)
	Code         *string             `json:"code,omitempty"`         // Код Сотрудника
	Created      *Timestamp          `json:"created,omitempty"`      // Момент создания Сотрудника
	Description  *string             `json:"description,omitempty"`  // Комментарий к Сотруднику
	Email        *string             `json:"email,omitempty"`        // Электронная почта сотрудника
	ExternalCode *string             `json:"externalCode,omitempty"` // Внешний код Сотрудника
	FirstName    *string             `json:"firstName,omitempty"`    // Имя
	FullName     *string             `json:"fullName,omitempty"`     // Имя Отчество Фамилия
	Group        *Group              `json:"group,omitempty"`        // Отдел сотрудника
	Updated      *Timestamp          `json:"updated,omitempty"`      // Момент последнего обновления Сотрудника
	AccountID    *uuid.UUID          `json:"accountId,omitempty"`    // ID учётной записи
	Cashiers     *MetaArray[Cashier] `json:"cashiers,omitempty"`     // Массив кассиров
	LastName     *string             `json:"lastName,omitempty"`     // Фамилия
	Meta         *Meta               `json:"meta,omitempty"`         // Метаданные Сотрудника
	MiddleName   *string             `json:"middleName,omitempty"`   // Отчество
	Name         *string             `json:"name,omitempty"`         // Наименование Сотрудника
	Archived     *bool               `json:"archived,omitempty"`     // Добавлен ли Сотрудник в архив
	Phone        *string             `json:"phone,omitempty"`        // Телефон сотрудника
	Position     *string             `json:"position,omitempty"`     // Должность сотрудника
	Salary       *Salary             `json:"salary,omitempty"`       // Оклад сотрудника
	Shared       *bool               `json:"shared,omitempty"`       // Общий доступ
	ShortFio     *string             `json:"shortFio,omitempty"`     // Краткое ФИО
	UID          *string             `json:"uid,omitempty"`          // Логин Сотрудника
	Attributes   Slice[Attribute]    `json:"attributes,omitempty"`   // Дополнительные поля Сотрудника
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (employee Employee) Clean() *Employee {
	if employee.Meta == nil {
		return nil
	}
	return &Employee{Meta: employee.Meta}
}

// AsEmployeeAgent реализует интерфейс [AgentEmployeeConverter].
func (employee Employee) AsEmployeeAgent() *Agent {
	return employee.AsAgent()
}

// AsAgent реализует интерфейс [AgentConverter].
func (employee Employee) AsAgent() *Agent {
	if employee.Meta == nil {
		return nil
	}
	return &Agent{Meta: employee.Meta}
}

// GetID возвращает ID Сотрудника.
func (employee Employee) GetID() uuid.UUID {
	return Deref(employee.ID)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (employee Employee) GetOwner() Employee {
	return Deref(employee.Owner)
}

// GetImage возвращает Фотографию сотрудника.
func (employee Employee) GetImage() Image {
	return Deref(employee.Image).getValue()
}

// GetINN возвращает ИНН.
func (employee Employee) GetINN() string {
	return Deref(employee.INN)
}

// GetCode возвращает Код Сотрудника.
func (employee Employee) GetCode() string {
	return Deref(employee.Code)
}

// GetCreated возвращает Момент создания Сотрудника.
func (employee Employee) GetCreated() time.Time {
	return Deref(employee.Created).Time()
}

// GetDescription возвращает Комментарий к Сотруднику.
func (employee Employee) GetDescription() string {
	return Deref(employee.Description)
}

// GetEmail возвращает Адрес электронной почты Сотрудника.
func (employee Employee) GetEmail() string {
	return Deref(employee.Email)
}

// GetExternalCode возвращает Внешний код Сотрудника.
func (employee Employee) GetExternalCode() string {
	return Deref(employee.ExternalCode)
}

// GetFirstName возвращает Имя.
func (employee Employee) GetFirstName() string {
	return Deref(employee.FirstName)
}

// GetFullName возвращает Имя Отчество Фамилия.
func (employee Employee) GetFullName() string {
	return Deref(employee.FullName)
}

// GetGroup возвращает Отдел сотрудника.
func (employee Employee) GetGroup() Group {
	return Deref(employee.Group)
}

// GetUpdated возвращает Момент последнего обновления Сотрудника.
func (employee Employee) GetUpdated() time.Time {
	return Deref(employee.Updated).Time()
}

// GetAccountID возвращает ID учётной записи.
func (employee Employee) GetAccountID() uuid.UUID {
	return Deref(employee.AccountID)
}

// GetCashiers возвращает Массив кассиров.
func (employee Employee) GetCashiers() MetaArray[Cashier] {
	return Deref(employee.Cashiers)
}

// GetLastName возвращает Фамилию.
func (employee Employee) GetLastName() string {
	return Deref(employee.LastName)
}

// GetMeta возвращает Метаданные Сотрудника.
func (employee Employee) GetMeta() Meta {
	return Deref(employee.Meta)
}

// GetMiddleName возвращает Отчество.
func (employee Employee) GetMiddleName() string {
	return Deref(employee.MiddleName)
}

// GetName возвращает Наименование Сотрудника.
func (employee Employee) GetName() string {
	return Deref(employee.Name)
}

// GetArchived возвращает true, если Сотрудник добавлен в архив.
func (employee Employee) GetArchived() bool {
	return Deref(employee.Archived)
}

// GetPhone возвращает Телефон сотрудника.
func (employee Employee) GetPhone() string {
	return Deref(employee.Phone)
}

// GetPosition возвращает Должность сотрудника.
func (employee Employee) GetPosition() string {
	return Deref(employee.Position)
}

// GetSalary возвращает Оклад сотрудника.
func (employee Employee) GetSalary() Salary {
	return Deref(employee.Salary)
}

// GetShared возвращает флаг Общего доступа.
func (employee Employee) GetShared() bool {
	return Deref(employee.Shared)
}

// GetShortFio возвращает Краткое ФИО.
func (employee Employee) GetShortFio() string {
	return Deref(employee.ShortFio)
}

// GetUID возвращает Логин Сотрудника.
func (employee Employee) GetUID() string {
	return Deref(employee.UID)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (employee Employee) GetAttributes() Slice[Attribute] {
	return employee.Attributes
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (employee *Employee) SetOwner(owner *Employee) *Employee {
	if owner != nil {
		employee.Owner = owner.Clean()
	}
	return employee
}

// SetImage устанавливает Фотографию сотрудника.
//
// Передача nil передаёт сброс значения (null).
func (employee *Employee) SetImage(image *Image) *Employee {
	employee.Image = NewNullValue(image)
	return employee
}

// SetINN устанавливает ИНН.
func (employee *Employee) SetINN(inn string) *Employee {
	employee.INN = &inn
	return employee
}

// SetCode устанавливает Код Сотрудника.
func (employee *Employee) SetCode(code string) *Employee {
	employee.Code = &code
	return employee
}

// SetDescription устанавливает Комментарий к Сотруднику.
func (employee *Employee) SetDescription(description string) *Employee {
	employee.Description = &description
	return employee
}

// SetEmail устанавливает Адрес электронной почты сотрудника.
func (employee *Employee) SetEmail(email string) *Employee {
	employee.Email = &email
	return employee
}

// SetExternalCode устанавливает Внешний код Сотрудника.
func (employee *Employee) SetExternalCode(externalCode string) *Employee {
	employee.ExternalCode = &externalCode
	return employee
}

// SetFirstName устанавливает Имя.
func (employee *Employee) SetFirstName(firstName string) *Employee {
	employee.FirstName = &firstName
	return employee
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (employee *Employee) SetGroup(group *Group) *Employee {
	if group != nil {
		employee.Group = group.Clean()
	}
	return employee
}

// SetLastName устанавливает Фамилию.
func (employee *Employee) SetLastName(lastName string) *Employee {
	employee.LastName = &lastName
	return employee
}

// SetMeta устанавливает Метаданные Сотрудника.
func (employee *Employee) SetMeta(meta *Meta) *Employee {
	employee.Meta = meta
	return employee
}

// SetMiddleName устанавливает Отчество.
func (employee *Employee) SetMiddleName(middleName string) *Employee {
	employee.MiddleName = &middleName
	return employee
}

// SetArchived устанавливает флаг нахождения сотрудника в архиве.
func (employee *Employee) SetArchived(archived bool) *Employee {
	employee.Archived = &archived
	return employee
}

// SetPhone устанавливает Телефон сотрудника.
func (employee *Employee) SetPhone(phone string) *Employee {
	employee.Phone = &phone
	return employee
}

// SetPosition устанавливает Должность сотрудника.
func (employee *Employee) SetPosition(position string) *Employee {
	employee.Position = &position
	return employee
}

// SetSalary устанавливает Оклад сотрудника.
func (employee *Employee) SetSalary(salary float64) *Employee {
	employee.Salary = &Salary{&salary}
	return employee
}

// SetShared устанавливает флаг общего доступа.
func (employee *Employee) SetShared(shared bool) *Employee {
	employee.Shared = &shared
	return employee
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (employee *Employee) SetAttributes(attributes ...*Attribute) *Employee {
	employee.Attributes.Push(attributes...)
	return employee
}

// String реализует интерфейс [fmt.Stringer].
func (employee Employee) String() string {
	return Stringify(employee)
}

// MetaType возвращает код сущности.
func (Employee) MetaType() MetaType {
	return MetaTypeEmployee
}

// Update shortcut
func (employee Employee) Update(ctx context.Context, client *Client, params ...*Params) (*Employee, *resty.Response, error) {
	return NewEmployeeService(client).Update(ctx, employee.GetID(), &employee, params...)
}

// Create shortcut
func (employee Employee) Create(ctx context.Context, client *Client, params ...*Params) (*Employee, *resty.Response, error) {
	return NewEmployeeService(client).Create(ctx, &employee, params...)
}

// Delete shortcut
func (employee Employee) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewEmployeeService(client).Delete(ctx, employee.GetID())
}

// Salary Оклад.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-sotrudniki-atributy-wlozhennyh-suschnostej-oklad
type Salary struct {
	Value *float64 `json:"value,omitempty"` // Сумма оклада
}

// GetValue возвращает Сумму оклада.
func (salary Salary) GetValue() float64 {
	return Deref(salary.Value)
}

// SetValue устанавливает Сумму оклада.
func (salary *Salary) SetValue(value float64) *Salary {
	salary.Value = &value
	return salary
}

// MailActivationRequired структура ответа на запрос активации сотрудника.
//
// Если поле mailActivationRequired равно «true» и сотрудник ранее не был активен, это означает,
// что на указанную у сотрудника почту было выслано письмо со ссылкой на вход для сотрудника.
//
// Если поле mailActivationRequired равно false, то можно использовать ранее заданный пароль для данного пользователя.
type MailActivationRequired struct {
	MailActivationRequired bool `json:"mailActivationRequired"`
}

// EmployeePermission Права Сотрудника.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-rabota-s-prawami-sotrudnika
type EmployeePermission struct {
	AuthorizedIpNetmask *string     `json:"authorizedIpNetmask,omitempty"` // Маска подсети с правом доступа на аккаунт
	AuthorizedIpNetwork *string     `json:"authorizedIpNetwork,omitempty"` // Ipv4 адрес, идентифицирующий соответствующую подсеть, с правом доступа на аккаунт
	Email               *string     `json:"email,omitempty"`               // Почта сотрудника
	Group               *MetaNameID `json:"group,omitempty"`               // Метаданные Группы, а также ее идентификатор и имя
	IsActive            *bool       `json:"isActive,omitempty"`            // Доступ к сервису МойСклад
	Login               *string     `json:"login,omitempty"`               // Логин сотрудника для входа в МойСклад
	Role                *Role       `json:"role,omitempty"`                // Информация о роли Сотрудника
	AuthorizedHosts     []string    `json:"authorizedHosts,omitempty"`     // Список ipv4 адресов, с которых разрешен доступ на аккаунт
}

// GetAuthorizedIpNetmask возвращает Маску подсети с правом доступа на аккаунт.
func (employeePermission EmployeePermission) GetAuthorizedIpNetmask() string {
	return Deref(employeePermission.AuthorizedIpNetmask)
}

// GetAuthorizedIpNetwork возвращает Ipv4 адрес, идентифицирующий соответствующую подсеть, с правом доступа на аккаунт.
func (employeePermission EmployeePermission) GetAuthorizedIpNetwork() string {
	return Deref(employeePermission.AuthorizedIpNetwork)
}

// GetEmail возвращает Почту сотрудника.
func (employeePermission EmployeePermission) GetEmail() string {
	return Deref(employeePermission.Email)
}

// GetGroup возвращает Метаданные Группы, а также ее идентификатор и имя.
func (employeePermission EmployeePermission) GetGroup() MetaNameID {
	return Deref(employeePermission.Group)
}

// GetActive возвращает флаг Доступа к сервису МойСклад.
func (employeePermission EmployeePermission) GetActive() bool {
	return Deref(employeePermission.IsActive)
}

// GetLogin возвращает Логин сотрудника для входа в МойСклад.
func (employeePermission EmployeePermission) GetLogin() string {
	return Deref(employeePermission.Login)
}

// GetRole возвращает Информация о роли сотрудника.
func (employeePermission EmployeePermission) GetRole() Role {
	return Deref(employeePermission.Role)
}

// GetAuthorizedHosts возвращает Список ipv4 адресов, с которых разрешен доступ на аккаунт.
func (employeePermission EmployeePermission) GetAuthorizedHosts() []string {
	return employeePermission.AuthorizedHosts
}

// SetAuthorizedIpNetmask устанавливает Маску подсети с правом доступа на аккаунт.
func (employeePermission *EmployeePermission) SetAuthorizedIpNetmask(authorizedIpNetmask string) *EmployeePermission {
	employeePermission.AuthorizedIpNetmask = &authorizedIpNetmask
	return employeePermission
}

// SetAuthorizedIpNetwork устанавливает Ipv4 адрес, идентифицирующий соответствующую подсеть, с правом доступа на аккаунт.
func (employeePermission *EmployeePermission) SetAuthorizedIpNetwork(authorizedIpNetwork string) *EmployeePermission {
	employeePermission.AuthorizedIpNetwork = &authorizedIpNetwork
	return employeePermission
}

// SetEmail устанавливает Почту сотрудника.
func (employeePermission *EmployeePermission) SetEmail(email string) *EmployeePermission {
	employeePermission.Email = &email
	return employeePermission
}

// SetGroup устанавливает Метаданные Группы.
func (employeePermission *EmployeePermission) SetGroup(group *Group) *EmployeePermission {
	if group != nil {
		employeePermission.Group = &MetaNameID{Meta: group.GetMeta()}
	}
	return employeePermission
}

// SetActive устанавливает Доступ к сервису МойСклад.
func (employeePermission *EmployeePermission) SetActive(isActive bool) *EmployeePermission {
	employeePermission.IsActive = &isActive
	return employeePermission
}

// SetLogin устанавливает Логин сотрудника для входа в МойСклад.
func (employeePermission *EmployeePermission) SetLogin(login string) *EmployeePermission {
	employeePermission.Login = &login
	return employeePermission
}

// SetRole устанавливает Роль сотрудника.
func (employeePermission *EmployeePermission) SetRole(role *Role) *EmployeePermission {
	if role != nil {
		employeePermission.Role = role.Clean()
	}
	return employeePermission
}

// SetAuthorizedHosts устанавливает Список ipv4 адресов, с которых разрешен доступ на аккаунт.
//
// Принимает множество string.
func (employeePermission *EmployeePermission) SetAuthorizedHosts(authorizedHosts ...string) *EmployeePermission {
	employeePermission.AuthorizedHosts = authorizedHosts
	return employeePermission
}

// String реализует интерфейс [fmt.Stringer].
func (employeePermission EmployeePermission) String() string {
	return Stringify(employeePermission)
}

// EmployeeService описывает методы сервиса для работы с сотрудниками.
type EmployeeService interface {
	// GetList выполняет запрос на получение списка сотрудников.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Employee], *resty.Response, error)

	// Create выполняет запрос на создание сотрудника.
	// Обязательные поля для заполнения:
	//	- lastName (Фамилия)
	// Принимает контекст, сотрудника и опционально объект параметров запроса Params.
	// Возвращает созданного сотрудника.
	Create(ctx context.Context, employee *Employee, params ...*Params) (*Employee, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение сотрудников.
	// Изменяемые сотрудники должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список сотрудников и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых сотрудников.
	CreateUpdateMany(ctx context.Context, employeeList Slice[Employee], params ...*Params) (*Slice[Employee], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление сотрудников.
	// Принимает контекст и множество сотрудников.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Employee) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление сотрудника.
	// Принимает контекст и ID сотрудника.
	// Возвращает «true» в случае успешного удаления сотрудника.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных сотрудников.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)

	// GetAttributeList выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект List.
	GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и/или изменение доп полей.
	// Изменяемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или изменённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает «true» в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного сотрудника по ID.
	// Принимает контекст, ID сотрудника и опционально объект параметров запроса Params.
	// Возвращает найденного сотрудника.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Employee, *resty.Response, error)

	// Update выполняет запрос на изменение сотрудника.
	// Принимает контекст, сотрудника и опционально объект параметров запроса Params.
	// Возвращает изменённого сотрудника.
	Update(ctx context.Context, id uuid.UUID, employee *Employee, params ...*Params) (*Employee, *resty.Response, error)

	// GetPermissions выполняет запрос на получение информации о правах сотрудника.
	// Принимает контекст и ID сотрудника.
	// Возвращает объект EmployeePermission.
	GetPermissions(ctx context.Context, id uuid.UUID) (*EmployeePermission, *resty.Response, error)

	// UpdatePermissions выполняет запрос на изменение информации о правах сотрудника.
	// Принимает контекст, ID сотрудника и объект EmployeePermission.
	// Возвращает объект EmployeePermission.
	UpdatePermissions(ctx context.Context, id uuid.UUID, permissions *EmployeePermission) (*EmployeePermission, *resty.Response, error)

	// Activate выполняет запрос на активацию сотрудника.
	//
	// Принимает контекст, ID сотрудника и объект EmployeePermission.
	//
	// Если пользователь ранее не был активным, то при запросе необходимо указать поле login.
	// Успешным результатом выполнения запроса будет json, содержащий поле mailActivationRequired со значением true.
	// Это означает, что на указанную у сотрудника почту было выслано письмо со ссылкой на вход для сотрудника.
	//
	// Если пользователь уже был ранее активным, то при активации не нужно указывать поле login.
	// Успешным результатом выполнения запроса будет json, содержащий поле mailActivationRequired со значением false.
	// В данном случае можно использовать ранее заданный пароль для данного пользователя.
	Activate(ctx context.Context, id uuid.UUID, permissions *EmployeePermission) (bool, *resty.Response, error)

	// Deactivate выполняет запрос на деактивацию сотрудника.
	// Принимает контекст и ID сотрудника.
	// Возвращает «true» в случае успешной деактивации.
	Deactivate(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// ResetPassword выполняет запрос на сброс пароля сотрудника.
	// Принимает контекст и ID сотрудника.
	// Возвращает «true» в случае успешного сброса пароля.
	ResetPassword(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

const (
	EndpointEmployee              = EndpointEntity + string(MetaTypeEmployee)
	EndpointEmployeeSecurity      = EndpointEmployee + "/%s/security"
	EndpointEmployeeActivate      = EndpointEmployee + "/%s/access/activate"
	EndpointEmployeeDeactivate    = EndpointEmployee + "/%s/access/deactivate"
	EndpointEmployeeResetPassword = EndpointEmployee + "/%s/access/resetpassword"
)

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

func (service *employeeService) GetPermissions(ctx context.Context, id uuid.UUID) (*EmployeePermission, *resty.Response, error) {
	path := fmt.Sprintf(EndpointEmployeeSecurity, id)
	return NewRequestBuilder[EmployeePermission](service.client, path).Get(ctx)
}

func (service *employeeService) UpdatePermissions(ctx context.Context, id uuid.UUID, permissions *EmployeePermission) (*EmployeePermission, *resty.Response, error) {
	path := fmt.Sprintf(EndpointEmployeeSecurity, id)
	return NewRequestBuilder[EmployeePermission](service.client, path).Put(ctx, permissions)
}

func (service *employeeService) Activate(ctx context.Context, id uuid.UUID, permissions *EmployeePermission) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointEmployeeActivate, id)
	mailActivationRequired, resp, err := NewRequestBuilder[MailActivationRequired](service.client, path).Put(ctx, permissions)
	return mailActivationRequired.MailActivationRequired, resp, err
}

func (service *employeeService) Deactivate(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointEmployeeDeactivate, id)
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusNoContent, resp, nil
}

func (service *employeeService) ResetPassword(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointEmployeeResetPassword, id)
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusNoContent, resp, nil
}

// NewEmployeeService принимает [Client] и возвращает сервис для работы с сотрудниками.
func NewEmployeeService(client *Client) EmployeeService {
	e := NewEndpoint(client, EndpointEmployee)
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
