package export

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"project-root/modules/export/Enums"
	"project-root/modules/post/DB/Models"

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

func Csv(report []Models.Post, fileName string) (string, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	record, headers := getRecords(report)

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

func Excle(report []Models.Post, fileName string) (string, error) {
	file := excelize.NewFile()
	defer file.Close()
	record, headers := getRecords(report)
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}
	for j := 0; j < len(record); j++ {
		for i, rec := range record[j] {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), j+2), rec)
		}
	}

	if err := file.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return "", err
	}

	return fileName, nil
}

func FinalExport(report []Models.Post, method int) (string, error) {

	fileName := generateUniqueFileName("xlsx")
	var (
		err error
	)
	switch method {
	case Enums.CsvFileType:
		fileName := generateUniqueFileName("csv")
		_, err = Csv(report, fileName)
	case Enums.ExcelFileType:
		fileName := generateUniqueFileName("xlsx")
		_, err = Excle(report, fileName)
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
