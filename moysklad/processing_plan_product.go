package moysklad

import (
	"github.com/google/uuid"
)

// ProcessingPlanProduct Продукт Тех. карты.
// Ключевое слово: processingplanresult
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-teh-karty-produkty-teh-karty
type ProcessingPlanProduct struct {
	AccountId  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара или модификации позиции
	Id         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Product    *Product            `json:"product,omitempty"`    // Метаданные товара позиции. В случае, если в поле assortment указана модификация, то это поле содержит товар, к которому относится эта модификация
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции
}

func (p ProcessingPlanProduct) String() string {
	return Stringify(p)
}

func (p ProcessingPlanProduct) MetaType() MetaType {
	return MetaTypeProcessingPlanProduct
}
