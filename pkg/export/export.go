package export

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/MHSaeedkia/web-crawler/internal/models"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

const (
	EXCLE      = 0
	CSV        = 1
	fileSystem = "/tmp/"
)

func Csv(report models.Posts, uuid, path string) (string, error) {
	fileName := fmt.Sprintf("%s/%s.%s", path, uuid /*uuid[:len(string(uuid))-1]*/, "csv")
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	record, headers := records(report)

	err = writer.Write(headers)
	if err != nil {
		return "", err
	}

	err = writer.Write(record)
	if err != nil {
		return "", err
	}
	writer.Flush()

	return fileName, nil
}

func Excle(report models.Posts, uuid, path string) (string, error) {
	fileName := fmt.Sprintf("%s/%s.%s", path, uuid /*string(uuid)[:len(string(uuid))-1]*/, "xlsx")
	file := excelize.NewFile()
	defer file.Close()
	record, headers := records(report)
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}
	for j, col := range record {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), 2), col)
	}

	if err := file.SaveAs(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}

func records(report models.Posts) ([]string, []string) {
	header := []string{
		"ID", "SrcSitesID", "CitiesID", "UsersID", "Status",
		"Title", "Description", "Price", "PriceHistory", "MainIMG",
		"GalleryIMGs", "SellerName", "LandArea", "BuiltYear",
		"Rooms", "IsApartment", "DealType", "Floors", "Elevator",
		"Storage", "Location", "PostDate", "DeletedAt", "CreatedAt", "UpdateAt",
	}

	body := []string{
		fmt.Sprintf("%v", report.ID),
		fmt.Sprintf("%v", report.SrcSitesID),
		fmt.Sprintf("%v", report.CitiesID),
		fmt.Sprintf("%v", report.UsersID),
		fmt.Sprintf("%v", report.Status),
		fmt.Sprintf("%s", report.ExternalSiteID),
		fmt.Sprintf("%s", report.Title),
		fmt.Sprintf("%s", report.Description),
		fmt.Sprintf("%v", report.Price),
		fmt.Sprintf("%s", report.PriceHistory),
		fmt.Sprintf("%v", report.MainIMG),
		fmt.Sprintf("%s", report.GalleryIMGs),
		fmt.Sprintf("%s", report.SellerName),
		fmt.Sprintf("%v", report.LandArea),
		fmt.Sprintf("%v", report.BuiltYear),
		fmt.Sprintf("%v", report.Rooms),
		fmt.Sprintf("%t", report.IsApartment),
		fmt.Sprintf("%v", report.DealType),
		fmt.Sprintf("%v", report.Floors),
		fmt.Sprintf("%t", report.Elevator),
		fmt.Sprintf("%t", report.Storage),
		fmt.Sprintf("%s", report.Location),
		fmt.Sprintf("%s", report.PostDate.String()),
		fmt.Sprintf("%s", report.DeletedAt.String()),
		fmt.Sprintf("%s", report.CreatedAt.String()),
		fmt.Sprintf("%s", report.UpdateAt.String()),
	}

	return body, header
}

func Export(report models.Posts, method int) (string, error) {
	var (
		uid      uuid.UUID
		fileName string
		err      error
	)

	uid = uuid.New()
	switch method {
	case CSV:
		fileName, err = Csv(report, uid.String(), fileSystem)
	case EXCLE:
		fileName, err = Excle(report, uid.String(), fileSystem)
	}
	return fileName, err
}

func ZipExport(pahts []string, path string) (string, error) {
	var (
		uid      uuid.UUID
		fileName string
		err      error
	)
	uid = uuid.New()

	fileName = fmt.Sprintf("%s/%s.zip", path, uid /*string(uuid)[:len(string(uuid))-1]*/)

	zipFile, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range pahts {
		fileToZip, err := os.Open(file)
		if err != nil {
			return "", err
		}
		defer fileToZip.Close()

		fileInfo, err := fileToZip.Stat()
		if err != nil {
			return "", err
		}
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return "", err
		}

		header.Name = file

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			return "", err
		}
	}

	return fileName, nil
}

func EmailExport(email, path string) error {

	return nil
}

func DeleteExport(pahts []string) error {
	for _, file := range pahts {
		err := os.Remove(file)
		if err != nil {
			return err
		}
	}
	return nil
}
