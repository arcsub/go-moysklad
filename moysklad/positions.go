package moysklad

type PositionTypes interface {
	BundleComponent | File | Image | CommissionReportInPosition | CommissionReportInReturnPosition |
		CommissionReportOutPosition | CustomerOrderPosition | DemandPosition | EnterPosition |
		InternalOrderPosition | InventoryPosition | InvoicePosition | LossPosition | MovePosition |
		PrepaymentPosition | PrepaymentReturnPosition | PriceListPosition | ProcessingOrderPosition |
		ProcessingPlanMaterial | ProcessingPlanProduct | ProcessingProcessPosition | PurchaseOrderPosition |
		PurchaseReturnPosition | RetailPosition | SalesReturnPosition | SupplyPosition
}

type Positions[T PositionTypes] MetaArray[T]

func (p *Positions[T]) Iter() *Iterator[T] {
	return p.Rows.Iter(MaxPositions)
}

func (p *Positions[T]) Set(elements Slice[T]) {
	if len(elements) > MaxPositions {
		elements = elements[:MaxPositions]
	}
	p.Rows = elements
}

func (p *Positions[T]) Push(elements ...T) {
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
