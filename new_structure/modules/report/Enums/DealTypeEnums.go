package Enums

const (
	RentDealType = iota
	SaleDealType
	BuyDealType
)

func GetDealTypes() []string {
	return []string{
		"Rent",
		"Sale",
		"Buy",
	}
}
