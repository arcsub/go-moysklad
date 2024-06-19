package moysklad

// TrackingCode Коды маркировки
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-kod-markirowki-atributy-suschnosti
type TrackingCode struct {
	ID               *string             `json:"id,omitempty"`                 // ID кода маркировки
	Cis              *string             `json:"cis,omitempty"`                // Код маркировки в стандартном формате
	Cis1162          *string             `json:"cis_1162,omitempty"`           // Код маркировки в формате тега 1162
	Type             TrackingCodeType    `json:"type,omitempty"`               // Тип кода маркировки
	TrackingCodes    Slice[TrackingCode] `json:"trackingCodes,omitempty"`      // Массив вложенных кодов маркировки. Может присутствовать, только если type имеет значения consumerpack или transportpack
	TrackingCode1162 Slice[TrackingCode] `json:"trackingCodes_1162,omitempty"` // Массив вложенных кодов маркировки. Может присутствовать, только если type имеет значения consumerpack или transportpack
}

func (trackingCode TrackingCode) GetID() string {
	return Deref(trackingCode.ID)
}

func (trackingCode TrackingCode) GetCis() string {
	return Deref(trackingCode.Cis)
}

func (trackingCode TrackingCode) GetCis1162() string {
	return Deref(trackingCode.Cis1162)
}

func (trackingCode TrackingCode) GetType() TrackingCodeType {
	return trackingCode.Type
}

func (trackingCode TrackingCode) GetTrackingCodes() Slice[TrackingCode] {
	return trackingCode.TrackingCodes
}

func (trackingCode TrackingCode) GetTrackingCode1162() Slice[TrackingCode] {
	return trackingCode.TrackingCode1162
}

func (trackingCode *TrackingCode) SetCis(cis string) *TrackingCode {
	trackingCode.Cis = &cis
	return trackingCode
}

func (trackingCode *TrackingCode) SetCis1162(cis1162 string) *TrackingCode {
	trackingCode.Cis1162 = &cis1162
	return trackingCode
}

func (trackingCode *TrackingCode) SetType(trackingCodeType TrackingCodeType) *TrackingCode {
	trackingCode.Type = trackingCodeType
	return trackingCode
}

func (trackingCode *TrackingCode) SetTrackingCodes(trackingCodes ...*TrackingCode) *TrackingCode {
	trackingCode.TrackingCodes = trackingCodes
	return trackingCode
}

func (trackingCode *TrackingCode) SetTrackingCode1162(trackingCode1162 ...*TrackingCode) *TrackingCode {
	trackingCode.TrackingCode1162 = trackingCode1162
	return trackingCode
}

func (trackingCode TrackingCode) String() string {
	return Stringify(trackingCode)
}

// TrackingCodeType Коды маркировки товаров и транспортных упаковок
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-kody-markirowki-towarow-i-transportnyh-upakowok
type TrackingCodeType string

const (
	TrackingCodeTypeTransportPack TrackingCodeType = "transportpack" // код маркировки товара
	TrackingCodeTypeConsumerPack  TrackingCodeType = "consumerpack"  // код маркировки потребительской упаковки
	TrackingCodeTypeTrackingCode  TrackingCodeType = "trackingcode"  // код транспортной упаковки
)
