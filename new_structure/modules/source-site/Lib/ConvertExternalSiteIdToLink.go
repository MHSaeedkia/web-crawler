package Lib

import "fmt"

func ConvertExternalSiteIdToLink(externalSiteId string, siteId int) string {
	switch siteId {
	case 1:
		return fmt.Sprintf("(https://divar.ir/v/%s)", externalSiteId)
	case 2:
		return fmt.Sprintf("(https://www.sheypoor.com/v/%s)", externalSiteId)
	default:
		return fmt.Sprintf("(https://divar.ir/v/%s)", externalSiteId)
	}
}
