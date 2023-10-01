package moysklad

import "github.com/google/uuid"

// Pack Упаковка Товара.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-upakowki-towara
type Pack struct {
	Barcodes *Barcodes  `json:"barcodes,omitempty"` // Массив штрихкодов упаковок товаров. Данный массив может содержать не более одного штрихкода. Если штрихкод в массиве отсутствует, то данное поле не выводится
	Id       *uuid.UUID `json:"id,omitempty"`       // ID упаковки товара
	Quantity *float64   `json:"quantity,omitempty"` // Количество Товаров в упаковке данного вида
	Uom      *Uom       `json:"uom,omitempty"`      // Метаданные единиц измерения
}

func (p Pack) String() string {
	return Stringify(p)
}

type Packs = Iterator[Pack]
