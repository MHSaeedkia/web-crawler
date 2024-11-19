package export

import (
	"fmt"
	"github.com/google/uuid"
	"project-root/modules/post/DB/Models"
	"project-root/sys-modules/env"
	"time"
)

func generateUniqueFileName(format string) string {
	baseSavePath := env.Env("EXPORT_SAVE_PATH")
	newUUID := uuid.New()
	fileName := fmt.Sprintf("%s/%s.%s", baseSavePath, newUUID.String(), format)
	return fileName
}

func getRecords(report []Models.Post) ([][]string, []string) {
	header := []string{
		"ID", "SrcSitesID", "CitiesID", "UsersID", "Status",
		"Title", "Description", "Price", "MainIMG", "SellerName", "LandArea", "BuiltYear",
		"Rooms", "IsApartment", "DealType", "Floors", "Elevator",
		"Storage", "Location", "PostDate", "CreatedAt", "UpdateAt",
	}
	var body [][]string
	for _, post := range report {
		body = append(body, []string{
			fmt.Sprintf("%d", post.ID),
			fmt.Sprintf("%d", post.SrcSitesID),
			nullSafeString(post.CityName),
			nullSafeInt(post.UsersID),
			fmt.Sprintf("%d", post.Status),
			nullSafeString(post.ExternalSiteID),
			post.Title,
			nullSafeString(post.Description),
			nullSafeInt64(post.Price),
			nullSafeString(post.MainIMG),
			nullSafeString(post.SellerName),
			nullSafeFloat64(post.LandArea),
			nullSafeInt(post.BuiltYear),
			nullSafeInt(post.Rooms),
			nullSafeBool(post.IsApartment),
			nullSafeInt(post.DealType),
			nullSafeInt(post.Floors),
			nullSafeBool(post.Elevator),
			nullSafeBool(post.Storage),
			nullSafeString(post.Location),
			nullSafeTime(post.PostDate),
			post.CreatedAt.Format("2006-01-02 15:04:05"),
			nullSafeTime(post.UpdateAt),
		})
	}
	return body, header
}

func nullSafeString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func nullSafeInt(ptr *int) string {
	if ptr == nil {
		return ""
	}
	return fmt.Sprintf("%d", *ptr)
}

func nullSafeInt64(ptr *int64) string {
	if ptr == nil {
		return ""
	}
	return fmt.Sprintf("%d", *ptr)
}

func nullSafeFloat64(ptr *float64) string {
	if ptr == nil {
		return ""
	}
	return fmt.Sprintf("%.2f", *ptr)
}

func nullSafeBool(ptr *bool) string {
	if ptr == nil {
		return "false"
	}
	return fmt.Sprintf("%t", *ptr)
}

func nullSafeTime(ptr *time.Time) string {
	if ptr == nil {
		return ""
	}
	return ptr.Format("2006-01-02 15:04:05")
}
