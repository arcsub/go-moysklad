package moysklad

import (
	"github.com/google/uuid"
)

// ProcessingPlanMaterial Материал Тех. карты.
// Ключевое слово: processingplanmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-teh-karty-materialy-teh-karty
type ProcessingPlanMaterial struct {
	AccountID                 *uuid.UUID          `json:"accountId,omitempty"`                 // ID учетной записи
	Assortment                *AssortmentPosition `json:"assortment,omitempty"`                // Метаданные товара или модификации позиции
	ID                        *uuid.UUID          `json:"id,omitempty"`                        // ID позиции
	Product                   *Product            `json:"product,omitempty"`                   // Метаданные товара позиции. В случае, если в поле assortment указана модификация, то это поле содержит товар, к которому относится эта модификация
	Quantity                  *float64            `json:"quantity,omitempty"`                  // Количество товаров данного вида в позиции
	ProcessingProcessPosition *Meta               `json:"processingProcessPosition,omitempty"` // Метаданные позиции Тех. процесса
	MaterialProcessingPlan    *Meta               `json:"materialProcessingPlan"`              // Метаданные техкарты материала [11-01-2024]
}

func (p ProcessingPlanMaterial) String() string {
	return Stringify(p)
}

func (p ProcessingPlanMaterial) MetaType() MetaType {
	return MetaTypeProcessingPlanMaterial
}
