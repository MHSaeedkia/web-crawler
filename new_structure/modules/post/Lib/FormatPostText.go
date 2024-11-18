package Lib

import (
	"fmt"
	"project-root/modules/post/DB/Models"
	PostEnums "project-root/modules/post/Enums"
	"project-root/modules/report/Enums"
	Lib2 "project-root/modules/source-site/Lib"
	"project-root/sys-modules/time/Lib"
	"strings"
	"time"
)

func FormatPostText(post *Models.Post) string {
	var result strings.Builder

	// --
	result.WriteString(fmt.Sprintf(
		"%s\nprice:                     %s\nstatus:                  %s\nref:                         %s\nseller name:        %s\n",
		post.Title,
		formatPrice(post.Price),
		formatStatus(post.Status),
		getRef(post.ExternalSiteID),
		getSellerName(post.SellerName),
	))

	// --
	result.WriteString(fmt.Sprintf(
		"bookmark:           %s\npublic:                   %s\ncreated at:            %s\nupdated at:          %s\nlink:                         [open](%s)\n",
		formatBool(post.IsPublic),
		formatBool(false),
		Lib.FormatTimeAgo(post.CreatedAt),
		Lib.FormatTimeAgo(*post.UpdateAt),
		Lib2.ConvertExternalSiteIdToLink(*post.ExternalSiteID, post.SrcSitesID),
	))

	if post.Description != nil {
		result.WriteString(fmt.Sprintf("description:          %s\n", *post.Description))
	}

	result.WriteString("---------------------- details --------------------\n")
	result.WriteString(fmt.Sprintf(
		"land area:               %s\nbuilt year:               %s\nroom count:          %s\nis apartment:        %s\ndeal_type:              %s\nfloor_count:          %s\nhas_elevator:       %s\nhas_storage:       %s\npost publish date: %s\ncity:                       %s\nneighborhood:  %s\n",
		formatNullableFloat(post.LandArea),
		formatNullableInt(post.BuiltYear),
		formatNullableInt(post.Rooms),
		formatBoolPtr(post.IsApartment),
		formatDealType(post.DealType),
		formatNullableInt(post.Floors),
		formatBoolPtr(post.Elevator),
		formatBoolPtr(post.Storage),
		Lib.FormatTimeAgo(*post.PostDate),
		getNullableString(post.CityName),
		getNullableString(post.NeighborhoodName),
	))

	return result.String()
}

// Helper functions

func formatPrice(price *int64) string {
	if price == nil {
		return "N/A"
	}
	return fmt.Sprintf("%d", *price)
}

func formatStatus(status int) string {
	switch status {
	case PostEnums.ProcessingPostStatus:
		return "processing"
	case PostEnums.SuccessfulPostStatus:
		return "successful"
	default:
		return "unknown"
	}
}

func getRef(externalSiteID *string) string {
	if externalSiteID == nil {
		return "N/A"
	}
	return *externalSiteID
}

func getSellerName(sellerName *string) string {
	if sellerName == nil {
		return "N/A"
	}
	return *sellerName
}

func formatBool(value bool) string {
	if value {
		return "✅"
	}
	return "❌"
}

func formatBoolPtr(value *bool) string {
	if value == nil {
		return "N/A"
	}
	return formatBool(*value)
}

func formatNullableFloat(value *float64) string {
	if value == nil {
		return "N/A"
	}
	return fmt.Sprintf("%.2f", *value)
}

func formatNullableInt(value *int) string {
	if value == nil {
		return "N/A"
	}
	return fmt.Sprintf("%d", *value)
}

func formatNullableTime(value *time.Time) string {
	if value == nil {
		return "N/A"
	}
	return value.Format("2006-01-02")
}

func getNullableString(value *string) string {
	if value == nil {
		return "N/A"
	}
	return *value
}

func formatDealType(value *int) string {
	if value == nil {
		return "N/A"
	}

	switch *value {
	case Enums.SaleDealType:
		return "sale"
	case Enums.RentDealType:
		return "rent"
	case Enums.BuyDealType:
		return "buy"
	default:
		return "unknown"
	}
}
