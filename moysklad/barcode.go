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

func (b Barcode) String() string {
	return Stringify(b)
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

func (b Barcode) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{string(b.Type): b.Value})
}

func (b *Barcode) UnmarshalJSON(bytes []byte) (err error) {
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

		b.Type = bType
		b.Value = bValue
	}
	return
}

type Barcodes = Iterator[Barcode]
