package moysklad

import "encoding/json"

type PositionTypes interface {
	BundleComponent | CommissionReportInPosition | CommissionReportInReturnPosition |
		CommissionReportOutPosition | CustomerOrderPosition | DemandPosition | EnterPosition |
		InternalOrderPosition | InventoryPosition | InvoicePosition | LossPosition | MovePosition |
		PrepaymentPosition | PrepaymentReturnPosition | PriceListPosition | ProcessingOrderPosition |
		ProcessingPlanMaterial | ProcessingPlanProduct | ProcessingProcessPosition | PurchaseOrderPosition |
		PurchaseReturnPosition | RetailPosition | SalesReturnPosition | SupplyPosition | ProductionTaskMaterial |
		ProductionStageCompletionMaterial | ProductionStageCompletionResult | ProductionRow | ProductionTaskResult
}

type Positions[T PositionTypes] MetaArray[T]

// MarshalJSON implements the json.Marshaler interface.
func (p Positions[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Rows)
}

func (p *Positions[T]) Push(elements ...*T) {
	p.Rows = append(p.Rows, elements...)

	if len(p.Rows) > MaxPositions {
		p.Rows = p.Rows[:MaxPositions]
	}
}

func NewPositions[T PositionTypes]() *Positions[T] {
	return &Positions[T]{Rows: make(Slice[T], 0)}
}
