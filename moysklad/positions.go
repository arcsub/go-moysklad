package moysklad

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

func (p *Positions[T]) Push(elements ...*T) {
	if len(elements) > MaxPositions {
		elements = elements[:MaxPositions]
	}
	limit := MaxPositions - len(p.Rows)
	if limit < 0 {
		limit = 0
	}
	elements = elements[:limit]
	p.Rows = append(p.Rows, elements...)
}

func NewPositions[T PositionTypes]() *Positions[T] {
	return &Positions[T]{}
}
