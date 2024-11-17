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
	"gopkg.in/gomail.v2"
)

const (
	EXCLE      = 0
	CSV        = 1
	FILESYSTEM = "/tmp/"
	SMTPHOST   = "smtp.gmail.com"
	SMTPPORT   = 587
	MYEMAIL    = "crawlerweb91@gmail.com"
	MYPASS     = ""
)

func Csv(report []models.Posts, uuid, path string) (string, error) {
	fileName := fmt.Sprintf("%s/%s.%s", path, uuid, "csv")
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

	for _, rec := range record {
		err = writer.Write(rec)
		if err != nil {
			return "", err
		}
		writer.Flush()
	}

	return fileName, nil
}

func Excle(report []models.Posts, uuid, path string) (string, error) {
	fileName := fmt.Sprintf("%s/%s.%s", path, uuid, "xlsx")
	file := excelize.NewFile()
	defer file.Close()
	record, headers := records(report)
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}
	for j := 0; j < len(record); j++ {
		for i, rec := range record[j] {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), j+2), rec)
		}
	}

	if err := file.SaveAs(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}

func records(report []models.Posts) ([][]string, []string) {
	header := []string{
		"ID", "SrcSitesID", "CitiesID", "UsersID", "Status",
		"Title", "Description", "Price", "PriceHistory", "MainIMG",
		"GalleryIMGs", "SellerName", "LandArea", "BuiltYear",
		"Rooms", "IsApartment", "DealType", "Floors", "Elevator",
		"Storage", "Location", "PostDate", "DeletedAt", "CreatedAt", "UpdateAt",
	}
	var body [][]string
	for _, post := range report {
		body = append(body, []string{
			fmt.Sprintf("%v", post.ID),
			fmt.Sprintf("%v", post.SrcSitesID),
			fmt.Sprintf("%v", post.CitiesID),
			fmt.Sprintf("%v", post.UsersID),
			fmt.Sprintf("%v", post.Status),
			fmt.Sprintf("%s", post.ExternalSiteID),
			fmt.Sprintf("%s", post.Title),
			fmt.Sprintf("%s", post.Description),
			fmt.Sprintf("%v", post.Price),
			fmt.Sprintf("%s", post.PriceHistory),
			fmt.Sprintf("%v", post.MainIMG),
			fmt.Sprintf("%s", post.GalleryIMGs),
			fmt.Sprintf("%s", post.SellerName),
			fmt.Sprintf("%v", post.LandArea),
			fmt.Sprintf("%v", post.BuiltYear),
			fmt.Sprintf("%v", post.Rooms),
			fmt.Sprintf("%t", post.IsApartment),
			fmt.Sprintf("%v", post.DealType),
			fmt.Sprintf("%v", post.Floors),
			fmt.Sprintf("%t", post.Elevator),
			fmt.Sprintf("%t", post.Storage),
			fmt.Sprintf("%s", post.Location),
			fmt.Sprintf("%s", post.PostDate.String()),
			fmt.Sprintf("%s", post.DeletedAt.String()),
			fmt.Sprintf("%s", post.CreatedAt.String()),
			fmt.Sprintf("%s", post.UpdateAt.String()),
		})
	}
	return body, header
}

func Export(report []models.Posts, method int) (string, error) {
	var (
		uid      uuid.UUID
		fileName string
		err      error
	)

	uid = uuid.New()
	switch method {
	case CSV:
		fileName, err = Csv(report, uid.String(), FILESYSTEM)
	case EXCLE:
		fileName, err = Excle(report, uid.String(), FILESYSTEM)
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
	msg := gomail.NewMessage()
	msg.SetHeader("From", MYEMAIL)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Web crawler")
	msg.Attach(path)
	dial := gomail.NewDialer(SMTPHOST, SMTPPORT, MYEMAIL, MYPASS)
	err := dial.DialAndSend(msg)
	return err
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
