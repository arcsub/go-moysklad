package moysklad

// TrackingCode Коды маркировки
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kody-markirowki-kod-markirowki-atributy-suschnosti
type TrackingCode struct {
	ID               *string             `json:"id,omitempty"`                 // ID кода маркировки
	Cis              *string             `json:"cis,omitempty"`                // Код маркировки в стандартном формате
	Cis1162          *string             `json:"cis_1162,omitempty"`           // Код маркировки в формате тега 1162
	Type             TrackingCodeType    `json:"type,omitempty"`               // Тип кода маркировки
	TrackingCodes    Slice[TrackingCode] `json:"trackingCodes,omitempty"`      // Массив вложенных кодов маркировки. Может присутствовать, только если type имеет значения consumerpack или transportpack
	TrackingCode1162 Slice[TrackingCode] `json:"trackingCodes_1162,omitempty"` // Массив вложенных кодов маркировки. Может присутствовать, только если type имеет значения consumerpack или transportpack
}

// GetID возвращает ID кода маркировки.
func (trackingCode TrackingCode) GetID() string {
	return Deref(trackingCode.ID)
}

// GetCis возвращает Код маркировки в стандартном формате.
func (trackingCode TrackingCode) GetCis() string {
	return Deref(trackingCode.Cis)
}

// GetCis1162 возвращает Код маркировки в формате тега 1162.
func (trackingCode TrackingCode) GetCis1162() string {
	return Deref(trackingCode.Cis1162)
}

// GetType возвращает Тип кода маркировки.
func (trackingCode TrackingCode) GetType() TrackingCodeType {
	return trackingCode.Type
}

// GetTrackingCodes возвращает Массив вложенных кодов маркировки.
//
// Может присутствовать, только если type имеет значения [TrackingCodeTypeConsumerPack] или [TrackingCodeTypeTransportPack].
func (trackingCode TrackingCode) GetTrackingCodes() Slice[TrackingCode] {
	return trackingCode.TrackingCodes
}

// GetTrackingCode1162 возвращает Массив вложенных кодов маркировки в формате тега 1162.
//
// Может присутствовать, только если type имеет значения [TrackingCodeTypeConsumerPack] или [TrackingCodeTypeTransportPack].
func (trackingCode TrackingCode) GetTrackingCode1162() Slice[TrackingCode] {
	return trackingCode.TrackingCode1162
}

// SetCis устанавливает Код маркировки в стандартном формате.
func (trackingCode *TrackingCode) SetCis(cis string) *TrackingCode {
	trackingCode.Cis = &cis
	return trackingCode
}

// SetCis1162 устанавливает Код маркировки в формате тега 1162.
func (trackingCode *TrackingCode) SetCis1162(cis1162 string) *TrackingCode {
	trackingCode.Cis1162 = &cis1162
	return trackingCode
}

// SetType устанавливает Тип кода маркировки.
func (trackingCode *TrackingCode) SetType(trackingCodeType TrackingCodeType) *TrackingCode {
	trackingCode.Type = trackingCodeType
	return trackingCode
}

// SetTypeTransportPack устанавливает Тип кода маркировки в значение [TrackingCodeTypeTransportPack].
func (trackingCode *TrackingCode) SetTypeTransportPack() *TrackingCode {
	trackingCode.Type = TrackingCodeTypeTransportPack
	return trackingCode
}

// SetTypeConsumerPack устанавливает Тип кода маркировки в значение [TrackingCodeTypeConsumerPack].
func (trackingCode *TrackingCode) SetTypeConsumerPack() *TrackingCode {
	trackingCode.Type = TrackingCodeTypeConsumerPack
	return trackingCode
}

// SetTypeTrackingCode устанавливает Тип кода маркировки в значение [TrackingCodeTypeTrackingCode].
func (trackingCode *TrackingCode) SetTypeTrackingCode() *TrackingCode {
	trackingCode.Type = TrackingCodeTypeTrackingCode
	return trackingCode
}

// SetTrackingCodes устанавливает Массив вложенных кодов маркировки.
//
// Принимает множество объектов [TrackingCode].
//
// Может присутствовать, только если type имеет значения [TrackingCodeTypeConsumerPack] или [TrackingCodeTypeTransportPack].
func (trackingCode *TrackingCode) SetTrackingCodes(trackingCodes ...*TrackingCode) *TrackingCode {
	trackingCode.TrackingCodes.Push(trackingCodes...)
	return trackingCode
}

// SetTrackingCode1162 устанавливает Массив вложенных кодов маркировкив формате тега 1162.
//
// Принимает множество объектов [TrackingCode].
//
// Может присутствовать, только если type имеет значения [TrackingCodeTypeConsumerPack] или [TrackingCodeTypeTransportPack].
func (trackingCode *TrackingCode) SetTrackingCode1162(trackingCode1162 ...*TrackingCode) *TrackingCode {
	trackingCode.TrackingCode1162.Push(trackingCode1162...)
	return trackingCode
}

// String реализует интерфейс [fmt.Stringer].
func (trackingCode TrackingCode) String() string {
	return Stringify(trackingCode)
}

// TrackingCodeType Коды маркировки товаров и транспортных упаковок.
//
// Возможные значения:
//   - TrackingCodeTypeTransportPack – Код маркировки товара
//   - TrackingCodeTypeConsumerPack  – Код маркировки потребительской упаковки
//   - TrackingCodeTypeTrackingCode  – Код транспортной упаковки
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-kody-markirowki-towarow-i-transportnyh-upakowok
type TrackingCodeType string

const (
	TrackingCodeTypeTransportPack TrackingCodeType = "transportpack" // Код маркировки товара
	TrackingCodeTypeConsumerPack  TrackingCodeType = "consumerpack"  // Код маркировки потребительской упаковки
	TrackingCodeTypeTrackingCode  TrackingCodeType = "trackingcode"  // Код транспортной упаковки
)
