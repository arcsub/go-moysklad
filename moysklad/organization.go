package moysklad

import (
	"github.com/google/uuid"
)

// Organization Юрлицо.
// Ключевое слово: organization
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jurlico
type Organization struct {
	AccountID              *uuid.UUID    `json:"accountId,omitempty"`              // ID учетной записи
	ActualAddress          *string       `json:"actualAddress,omitempty"`          // Фактический адрес Юрлица
	ActualAddressFull      *Address      `json:"actualAddressFull,omitempty"`      // Фактический адрес Юрлица с детализацией по отдельным полям
	Archived               *bool         `json:"archived,omitempty"`               // Добавлено ли Юрлицо в архив
	BonusPoints            *int          `json:"bonusPoints,omitempty"`            // Бонусные баллы по активной бонусной программе
	BonusProgram           *BonusProgram `json:"bonusProgram,omitempty"`           // Метаданные активной бонусной программы
	Code                   *string       `json:"code,omitempty"`                   // Код Юрлица
	CompanyType            CompanyType   `json:"companyType,omitempty"`            // Тип Юрлица . В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	Created                *Timestamp    `json:"created,omitempty"`                // Дата создания
	Description            *string       `json:"description,omitempty"`            // Комментарий к Юрлицу
	ExternalCode           *string       `json:"externalCode,omitempty"`           // Внешний код Юрлица
	Group                  *Group        `json:"group,omitempty"`                  // Отдел сотрудника
	ID                     *uuid.UUID    `json:"id,omitempty"`                     // ID сущности
	Meta                   *Meta         `json:"meta,omitempty"`                   // Метаданные
	Name                   *string       `json:"name,omitempty"`                   // Наименование
	Owner                  *Employee     `json:"owner,omitempty"`                  // Владелец (Сотрудник)
	Shared                 *bool         `json:"shared,omitempty"`                 // Общий доступ
	SyncID                 *uuid.UUID    `json:"syncId,omitempty"`                 // ID синхронизации
	TrackingContractDate   *Timestamp    `json:"trackingContractDate,omitempty"`   // Дата договора с ЦРПТ
	TrackingContractNumber *string       `json:"trackingContractNumber,omitempty"` // Номер договора с ЦРПТ
	Updated                *Timestamp    `json:"updated,omitempty"`                // Момент последнего обновления Юрлица
	Accounts               *MetaWrapper  `json:"accounts,omitempty"`               // Метаданные счетов юрлица
	Attributes             *Attributes   `json:"attributes,omitempty"`             // Массив метаданных дополнительных полей юрлица
	CertificateDate        *Timestamp    `json:"certificateDate,omitempty"`        // Дата свидетельства
	CertificateNumber      *string       `json:"certificateNumber,omitempty"`      // Номер свидетельства
	ChiefAccountSign       *Image        `json:"chiefAccountSign,omitempty"`       // Подпись главного бухгалтера
	ChiefAccountant        *string       `json:"chiefAccountant,omitempty"`        // Главный бухгалтер
	Director               *string       `json:"director,omitempty"`               // Руководитель
	DirectorPosition       *string       `json:"directorPosition,omitempty"`       // Должность руководителя
	DirectorSign           *Image        `json:"directorSign,omitempty"`           // Подпись руководителя
	Email                  *string       `json:"email,omitempty"`                  // Электронная почта
	Fax                    *string       `json:"fax,omitempty"`                    // Факс
	FSRARId                *string       `json:"fsrarId,omitempty"`                // Идентификатор в ФСРАР
	INN                    *string       `json:"inn,omitempty"`                    // ИНН
	IsEGAISEnable          *bool         `json:"isEgaisEnable,omitempty"`          // Включен ли ЕГАИС для данного юрлица
	KPP                    *string       `json:"kpp,omitempty"`                    // КПП
	LegalAddress           *string       `json:"legalAddress,omitempty"`           // Юридический адрес юрлица
	LegalAddressFull       *Address      `json:"legalAddressFull,omitempty"`       // Структурированный Юридический адрес юрлица
	LegalFirstName         *string       `json:"legalFirstName,omitempty"`         // Имя для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalLastName          *string       `json:"legalLastName,omitempty"`          // Фамилия для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalMiddleName        *string       `json:"legalMiddleName,omitempty"`        // Отчество для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalTitle             *string       `json:"legalTitle,omitempty"`             // Полное наименование. Игнорируется, если передано одно из значений для ФИО. Формируется автоматически на основе получаемых ФИО Юрлица
	OGRN                   *string       `json:"ogrn,omitempty"`                   // ОГРН
	OGRNIP                 *string       `json:"ogrnip,omitempty"`                 // ОГРНИП
	OKPO                   *string       `json:"okpo,omitempty"`                   // ОКПО
	PayerVat               *bool         `json:"payerVat,omitempty"`               // Является ли данное юрлицо плательщиком НДС
	Phone                  *string       `json:"phone,omitempty"`                  // Телефон
	Stamp                  *Image        `json:"stamp,omitempty"`                  // Печать
	UTMUrl                 *string       `json:"utmUrl,omitempty"`                 // IP-адрес УТМ
}

func (o Organization) String() string {
	return Stringify(o)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (o Organization) GetMeta() *Meta {
	return o.Meta
}

func (o Organization) MetaType() MetaType {
	return MetaTypeOrganization
}
