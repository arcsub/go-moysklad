package moysklad

// TrackingType Тип маркируемой продукции.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-tip-markiruemoj-produkcii
type TrackingType string

const (
	TrackingTypeFoodSupplement TrackingType = "FOOD_SUPPLEMENT" // [22-12-2023]
	TrackingTypeBeerAlcohol    TrackingType = "BEER_ALCOHOL"    // [26-10-2023]
	TrackingTypeElectronics    TrackingType = "ELECTRONICS"
	TrackingTypeClothes        TrackingType = "LP_CLOTHES"
	TrackingTypeLinens         TrackingType = "LP_LINENS"
	TrackingTypeMilk           TrackingType = "MILK"
	TrackingTypeNcp            TrackingType = "NCP"
	TrackingTypeNotTracked     TrackingType = "NOT_TRACKED"
	TrackingTypeOtp            TrackingType = "OTP"
	TrackingTypePerfumery      TrackingType = "PERFUMERY"
	TrackingTypeShoes          TrackingType = "SHOES"
	TrackingTypeTires          TrackingType = "TIRES"
	TrackingTypeTobacco        TrackingType = "TOBACCO"
	TrackingTypeWater          TrackingType = "WATER"
)
