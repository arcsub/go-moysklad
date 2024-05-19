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

func newPositions[T PositionTypes]() *Positions[T] {
	return &Positions[T]{Rows: make(Slice[T], 0)}
}

func NewBundleComponents() *Positions[BundleComponent] {
	return newPositions[BundleComponent]()
}

func NewCommissionReportInPositions() *Positions[CommissionReportInPosition] {
	return newPositions[CommissionReportInPosition]()
}

func NewCommissionReportInReturnPositions() *Positions[CommissionReportInReturnPosition] {
	return newPositions[CommissionReportInReturnPosition]()
}

func NewCommissionReportOutPositions() *Positions[CommissionReportOutPosition] {
	return newPositions[CommissionReportOutPosition]()
}

func NewCustomerOrderPositions() *Positions[CustomerOrderPosition] {
	return newPositions[CustomerOrderPosition]()
}

func NewDemandPositions() *Positions[DemandPosition] {
	return newPositions[DemandPosition]()
}

func NewEnterPositions() *Positions[EnterPosition] {
	return newPositions[EnterPosition]()
}

func NewInternalOrderPositions() *Positions[InternalOrderPosition] {
	return newPositions[InternalOrderPosition]()
}

func NewInventoryPositions() *Positions[InventoryPosition] {
	return newPositions[InventoryPosition]()
}

func NewInvoicePositions() *Positions[InvoicePosition] {
	return newPositions[InvoicePosition]()
}

func NewLossPositions() *Positions[LossPosition] {
	return newPositions[LossPosition]()
}

func NewMovePositions() *Positions[MovePosition] {
	return newPositions[MovePosition]()
}

func NewPrepaymentPositions() *Positions[PrepaymentPosition] {
	return newPositions[PrepaymentPosition]()
}

func NewPrepaymentReturnPositions() *Positions[PrepaymentReturnPosition] {
	return newPositions[PrepaymentReturnPosition]()
}

func NewPriceListPositions() *Positions[PriceListPosition] {
	return newPositions[PriceListPosition]()
}

func NewProcessingOrderPositions() *Positions[ProcessingOrderPosition] {
	return newPositions[ProcessingOrderPosition]()
}

func NewProcessingPlanMaterials() *Positions[ProcessingPlanMaterial] {
	return newPositions[ProcessingPlanMaterial]()
}

func NewProcessingPlanProducts() *Positions[ProcessingPlanProduct] {
	return newPositions[ProcessingPlanProduct]()
}

func NewProcessingProcessPositions() *Positions[ProcessingProcessPosition] {
	return newPositions[ProcessingProcessPosition]()
}

func NewPurchaseOrderPositions() *Positions[PurchaseOrderPosition] {
	return newPositions[PurchaseOrderPosition]()
}

func NewPurchaseReturnPositions() *Positions[PurchaseReturnPosition] {
	return newPositions[PurchaseReturnPosition]()
}

func NewRetailPositions() *Positions[RetailPosition] {
	return newPositions[RetailPosition]()
}

func NewSalesReturnPositions() *Positions[SalesReturnPosition] {
	return newPositions[SalesReturnPosition]()
}

func NewSupplyPositions() *Positions[SupplyPosition] {
	return newPositions[SupplyPosition]()
}

func NewProductionTaskMaterials() *Positions[ProductionTaskMaterial] {
	return newPositions[ProductionTaskMaterial]()
}

func NewProductionStageCompletionMaterials() *Positions[ProductionStageCompletionMaterial] {
	return newPositions[ProductionStageCompletionMaterial]()
}

func NewProductionStageCompletionResults() *Positions[ProductionStageCompletionResult] {
	return newPositions[ProductionStageCompletionResult]()
}

func NewProductionRows() *Positions[ProductionRow] {
	return newPositions[ProductionRow]()
}

func NewProductionTaskResults() *Positions[ProductionTaskResult] {
	return newPositions[ProductionTaskResult]()
}
