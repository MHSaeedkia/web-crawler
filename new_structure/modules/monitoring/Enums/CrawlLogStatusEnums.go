package Enums

const (
	ProcessingCrawlLogStatus = 0
	SuccessfulCrawlLogStatus = 1
	FailCrawlLogStatus       = 2
)

func GetCrawlLogStatusLabel(statusEnum int) string {
	status := ""
	switch statusEnum {
	case ProcessingCrawlLogStatus:
		status = "⏳processing"
	case SuccessfulCrawlLogStatus:
		status = "✅successful"
	case FailCrawlLogStatus:
		status = "❌failed"
	}
	return status
}
