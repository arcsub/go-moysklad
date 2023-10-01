package moysklad

import "github.com/google/uuid"

// ProcessingPosition общие поля для материала тех.операции и продукта тех.операции
type ProcessingPosition struct {
	AccountId  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/серии/модификации, которую представляет собой позиция
	Id         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции
}

func (p ProcessingPosition) String() string {
	return Stringify(p)
}
