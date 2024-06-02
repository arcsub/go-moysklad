package moysklad

import "github.com/google/uuid"

// Pack Упаковка Товара.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-upakowki-towara
type Pack struct {
	ID       *uuid.UUID `json:"id,omitempty"`
	Quantity *float64   `json:"quantity,omitempty"`
	Uom      *Uom       `json:"uom,omitempty"`
	Barcodes Barcodes   `json:"barcodes,omitempty"`
}

func (p Pack) String() string {
	return Stringify(p)
}

type Packs = Slice[Pack]
