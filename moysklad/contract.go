package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Contract Договор.
//
// Код сущности: contract
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-dogowor
type Contract struct {
	AgentAccount        *AgentAccount     `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Published           *bool             `json:"published,omitempty"`           // Опубликован ли документ
	RewardPercent       *float64          `json:"rewardPercent,omitempty"`       // Вознаграждение в процентах (от 0 до 100)
	Archived            *bool             `json:"archived,omitempty"`            // Добавлен ли Договор в архив
	Agent               *Counterparty     `json:"agent,omitempty"`               // Метаданные Контрагента
	Code                *string           `json:"code,omitempty"`                // Код Договора
	Name                *string           `json:"name,omitempty"`                // Номер договора
	Description         *string           `json:"description,omitempty"`         // Описание Договора
	ExternalCode        *string           `json:"externalCode,omitempty"`        // Внешний код Договора
	Group               *Group            `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID        `json:"id,omitempty"`                  // ID Договора
	Meta                *Meta             `json:"meta,omitempty"`                // Метаданные Договора
	Moment              *Timestamp        `json:"moment,omitempty"`              // Дата Договора
	Printed             *bool             `json:"printed,omitempty"`             // Напечатан ли документ
	OrganizationAccount *AgentAccount     `json:"organizationAccount,omitempty"` // Метаданные счета вашего юрлица
	OwnAgent            *Organization     `json:"ownAgent,omitempty"`            // Метаданные вашего юрлица
	Owner               *Employee         `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Rate                *NullValue[Rate]  `json:"rate,omitempty"`                // Валюта
	AccountID           *uuid.UUID        `json:"accountId,omitempty"`           // ID учётной записи
	Updated             *Timestamp        `json:"updated,omitempty"`             // Момент последнего обновления сущности
	Shared              *bool             `json:"shared,omitempty"`              // Общий доступ
	State               *NullValue[State] `json:"state,omitempty"`               // Метаданные статуса договора
	Sum                 *float64          `json:"sum,omitempty"`                 // Сумма Договора
	SyncID              *uuid.UUID        `json:"syncId,omitempty"`              // ID синхронизации
	ContractType        ContractType      `json:"contractType,omitempty"`        // Тип Договора
	RewardType          RewardType        `json:"rewardType,omitempty"`          // Тип Вознаграждения
	Attributes          Slice[Attribute]  `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (contract Contract) Clean() *Contract {
	if contract.Meta == nil {
		return nil
	}
	return &Contract{Meta: contract.Meta}
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (contract Contract) GetAgentAccount() AgentAccount {
	return Deref(contract.AgentAccount)
}

// GetPublished возвращает true, если документ опубликован.
func (contract Contract) GetPublished() bool {
	return Deref(contract.Published)
}

// GetRewardPercent возвращает Вознаграждение в процентах (от 0 до 100).
func (contract Contract) GetRewardPercent() float64 {
	return Deref(contract.RewardPercent)
}

// GetArchived возвращает true, если Договор добавлен в архив.
func (contract Contract) GetArchived() bool {
	return Deref(contract.Archived)
}

// GetAgent возвращает Метаданные контрагента.
func (contract Contract) GetAgent() Counterparty {
	return Deref(contract.Agent)
}

// GetCode возвращает Код Договора.
func (contract Contract) GetCode() string {
	return Deref(contract.Code)
}

// GetName возвращает Номер Договора.
func (contract Contract) GetName() string {
	return Deref(contract.Name)
}

// GetDescription возвращает Описание Договора.
func (contract Contract) GetDescription() string {
	return Deref(contract.Description)
}

// GetExternalCode возвращает Внешний код Договора.
func (contract Contract) GetExternalCode() string {
	return Deref(contract.ExternalCode)
}

// GetGroup возвращает Отдел сотрудника.
func (contract Contract) GetGroup() Group {
	return Deref(contract.Group)
}

// GetID возвращает ID Договора.
func (contract Contract) GetID() uuid.UUID {
	return Deref(contract.ID)
}

// GetMeta возвращает Метаданные Договора.
func (contract Contract) GetMeta() Meta {
	return Deref(contract.Meta)
}

// GetMoment возвращает Дату Договора.
func (contract Contract) GetMoment() Timestamp {
	return Deref(contract.Moment)
}

// GetPrinted возвращает true, если Договор напечатан.
func (contract Contract) GetPrinted() bool {
	return Deref(contract.Printed)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (contract Contract) GetOrganizationAccount() AgentAccount {
	return Deref(contract.OrganizationAccount)
}

// GetOwnAgent возвращает Метаданные вашего юрлица.
func (contract Contract) GetOwnAgent() Organization {
	return Deref(contract.OwnAgent)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (contract Contract) GetOwner() Employee {
	return Deref(contract.Owner)
}

// GetRate возвращает Валюту.
func (contract Contract) GetRate() Rate {
	return Deref(contract.Rate).GetValue()
}

// GetAccountID возвращает ID учётной записи.
func (contract Contract) GetAccountID() uuid.UUID {
	return Deref(contract.AccountID)
}

// GetUpdated возвращает Момент последнего обновления.
func (contract Contract) GetUpdated() Timestamp {
	return Deref(contract.Updated)
}

// GetShared возвращает флаг Общего доступа.
func (contract Contract) GetShared() bool {
	return Deref(contract.Shared)
}

// GetState возвращает Метаданные статуса Договора.
func (contract Contract) GetState() State {
	return Deref(contract.State).GetValue()
}

// GetSum возвращает Сумму Договора.
func (contract Contract) GetSum() float64 {
	return Deref(contract.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (contract Contract) GetSyncID() uuid.UUID {
	return Deref(contract.SyncID)
}

// GetContractType возвращает Тип Договора.
func (contract Contract) GetContractType() ContractType {
	return contract.ContractType
}

// GetRewardType возвращает Тип Вознаграждения
func (contract Contract) GetRewardType() RewardType {
	return contract.RewardType
}

// GetAttributes возвращает Список метаданных доп. полей.
func (contract Contract) GetAttributes() Slice[Attribute] {
	return contract.Attributes
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (contract *Contract) SetAgentAccount(agentAccount *AgentAccount) *Contract {
	if agentAccount != nil {
		contract.AgentAccount = agentAccount.Clean()
	}
	return contract
}

// SetRewardPercent устанавливает Вознаграждение в процентах (от 0 до 100).
func (contract *Contract) SetRewardPercent(rewardPercent float64) *Contract {
	contract.RewardPercent = &rewardPercent
	return contract
}

// SetArchived устанавливает флаг нахождения договора в архиве.
func (contract *Contract) SetArchived(archived bool) *Contract {
	contract.Archived = &archived
	return contract
}

// SetAgent устанавливает Метаданные контрагента.
func (contract *Contract) SetAgent(agent *Counterparty) *Contract {
	if agent != nil {
		contract.Agent = agent.Clean()
	}
	return contract
}

// SetCode устанавливает Код Договора.
func (contract *Contract) SetCode(code string) *Contract {
	contract.Code = &code
	return contract
}

// SetName устанавливает Номер Договора.
func (contract *Contract) SetName(name string) *Contract {
	contract.Name = &name
	return contract
}

// SetDescription устанавливает Описание Договора.
func (contract *Contract) SetDescription(description string) *Contract {
	contract.Description = &description
	return contract
}

// SetExternalCode устанавливает Внешний код Договора.
func (contract *Contract) SetExternalCode(externalCode string) *Contract {
	contract.ExternalCode = &externalCode
	return contract
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (contract *Contract) SetGroup(group *Group) *Contract {
	if group != nil {
		contract.Group = group.Clean()
	}
	return contract
}

// SetMeta устанавливает Метаданные Договора.
func (contract *Contract) SetMeta(meta *Meta) *Contract {
	contract.Meta = meta
	return contract
}

// SetMoment устанавливает Дату Договора.
func (contract *Contract) SetMoment(moment time.Time) *Contract {
	contract.Moment = NewTimestamp(moment)
	return contract
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (contract *Contract) SetOrganizationAccount(organizationAccount *AgentAccount) *Contract {
	if organizationAccount != nil {
		contract.OrganizationAccount = organizationAccount.Clean()
	}
	return contract
}

// SetOwnAgent устанавливает Метаданные вашего юрлица.
func (contract *Contract) SetOwnAgent(ownAgent *Organization) *Contract {
	if ownAgent != nil {
		contract.OwnAgent = ownAgent.Clean()
	}
	return contract
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (contract *Contract) SetOwner(owner *Employee) *Contract {
	if owner != nil {
		contract.Owner = owner.Clean()
	}
	return contract
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (contract *Contract) SetRate(rate *Rate) *Contract {
	contract.Rate = NewNullValue(rate)
	return contract
}

// SetShared устанавливает флаг общего доступа.
func (contract *Contract) SetShared(shared bool) *Contract {
	contract.Shared = &shared
	return contract
}

// SetState устанавливает Метаданные статуса Договора.
//
// Передача nil передаёт сброс значения (null).
func (contract *Contract) SetState(state *State) *Contract {
	contract.State = NewNullValue(state)
	return contract
}

// SetSum устанавливает Сумму Договора.
func (contract *Contract) SetSum(sum *float64) *Contract {
	contract.Sum = sum
	return contract
}

// SetSyncID устанавливает ID синхронизации.
func (contract *Contract) SetSyncID(syncID uuid.UUID) *Contract {
	contract.SyncID = &syncID
	return contract
}

// SetContractType устанавливает Тип Договора.
func (contract *Contract) SetContractType(contractType ContractType) *Contract {
	contract.ContractType = contractType
	return contract
}

// SetContractTypeCommission устанавливает Тип Договора - Договор комиссии.
func (contract *Contract) SetContractTypeCommission() *Contract {
	contract.ContractType = ContractTypeCommission
	return contract
}

// SetContractTypeSales устанавливает Тип Договора - Договор купли-продажи.
func (contract *Contract) SetContractTypeSales() *Contract {
	contract.ContractType = ContractTypeSales
	return contract
}

// SetRewardType устанавливает Тип Вознаграждения.
func (contract *Contract) SetRewardType(rewardType RewardType) *Contract {
	contract.RewardType = rewardType
	return contract
}

// SetRewardTypePercentOfSales устанавливает Тип Вознаграждения - Процент от суммы продажи.
func (contract *Contract) SetRewardTypePercentOfSales() *Contract {
	contract.RewardType = RewardTypePercentOfSales
	return contract
}

// SetRewardTypeNone устанавливает Тип Вознаграждения - Не рассчитывать.
func (contract *Contract) SetRewardTypeNone() *Contract {
	contract.RewardType = RewardTypeNone
	return contract
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (contract *Contract) SetAttributes(attributes ...*Attribute) *Contract {
	contract.Attributes.Push(attributes...)
	return contract
}

// String реализует интерфейс [fmt.Stringer].
func (contract Contract) String() string {
	return Stringify(contract)
}

// MetaType возвращает код сущности.
func (Contract) MetaType() MetaType {
	return MetaTypeContract
}

// ContractType Тип Договора.
//
// Возможные значения:
//   - ContractTypeCommission – Договор комиссии
//   - ContractTypeSales      – Договор купли-продажи
type ContractType string

const (
	ContractTypeCommission ContractType = "Commission" // Договор комиссии
	ContractTypeSales      ContractType = "Sales"      // Договор купли-продажи
)

// RewardType Тип Вознаграждения.
//
// Возможные значения:
//   - RewardTypePercentOfSales – Процент от суммы продажи
//   - RewardTypeNone           – Не рассчитывать
type RewardType string

const (
	RewardTypePercentOfSales RewardType = "PercentOfSales" // Процент от суммы продажи
	RewardTypeNone           RewardType = "None"           // Не рассчитывать
)

// ContractService описывает методы сервиса для работы с договорами.
type ContractService interface {
	// GetList выполняет запрос на получение списка договоров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Contract], *resty.Response, error)

	// Create выполняет запрос на создание договора.
	// Обязательные поля для заполнения:
	//	- name (Номер договора)
	//	- ownAgent (Метаданные вашего юрлица)
	//	- agent (Метаданные Контрагента)
	// Принимает контекст, договор и опционально объект параметров запроса Params.
	// Возвращает созданный договор.
	Create(ctx context.Context, contract *Contract, params ...*Params) (*Contract, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение договоров.
	// Изменяемые договоры должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список договоров и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых договоров.
	CreateUpdateMany(ctx context.Context, contractList Slice[Contract], params ...*Params) (*Slice[Contract], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление договоров.
	// Принимает контекст и множество договоров.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Contract) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление договора.
	// Принимает контекст и ID договора.
	// Возвращает true в случае успешного удаления договора.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных договоров.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного договора по ID.
	// Принимает контекст, ID договора и опционально объект параметров запроса Params.
	// Возвращает найденный договор.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Contract, *resty.Response, error)

	// Update выполняет запрос на изменение договора.
	// Принимает контекст, договор и опционально объект параметров запроса Params.
	// Возвращает изменённый договор.
	Update(ctx context.Context, id uuid.UUID, contract *Contract, params ...*Params) (*Contract, *resty.Response, error)

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
	// Возвращает true в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetPublicationList выполняет запрос на получение списка публикаций.
	// Принимает контекст и ID документа.
	// Возвращает объект List.
	GetPublicationList(ctx context.Context, id uuid.UUID) (*List[Publication], *resty.Response, error)

	// GetPublicationByID выполняет запрос на получение отдельной публикации по ID.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает найденную публикацию.
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)

	// Publish выполняет запрос на создание публикации.
	// Принимает контекст, ID документа и шаблон.
	// Возвращает созданную публикацию.
	Publish(ctx context.Context, id uuid.UUID, template TemplateInterface) (*Publication, *resty.Response, error)

	// DeletePublication выполняет запрос на удаление публикации.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает true в случае успешного удаления публикации.
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает true в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetStateByID выполняет запрос на получение статуса документа по ID.
	// Принимает контекст и ID статуса.
	// Возвращает найденный статус.
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)

	// CreateState выполняет запрос на создание статуса документа.
	// Принимает контекст и статус.
	// Возвращает созданный статус.
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)

	// UpdateState выполняет запрос на изменение статуса документа.
	// Принимает контекст, ID статуса и статус.
	// Возвращает изменённый статус.
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)

	// CreateUpdateStateMany выполняет запрос на массовое создание и/или изменение статусов документа.
	// Принимает контекст и множество статусов.
	// Возвращает список созданных и/или изменённых статусов.
	CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error)

	// DeleteState выполняет запрос на удаление статуса документа.
	// Принимает контекст и ID статуса.
	// Возвращает true в случае успешного удаления статуса.
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

// NewContractService принимает [Client] и возвращает сервис для работы с договорами.
func NewContractService(client *Client) ContractService {
	return newMainService[Contract, any, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/contract"))
}
