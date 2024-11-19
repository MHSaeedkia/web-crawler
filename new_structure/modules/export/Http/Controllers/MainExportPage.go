package Controllers

import (
	"errors"
	"fmt"
	tele "gopkg.in/telebot.v4"
	Models3 "project-root/modules/export/DB/Models"
	exportEnums "project-root/modules/export/Enums"
	"project-root/modules/export/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/modules/user/Enums"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
	"project-root/sys-modules/time/Lib"
	"strconv"
	"strings"
)

type MainExportPage struct{}

func (p *MainExportPage) PageNumber() int {
	return exportEnums.MainExportPageNumber
}

func (p *MainExportPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {

	exports, countAllPage, _ := Facades.ExportRepo().GetExportsWithPagination(*telSession.LoggedUserID, StaticBtns.GetDefaultPerPage(), telSession.GetExportTempData().LastPageNumber)
	//crawlLogs, countAllPage, _ := Facades2.CrawlLogRepo().GetCrawlLogsWithPagination(StaticBtns.GetDefaultPerPage(), telSession.GetMonitoringTempData().LastPageNumber)

	// dynamic btn
	paginationReplyMarkup := StaticBtns.PaginationReplyMarkupData{
		Items:        []StaticBtns.PaginationReplyMarkupItem{},
		StaticRowBtn: []tele.Row{},
	}

	// custom btn
	return &Page.PageContentOV{
		Message:     formatExportDetails(*exports, telSession.GetExportTempData().LastPageNumber, countAllPage),
		ReplyMarkup: paginationReplyMarkup.GetReplyMarkup("export"),
	}
}

func formatExportDetails(exports []Models3.Export, currentPage int, totalPages int) string {
	var result strings.Builder

	result.WriteString("Select the export ids to download, separate each one with an “,” for example “12,2,3”\n❓Tip: If you select multiple files, they will be zipped\n--------------------------------------------------\nExports By Report Name\n\n")
	for _, export := range exports {
		reportTitle := export.Report.Title
		fileType := exportEnums.ConvertFileTypeToStr(export.FileType)
		createdAt := Lib.FormatTimeAgo(export.CreatedAt)
		result.WriteString(fmt.Sprintf("%s\n ( %d ) - %s - %s\n", reportTitle, export.ID, createdAt, fileType))
		result.WriteString("---------\n")
	}
	result.WriteString(strconv.Itoa(currentPage) + " page of " + strconv.Itoa(totalPages))

	return result.String()
}

func (p *MainExportPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	// validation input ids
	ids, err := ParseIDs(value)
	if err != nil {
		telSession.GetGeneralTempData().LastMessage = err.Error()
		return Page.GetPage(p.PageNumber())
	}
	hasError := Facades.ExportRepo().ValidateExportIDs(ids, *telSession.LoggedUserID)
	if hasError != nil {
		telSession.GetGeneralTempData().LastMessage = hasError.Error()
		return Page.GetPage(p.PageNumber())
	}
	telSession.GetExportTempData().ExportIds = ids
	return Page.GetPage(exportEnums.SelectExportMethodPageNumber)
}

func ParseIDs(input string) ([]int, error) {
	if input == "" {
		return nil, errors.New("input is empty")
	}

	parts := strings.Split(input, ",")
	if len(parts) > 5 {
		return nil, errors.New("too many IDs, maximum allowed is 5")
	}

	ids := make([]int, 0, len(parts))
	idMap := make(map[int]bool)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, errors.New("empty ID found in input")
		}

		id, err := strconv.Atoi(part)
		if err != nil || id <= 0 {
			return nil, errors.New("invalid ID: " + part)
		}

		// Check for duplicates
		if idMap[id] {
			return nil, errors.New("duplicate ID found: " + part)
		}

		idMap[id] = true
		ids = append(ids, id)
	}

	return ids, nil
}

func (p *MainExportPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	// pagination btn
	paginationHandleBtn := StaticBtns.PaginationOnClickBtnData{
		PrefixBtnKey:      "export",
		CurrentPageNumber: p.PageNumber(),
		BackPageNumber:    Enums.MainUserPageNumber,
		GetPageNumberSaved: func() int {
			return telSession.GetExportTempData().LastPageNumber
		},
		SavePageNumber: func(pageNum int) {
			telSession.GetExportTempData().LastPageNumber = pageNum
		},
		OnSelectItemId: func(itemId int) Page.PageInterface {
			return nil
		},
	}
	return paginationHandleBtn.HandleInputPagination(btnKey)
}

var _ Page.PageInterface = &MainExportPage{}
