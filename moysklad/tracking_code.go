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

func (t TrackingCode) String() string {
	return Stringify(t)
}

type TrackingCodes = Slice[TrackingCode]

// TrackingCodeType Коды маркировки товаров и транспортных упаковок
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-kody-markirowki-towarow-i-transportnyh-upakowok
type TrackingCodeType string

const (
	TrackingCodeTypeTransportPack TrackingCodeType = "transportpack" // код маркировки товара
	TrackingCodeTypeConsumerPack  TrackingCodeType = "consumerpack"  // код маркировки потребительской упаковки
	TrackingCodeTypeTrackingCode  TrackingCodeType = "trackingcode"  // код транспортной упаковки
)
