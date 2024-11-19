package Enums

const (
	CsvFileType   = 0
	ExcelFileType = 1
	ZipFileType   = 2
)

func ConvertFileTypeToStr(fileType int) string {
	status := ""
	switch fileType {
	case CsvFileType:
		status = "csv"
	case ExcelFileType:
		status = "xlsx"
	case ZipFileType:
		status = "zip"
	}
	return status
}
