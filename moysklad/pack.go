package moysklad

// Pack Упаковка Товара.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-upakowki-towara
type Pack struct {
	ID       *string        `json:"id,omitempty"`       // ID упаковки товара
	Quantity *float64       `json:"quantity,omitempty"` // Количество Товаров в упаковке данного вида
	Uom      *Uom           `json:"uom,omitempty"`      // Единица измерения
	Barcodes Slice[Barcode] `json:"barcodes,omitempty"` // Штрихкоды
}

// GetID возвращает ID упаковки товара.
func (pack Pack) GetID() string {
	return Deref(pack.ID)
}

// GetQuantity возвращает Количество Товаров в упаковке данного вида.
func (pack Pack) GetQuantity() float64 {
	return Deref(pack.Quantity)
}

// GetUom возвращает Единицу измерения.
func (pack Pack) GetUom() Uom {
	return Deref(pack.Uom)
}

// GetBarcodes возвращает Штрихкоды.
func (pack Pack) GetBarcodes() Slice[Barcode] {
	return pack.Barcodes
}

// SetQuantity устанавливает Количество Товаров в упаковке данного вида.
func (pack *Pack) SetQuantity(quantity float64) *Pack {
	pack.Quantity = &quantity
	return pack
}

// SetUom устанавливает Единицу измерения.
func (pack *Pack) SetUom(uom *Uom) *Pack {
	if uom != nil {
		pack.Uom = uom.Clean()
	}
	return pack
}

// SetBarcodes устанавливает Штрихкоды.
//
// Для обновления списка штрихкодов необходимо передавать их полный список, включающий как старые,
// так и новые значения. Отсутствующие значения штрихкодов при обновлении будут удалены.
// При обновлении списка штрихкодов валидируются только новые значения.
// Ранее сохраненные штрихкоды не валидируются.
// То есть, если один из старых штрихкодов не соответствует требованиям к валидации,
// то ошибки при обновлении списка не будет. Если на вход передан пустой список штрихкодов
// или список из пустых значений, то ранее созданные штрихкоды будут удалены.
//
// Принимает множество объектов [Barcode].
func (pack *Pack) SetBarcodes(barcodes ...*Barcode) *Pack {
	pack.Barcodes = barcodes
	return pack
}

// String реализует интерфейс [fmt.Stringer].
func (pack Pack) String() string {
	return Stringify(pack)
}
