package Enums

const (
	Rent = iota
	Sale
	Buy
)

func GetDealTypes() []string {
	return []string{
		"Rent",
		"Sale",
		"Buy",
	}
}
