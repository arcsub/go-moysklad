package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Organization Юрлицо.
//
// Код сущности: organization
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jurlico
type Organization struct {
	Updated                *Timestamp               `json:"updated,omitempty"`                // Момент последнего обновления Юрлица
	ChiefAccountSign       *Image                   `json:"chiefAccountSign,omitempty"`       // Подпись главного бухгалтера
	ActualAddressFull      *Address                 `json:"actualAddressFull,omitempty"`      // Фактический адрес Юрлица с детализацией по отдельным полям
	Archived               *bool                    `json:"archived,omitempty"`               // Добавлено ли Юрлицо в архив
	BonusPoints            *int                     `json:"bonusPoints,omitempty"`            // Бонусные баллы по активной бонусной программе
	BonusProgram           *NullValue[BonusProgram] `json:"bonusProgram,omitempty"`           // Метаданные активной бонусной программы
	ActualAddress          *string                  `json:"actualAddress,omitempty"`          // Фактический адрес Юрлица
	UTMUrl                 *string                  `json:"utmUrl,omitempty"`                 // IP-адрес УТМ
	Created                *Timestamp               `json:"created,omitempty"`                // Дата создания
	Description            *string                  `json:"description,omitempty"`            // Комментарий к Юрлицу
	ExternalCode           *string                  `json:"externalCode,omitempty"`           // Внешний код Юрлица
	Group                  *Group                   `json:"group,omitempty"`                  // Отдел сотрудника
	ID                     *uuid.UUID               `json:"id,omitempty"`                     // ID Юрлица
	Meta                   *Meta                    `json:"meta,omitempty"`                   // Метаданные Юрлица
	Name                   *string                  `json:"name,omitempty"`                   // Наименование Юрлица
	Owner                  *Employee                `json:"owner,omitempty"`                  // Метаданные владельца (Сотрудника)
	Shared                 *bool                    `json:"shared,omitempty"`                 // Общий доступ
	SyncID                 *uuid.UUID               `json:"syncId,omitempty"`                 // ID синхронизации
	TrackingContractDate   *Timestamp               `json:"trackingContractDate,omitempty"`   // Дата договора с ЦРПТ
	TrackingContractNumber *string                  `json:"trackingContractNumber,omitempty"` // Номер договора с ЦРПТ
	CertificateNumber      *string                  `json:"certificateNumber,omitempty"`      // Номер свидетельства
	Accounts               *MetaArray[AgentAccount] `json:"accounts,omitempty"`               // Метаданные счетов юрлица
	Stamp                  *Image                   `json:"stamp,omitempty"`                  // Печать
	CertificateDate        *Timestamp               `json:"certificateDate,omitempty"`        // Дата свидетельства
	AccountID              *uuid.UUID               `json:"accountId,omitempty"`              // ID учётной записи
	Code                   *string                  `json:"code,omitempty"`                   // Код Юрлица
	ChiefAccountant        *string                  `json:"chiefAccountant,omitempty"`        // Главный бухгалтер
	Director               *string                  `json:"director,omitempty"`               // Руководитель
	DirectorPosition       *string                  `json:"directorPosition,omitempty"`       // Должность руководителя
	DirectorSign           *Image                   `json:"directorSign,omitempty"`           // Подпись руководителя
	Email                  *string                  `json:"email,omitempty"`                  // Адрес электронной почты
	Fax                    *string                  `json:"fax,omitempty"`                    // Номер факса
	FSRARID                *string                  `json:"fsrarId,omitempty"`                // Идентификатор в ФСРАР
	INN                    *string                  `json:"inn,omitempty"`                    // ИНН
	IsEGAISEnable          *bool                    `json:"isEgaisEnable,omitempty"`          // Включен ли ЕГАИС для данного юрлица
	KPP                    *string                  `json:"kpp,omitempty"`                    // КПП
	LegalAddress           *string                  `json:"legalAddress,omitempty"`           // Юридический адреса Юрлица
	LegalAddressFull       *Address                 `json:"legalAddressFull,omitempty"`       // Юридический адрес Юрлица с детализацией по отдельным полям
	LegalFirstName         *string                  `json:"legalFirstName,omitempty"`         // Имя для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalLastName          *string                  `json:"legalLastName,omitempty"`          // Фамилия для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalMiddleName        *string                  `json:"legalMiddleName,omitempty"`        // Отчество для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalTitle             *string                  `json:"legalTitle,omitempty"`             // Полное наименование. Игнорируется, если передано одно из значений для ФИО. Формируется автоматически на основе получаемых ФИО Юрлица
	OGRN                   *string                  `json:"ogrn,omitempty"`                   // ОГРН
	OGRNIP                 *string                  `json:"ogrnip,omitempty"`                 // ОГРНИП
	OKPO                   *string                  `json:"okpo,omitempty"`                   // ОКПО
	PayerVat               *bool                    `json:"payerVat,omitempty"`               // Является ли данное юрлицо плательщиком НДС
	Phone                  *string                  `json:"phone,omitempty"`                  // Номер городского телефона
	CompanyType            CompanyType              `json:"companyType,omitempty"`            // Тип Юрлица . В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	Attributes             Slice[Attribute]         `json:"attributes,omitempty"`             // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (organization Organization) Clean() *Organization {
	if organization.Meta == nil {
		return nil
	}
	return &Organization{Meta: organization.Meta}
}

// AsTaskAgent реализует интерфейс [AgentOrganizationConverter].
func (organization Organization) AsOrganizationAgent() *Agent {
	return organization.AsAgent()
}

// AsAgent реализует интерфейс [AgentConverter].
func (organization Organization) AsAgent() *Agent {
	if organization.Meta == nil {
		return nil
	}
	return &Agent{Meta: organization.Meta}
}

// GetUpdated возвращает Момент последнего обновления Юрлица.
func (organization Organization) GetUpdated() time.Time {
	return Deref(organization.Updated).Time()
}

// GetChiefAccountSign возвращает Подпись главного бухгалтера.
func (organization Organization) GetChiefAccountSign() Image {
	return Deref(organization.ChiefAccountSign)
}

// GetActualAddressFull возвращает Фактический адрес Юрлица с детализацией по отдельным полям.
func (organization Organization) GetActualAddressFull() Address {
	return Deref(organization.ActualAddressFull)
}

// GetArchived возвращает true, если Юрлицо добавлено в архив.
func (organization Organization) GetArchived() bool {
	return Deref(organization.Archived)
}

// GetBonusPoints возвращает Бонусные баллы по активной бонусной программе.
func (organization Organization) GetBonusPoints() int {
	return Deref(organization.BonusPoints)
}

// GetBonusProgram возвращает Метаданные активной бонусной программы.
func (organization Organization) GetBonusProgram() BonusProgram {
	return Deref(organization.BonusProgram).getValue()
}

// GetActualAddress возвращает Фактический адрес Юрлица.
func (organization Organization) GetActualAddress() string {
	return Deref(organization.ActualAddress)
}

// GetUTMUrl возвращает IP-адрес УТ.
func (organization Organization) GetUTMUrl() string {
	return Deref(organization.UTMUrl)
}

// GetCreated возвращает Дату создания.
func (organization Organization) GetCreated() time.Time {
	return Deref(organization.Created).Time()
}

// GetDescription возвращает Комментарий к Юрлицу.
func (organization Organization) GetDescription() string {
	return Deref(organization.Description)
}

// GetExternalCode возвращает Внешний код Юрлица.
func (organization Organization) GetExternalCode() string {
	return Deref(organization.ExternalCode)
}

// GetGroup возвращает Отдел сотрудника.
func (organization Organization) GetGroup() Group {
	return Deref(organization.Group)
}

// GetID возвращает ID Юрлица.
func (organization Organization) GetID() uuid.UUID {
	return Deref(organization.ID)
}

// GetMeta возвращает Метаданные Юрлица.
func (organization Organization) GetMeta() Meta {
	return Deref(organization.Meta)
}

// GetName возвращает Наименование Юрлица.
func (organization Organization) GetName() string {
	return Deref(organization.Name)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (organization Organization) GetOwner() Employee {
	return Deref(organization.Owner)
}

// GetShared возвращает флаг Общего доступа.
func (organization Organization) GetShared() bool {
	return Deref(organization.Shared)
}

// GetSyncID возвращает ID синхронизации.
func (organization Organization) GetSyncID() uuid.UUID {
	return Deref(organization.SyncID)
}

// GetTrackingContractDate возвращает Дату договора с ЦРПТ.
func (organization Organization) GetTrackingContractDate() time.Time {
	return Deref(organization.TrackingContractDate).Time()
}

// GetTrackingContractNumber возвращает Номер договора с ЦРПТ.
func (organization Organization) GetTrackingContractNumber() string {
	return Deref(organization.TrackingContractNumber)
}

// GetCertificateNumber возвращает Номер свидетельства.
func (organization Organization) GetCertificateNumber() string {
	return Deref(organization.CertificateNumber)
}

// GetAccounts возвращает Массив счетов юрлица.
func (organization Organization) GetAccounts() MetaArray[AgentAccount] {
	return Deref(organization.Accounts)
}

// GetStamp возвращает Печать.
func (organization Organization) GetStamp() Image {
	return Deref(organization.Stamp)
}

// GetCertificateDate возвращает Дату свидетельства.
func (organization Organization) GetCertificateDate() time.Time {
	return Deref(organization.CertificateDate).Time()
}

// GetAccountID возвращает ID учётной записи.
func (organization Organization) GetAccountID() uuid.UUID {
	return Deref(organization.AccountID)
}

// GetCode возвращает Код Юрлица.
func (organization Organization) GetCode() string {
	return Deref(organization.Code)
}

// GetChiefAccountant возвращает Главного бухгалтера.
func (organization Organization) GetChiefAccountant() string {
	return Deref(organization.ChiefAccountant)
}

// GetDirector возвращает Руководителя.
func (organization Organization) GetDirector() string {
	return Deref(organization.Director)
}

// GetDirectorPosition возвращает Должность руководителя.
func (organization Organization) GetDirectorPosition() string {
	return Deref(organization.DirectorPosition)
}

// GetDirectorSign возвращает Подпись руководителя.
func (organization Organization) GetDirectorSign() Image {
	return Deref(organization.DirectorSign)
}

// GetEmail возвращает Адрес электронной почты.
func (organization Organization) GetEmail() string {
	return Deref(organization.Email)
}

// GetFax возвращает Номер факса.
func (organization Organization) GetFax() string {
	return Deref(organization.Fax)
}

// GetFSRARID возвращает Идентификатор в ФСРАР.
func (organization Organization) GetFSRARID() string {
	return Deref(organization.FSRARID)
}

// GetINN возвращает ИНН.
func (organization Organization) GetINN() string {
	return Deref(organization.INN)
}

// GetIsEGAISEnable возвращает true, если для данного юрлица включен ЕГАИС.
func (organization Organization) GetIsEGAISEnable() bool {
	return Deref(organization.IsEGAISEnable)
}

// GetKPP возвращает КПП.
func (organization Organization) GetKPP() string {
	return Deref(organization.KPP)
}

// GetLegalAddress возвращает Юридический адрес Юрлица.
func (organization Organization) GetLegalAddress() string {
	return Deref(organization.LegalAddress)
}

// GetLegalAddressFull возвращает Юридический адрес Юрлица с детализацией по отдельным полям.
func (organization Organization) GetLegalAddressFull() Address {
	return Deref(organization.LegalAddressFull)
}

// GetLegalFirstName возвращает Имя для Юрлица типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (organization Organization) GetLegalFirstName() string {
	return Deref(organization.LegalFirstName)
}

// GetLegalLastName возвращает Фамилию для Юрлица типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (organization Organization) GetLegalLastName() string {
	return Deref(organization.LegalLastName)
}

// GetLegalMiddleName возвращает Отчество для Юрлица типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (organization Organization) GetLegalMiddleName() string {
	return Deref(organization.LegalMiddleName)
}

// GetLegalTitle возвращает Полное наименование.
//
// Игнорируется для Контрагентов типа [Индивидуальный предприниматель, Физическое лицо],
// если передано одно из значений для ФИО и формируется автоматически на основе получаемых ФИО Юрлица.
func (organization Organization) GetLegalTitle() string {
	return Deref(organization.LegalTitle)
}

// GetOGRN возвращает ОГРН.
func (organization Organization) GetOGRN() string {
	return Deref(organization.OGRN)
}

// GetOGRNIP возвращает ОГРНИП.
func (organization Organization) GetOGRNIP() string {
	return Deref(organization.OGRNIP)
}

// GetOKPO возвращает ОКПО.
func (organization Organization) GetOKPO() string {
	return Deref(organization.OKPO)
}

// GetPayerVat возвращает true, если данное юрлицо является плательщиком НДС.
func (organization Organization) GetPayerVat() bool {
	return Deref(organization.PayerVat)
}

// GetPhone возвращает Номер городского телефона.
func (organization Organization) GetPhone() string {
	return Deref(organization.Phone)
}

// GetCompanyType возвращает Тип Юрлица.
//
// В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться.
func (organization Organization) GetCompanyType() CompanyType {
	return organization.CompanyType
}

// GetAttributes возвращает Список метаданных доп. полей.
func (organization Organization) GetAttributes() Slice[Attribute] {
	return organization.Attributes
}

// SetChiefAccountSign устанавливает Подпись главного бухгалтера.
func (organization *Organization) SetChiefAccountSign(chiefAccountSign *Image) *Organization {
	if chiefAccountSign != nil {
		organization.ChiefAccountSign = chiefAccountSign
	}
	return organization
}

// SetActualAddressFull устанавливает Фактический адрес Контрагента с детализацией по отдельным полям.
//
// Передача nil передаёт сброс значения (null).
func (organization *Organization) SetActualAddressFull(actualAddressFull *Address) *Organization {
	if actualAddressFull == nil {
		organization.SetActualAddress("")
	} else {
		organization.ActualAddressFull = actualAddressFull
	}
	return organization
}

// SetArchived устанавливает флаг юрлица комплекта в архиве.
func (organization *Organization) SetArchived(archived bool) *Organization {
	organization.Archived = &archived
	return organization
}

// SetBonusProgram устанавливает Метаданные бонусной программы.
//
// Передача nil передаёт сброс значения (null).
func (organization *Organization) SetBonusProgram(bonusProgram *BonusProgram) *Organization {
	organization.BonusProgram = NewNullValue(bonusProgram)
	return organization
}

// SetActualAddress устанавливает Фактический адрес Юрлица.
func (organization *Organization) SetActualAddress(actualAddress string) *Organization {
	organization.ActualAddress = &actualAddress
	return organization
}

// SetUTMUrl устанавливает IP-адрес УТМ.
func (organization *Organization) SetUTMUrl(utmUrl string) *Organization {
	organization.UTMUrl = &utmUrl
	return organization
}

// SetDescription устанавливает Комментарий к Юрлицу.
func (organization *Organization) SetDescription(description string) *Organization {
	organization.Description = &description
	return organization
}

// SetExternalCode устанавливает Внешний код Юрлица.
func (organization *Organization) SetExternalCode(externalCode string) *Organization {
	organization.ExternalCode = &externalCode
	return organization
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (organization *Organization) SetGroup(group *Group) *Organization {
	if group != nil {
		organization.Group = group.Clean()
	}
	return organization
}

// SetMeta устанавливает Метаданные Юрлица.
func (organization *Organization) SetMeta(meta *Meta) *Organization {
	organization.Meta = meta
	return organization
}

// SetName устанавливает Наименование Юрлица.
func (organization *Organization) SetName(name string) *Organization {
	organization.Name = &name
	return organization
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (organization *Organization) SetOwner(owner *Employee) *Organization {
	if owner != nil {
		organization.Owner = owner.Clean()
	}
	return organization
}

// SetShared устанавливает флаг общего доступа.
func (organization *Organization) SetShared(shared bool) *Organization {
	organization.Shared = &shared
	return organization
}

// SetSyncID устанавливает ID синхронизации.
func (organization *Organization) SetSyncID(syncID uuid.UUID) *Organization {
	organization.SyncID = &syncID
	return organization
}

// SetTrackingContractDate устанавливает Дату договора с ЦРПТ.
func (organization *Organization) SetTrackingContractDate(trackingContractDate time.Time) *Organization {
	organization.TrackingContractDate = NewTimestamp(trackingContractDate)
	return organization
}

// SetTrackingContractNumber устанавливает Номер договора с ЦРПТ.
func (organization *Organization) SetTrackingContractNumber(trackingContractNumber string) *Organization {
	organization.TrackingContractNumber = &trackingContractNumber
	return organization
}

// SetCertificateNumber устанавливает Номер свидетельства.
func (organization *Organization) SetCertificateNumber(certificateNumber string) *Organization {
	organization.CertificateNumber = &certificateNumber
	return organization
}

// SetAccounts устанавливает Массив счетов Контрагентов.
//
// Принимает множества объектов [AgentAccount].
func (organization *Organization) SetAccounts(accounts ...*AgentAccount) *Organization {
	organization.Accounts = NewMetaArrayFrom(accounts)
	return organization
}

// SetStamp устанавливает Печать.
func (organization *Organization) SetStamp(stamp *Image) *Organization {
	if stamp != nil {
		organization.Stamp = stamp
	}
	return organization
}

// SetCertificateDate устанавливает Дату свидетельства.
func (organization *Organization) SetCertificateDate(certificateDate time.Time) *Organization {
	organization.CertificateDate = NewTimestamp(certificateDate)
	return organization
}

// SetCode устанавливает Код Юрлица.
func (organization *Organization) SetCode(code string) *Organization {
	organization.Code = &code
	return organization
}

// SetChiefAccountant устанавливает Главного бухгалтера.
func (organization *Organization) SetChiefAccountant(chiefAccountant string) *Organization {
	organization.ChiefAccountant = &chiefAccountant
	return organization
}

// SetDirector устанавливает Руководитель.
func (organization *Organization) SetDirector(director string) *Organization {
	organization.Director = &director
	return organization
}

// SetDirectorPosition устанавливает Должность руководителя.
func (organization *Organization) SetDirectorPosition(directorPosition string) *Organization {
	organization.DirectorPosition = &directorPosition
	return organization
}

// SetDirectorSign устанавливает Подпись руководителя.
func (organization *Organization) SetDirectorSign(directorSign *Image) *Organization {
	if directorSign != nil {
		organization.DirectorSign = directorSign
	}
	return organization
}

// SetEmail устанавливает Адрес электронной почты.
func (organization *Organization) SetEmail(email string) *Organization {
	organization.Email = &email
	return organization
}

// SetFax устанавливает Номер факса.
func (organization *Organization) SetFax(fax string) *Organization {
	organization.Fax = &fax
	return organization
}

// SetFSRARID устанавливает Идентификатор в ФСРАР.
func (organization *Organization) SetFSRARID(fsrarID string) *Organization {
	organization.FSRARID = &fsrarID
	return organization
}

// SetINN устанавливает ИНН.
func (organization *Organization) SetINN(inn string) *Organization {
	organization.INN = &inn
	return organization
}

// SetEGAISEnable устанавливает признак включения ЕГАИС для данного юрлица.
func (organization *Organization) SetEGAISEnable(isEGAISEnable bool) *Organization {
	organization.IsEGAISEnable = &isEGAISEnable
	return organization
}

// SetKPP устанавливает КПП.
func (organization *Organization) SetKPP(kpp string) *Organization {
	organization.KPP = &kpp
	return organization
}

// SetLegalAddress устанавливает Юридический адрес Юрлица.
func (organization *Organization) SetLegalAddress(legalAddress string) *Organization {
	organization.LegalAddress = &legalAddress
	return organization
}

// SetLegalAddressFull устанавливает Юридический адрес Юрлица с детализацией по отдельным полям.
//
// Передача nil передаёт сброс значения (null).
func (organization *Organization) SetLegalAddressFull(legalAddressFull *Address) *Organization {
	if legalAddressFull == nil {
		organization.SetLegalAddress("")
	} else {
		organization.LegalAddressFull = legalAddressFull
	}
	return organization
}

// SetLegalFirstName устанавливает Имя для Юрлица типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (organization *Organization) SetLegalFirstName(legalFirstName string) *Organization {
	organization.LegalFirstName = &legalFirstName
	return organization
}

// SetLegalLastName устанавливает Фамилию для Юрлица типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (organization *Organization) SetLegalLastName(legalLastName string) *Organization {
	organization.LegalLastName = &legalLastName
	return organization
}

// SetLegalMiddleName устанавливает Отчество для Юрлица типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (organization *Organization) SetLegalMiddleName(legalMiddleName string) *Organization {
	organization.LegalMiddleName = &legalMiddleName
	return organization
}

// SetOGRN устанавливает ОГРН.
func (organization *Organization) SetOGRN(ogrn string) *Organization {
	organization.OGRN = &ogrn
	return organization
}

// SetOGRNIP устанавливает ОГРНИП.
func (organization *Organization) SetOGRNIP(ogrnip string) *Organization {
	organization.OGRNIP = &ogrnip
	return organization
}

// SetOKPO устанавливает ОКПО.
func (organization *Organization) SetOKPO(okpo string) *Organization {
	organization.OKPO = &okpo
	return organization
}

// SetPayerVat устанавливает признак плательщика НДС.
func (organization *Organization) SetPayerVat(payerVat bool) *Organization {
	organization.PayerVat = &payerVat
	return organization
}

// SetPhone устанавливает Номер городского телефона.
func (organization *Organization) SetPhone(phone string) *Organization {
	organization.Phone = &phone
	return organization
}

// SetCompanyType устанавливает Тип Контрагента.
func (organization *Organization) SetCompanyType(companyType CompanyType) *Organization {
	organization.CompanyType = companyType
	return organization
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (organization *Organization) SetAttributes(attributes ...*Attribute) *Organization {
	organization.Attributes.Push(attributes...)
	return organization
}

// String реализует интерфейс [fmt.Stringer].
func (organization Organization) String() string {
	return Stringify(organization)
}

// MetaType возвращает код сущности.
func (Organization) MetaType() MetaType {
	return MetaTypeOrganization
}

// Update shortcut
func (organization *Organization) Update(ctx context.Context, client *Client, params ...*Params) (*Organization, *resty.Response, error) {
	return NewOrganizationService(client).Update(ctx, organization.GetID(), organization, params...)
}

// Create shortcut
func (organization *Organization) Create(ctx context.Context, client *Client, params ...*Params) (*Organization, *resty.Response, error) {
	return NewOrganizationService(client).Create(ctx, organization, params...)
}

// Delete shortcut
func (organization *Organization) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewOrganizationService(client).Delete(ctx, organization)
}

// OrganizationService описывает методы сервиса для работы с юридическими лицами.
type OrganizationService interface {
	// GetList выполняет запрос на получение списка юрлиц.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Organization], *resty.Response, error)

	// Create выполняет запрос на создание юрлица.
	// Обязательные поля для заполнения:
	//	- name (Наименование Юрлица)
	// Принимает контекст, юрлицо и опционально объект параметров запроса Params.
	// Возвращает созданное юрлицо.
	Create(ctx context.Context, organization *Organization, params ...*Params) (*Organization, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение юрлиц.
	// Изменяемые юрлица должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список юрлиц и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых юрлиц.
	CreateUpdateMany(ctx context.Context, organizationList Slice[Organization], params ...*Params) (*Slice[Organization], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление юрлиц.
	// Принимает контекст и множество юрлиц.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Organization) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление юрлица по ID.
	// Принимает контекст и ID юрлица.
	// Возвращает «true» в случае успешного удаления юрлица.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление юрлица.
	// Принимает контекст и юрлицо.
	// Возвращает «true» в случае успешного удаления юрлица.
	Delete(ctx context.Context, entity *Organization) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного юрлица по ID.
	// Принимает контекст, ID юрлица и опционально объект параметров запроса Params.
	// Возвращает найденное юрлицо.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Organization, *resty.Response, error)

	// Update выполняет запрос на изменение юрлица.
	// Принимает контекст, юрлицо и опционально объект параметров запроса Params.
	// Возвращает изменённое юрлицо.
	Update(ctx context.Context, id uuid.UUID, organization *Organization, params ...*Params) (*Organization, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных юрлиц.
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

	// GetAccountList выполняет запрос на получение списка счетов юрлица.
	// Принимает контекст и ID юрлица.
	// Возвращает объект List.
	GetAccountList(ctx context.Context, id uuid.UUID) (*List[AgentAccount], *resty.Response, error)

	// GetAccountByID выполняет запрос на получение отдельного счёта юрлица по ID.
	// Принимает контекст, ID юрлица и ID счёта юрлица.
	// Возвращает найденный счёт юрлица.
	GetAccountByID(ctx context.Context, id uuid.UUID, accountID uuid.UUID) (*AgentAccount, *resty.Response, error)

	// UpdateAccountMany выполняет запрос на массовое изменение счетов юрлица.
	// Принимает контекст, ID юрлица и множество счетов юрлица.
	// Возвращает список изменённых счетов юрлица.
	UpdateAccountMany(ctx context.Context, id uuid.UUID, accounts ...*AgentAccount) (*MetaArray[AgentAccount], *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Organization, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
}

const (
	EndpointOrganization = EndpointEntity + string(MetaTypeOrganization)
)

// NewOrganizationService принимает [Client] и возвращает сервис для работы с юридическими лицами.
func NewOrganizationService(client *Client) OrganizationService {
	return newMainService[Organization, any, MetaAttributesSharedWrapper, any](client, EndpointOrganization)
}
