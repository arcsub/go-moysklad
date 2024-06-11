package moysklad

import "github.com/google/uuid"

// Pack Упаковка Товара.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-upakowki-towara
type Pack struct {
	ID       *uuid.UUID     `json:"id,omitempty"`
	Quantity *float64       `json:"quantity,omitempty"`
	Uom      *Uom           `json:"uom,omitempty"`
	Barcodes Slice[Barcode] `json:"barcodes,omitempty"`
}

func (pack Pack) GetID() uuid.UUID {
	return Deref(pack.ID)
}

func (pack Pack) GetQuantity() float64 {
	return Deref(pack.Quantity)
}

func (pack Pack) GetUom() Uom {
	return Deref(pack.Uom)
}

func (pack Pack) GetBarcodes() Slice[Barcode] {
	return pack.Barcodes
}

func (pack *Pack) SetQuantity(quantity float64) *Pack {
	pack.Quantity = &quantity
	return pack
}

func (pack *Pack) SetUom(uom *Uom) *Pack {
	pack.Uom = uom
	return pack
}

func (pack *Pack) SetBarcodes(barcodes Slice[Barcode]) *Pack {
	pack.Barcodes = barcodes
	return pack
}

func (pack Pack) String() string {
	return Stringify(pack)
}
