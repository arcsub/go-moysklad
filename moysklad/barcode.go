package moysklad

import (
	"github.com/goccy/go-json"
	"reflect"
)

// Barcode Штрихкод.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-shtrihkody
type Barcode struct {
	Type  BarcodeType // Тип штрихкода
	Value string      // Штрихкод
}

// String реализует интерфейс [fmt.Stringer].
func (barcode Barcode) String() string {
	return Stringify(barcode)
}

// NewBarcodeEAN13 принимает значение штрихкода типа EAN13 и возвращает готовый [Barcode] с типом [BarcodeEAN13]
func NewBarcodeEAN13(value string) *Barcode {
	return &Barcode{BarcodeEAN13, value}
}

// NewBarcodeEAN8 принимает значение штрихкода типа EAN8 и возвращает готовый [Barcode] с типом [BarcodeEAN8]
func NewBarcodeEAN8(value string) *Barcode {
	return &Barcode{BarcodeEAN8, value}
}

// NewBarcodeCode128 принимает значение штрихкода типа code128 и возвращает готовый [Barcode] с типом [BarcodeCode128]
func NewBarcodeCode128(value string) *Barcode {
	return &Barcode{BarcodeCode128, value}
}

// NewBarcodeGTIN принимает значение штрихкода типа GTIN и возвращает готовый [Barcode] с типом [BarcodeGTIN]
func NewBarcodeGTIN(value string) *Barcode {
	return &Barcode{BarcodeGTIN, value}
}

// NewBarcodeUPC принимает значение штрихкода типа UPC и возвращает готовый [Barcode] с типом [BarcodeUPC]
func NewBarcodeUPC(value string) *Barcode {
	return &Barcode{BarcodeUPC, value}
}

// BarcodeType Тип штрихкода
//
// Возможные значения:
//   - BarcodeEAN13   – штрихкод в формате EAN13, если требуется создать штрихкод в формате EAN13
//   - BarcodeEAN8    – штрихкод в формате EAN8, если требуется создать штрихкод в формате EAN8
//   - BarcodeCode128 – штрихкод в формате Code128, если требуется создать штрихкод в формате Code128
//   - BarcodeGTIN    – штрихкод в формате GTIN, если требуется создать штрихкод в формате GTIN. Валидируется на соответствие формату GS1
//   - BarcodeUPC     – штрихкод в формате UPC, если требуется создать штрихкод в формате UPC
type BarcodeType string

const (
	BarcodeEAN13   BarcodeType = "ean13"   // штрихкод в формате EAN13, если требуется создать штрихкод в формате EAN13
	BarcodeEAN8    BarcodeType = "ean8"    // штрихкод в формате EAN8, если требуется создать штрихкод в формате EAN8
	BarcodeCode128 BarcodeType = "code128" // штрихкод в формате Code128, если требуется создать штрихкод в формате Code128
	BarcodeGTIN    BarcodeType = "gtin"    // штрихкод в формате GTIN, если требуется создать штрихкод в формате GTIN. Валидируется на соответствие формату GS1
	BarcodeUPC     BarcodeType = "upc"     // штрихкод в формате UPC, если требуется создать штрихкод в формате UPC
)

// MarshalJSON реализует интерфейс [json.Marshaler]
func (barcode Barcode) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{string(barcode.Type): barcode.Value})
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler]
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

// NewBarcodes возвращает пустой срез для удобной работы с множеством объектов типа [Barcode].
func NewBarcodes() Slice[Barcode] {
	return NewSlice[Barcode]()
}
