package moysklad

import (
	"encoding/json"
	"reflect"
)

// Barcode Штрихкод.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-shtrihkody
type Barcode struct {
	Type  BarcodeType // Тип штрихкода
	Value string      // Штрихкод
}

func (barcode Barcode) String() string {
	return Stringify(barcode)
}

func NewBarcodeEAN13(value string) *Barcode {
	return &Barcode{BarcodeEAN13, value}
}

func NewBarcodeEAN8(value string) *Barcode {
	return &Barcode{BarcodeEAN8, value}
}

func NewBarcodeCode128(value string) *Barcode {
	return &Barcode{BarcodeCode128, value}
}

func NewBarcodeGTIN(value string) *Barcode {
	return &Barcode{BarcodeGTIN, value}
}

func NewBarcodeUPC(value string) *Barcode {
	return &Barcode{BarcodeUPC, value}
}

// BarcodeType Тип штрихкода
type BarcodeType string

const (
	BarcodeEAN13   BarcodeType = "ean13"   // штрихкод в формате EAN13, если требуется создать штрихкод в формате EAN13
	BarcodeEAN8    BarcodeType = "ean8"    // штрихкод в формате EAN8, если требуется создать штрихкод в формате EAN8
	BarcodeCode128 BarcodeType = "code128" // штрихкод в формате Code128, если требуется создать штрихкод в формате Code128
	BarcodeGTIN    BarcodeType = "gtin"    // штрихкод в формате GTIN, если требуется создать штрихкод в формате GTIN. Валидируется на соответствие формату GS1
	BarcodeUPC     BarcodeType = "upc"     // штрихкод в формате UPC, если требуется создать штрихкод в формате UPC
)

// MarshalJSON implements the json.Marshaler interface.
func (barcode Barcode) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{string(barcode.Type): barcode.Value})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (barcode *Barcode) UnmarshalJSON(bytes []byte) (err error) {
	tmp := map[string]string{}
	if err = json.Unmarshal(bytes, &tmp); err != nil {
		return
	}

	rv := reflect.ValueOf(tmp)
	mapKeys := rv.MapKeys()

	if len(mapKeys) > 0 {
		key := mapKeys[0]
		bType := BarcodeType(key.String())
		bValue := (rv.MapIndex(key)).String()

		barcode.Type = bType
		barcode.Value = bValue
	}
	return
}

type Barcodes = Slice[Barcode]

func NewBarcodes() Barcodes {
	return make(Barcodes, 0)
}
