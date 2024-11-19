package Download

import (
	tele "gopkg.in/telebot.v4"
	"project-root/modules/export/Enums"
	"project-root/modules/export/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type SendFileExportInRobotMethodPage struct{}

func (p *SendFileExportInRobotMethodPage) PageNumber() int {
	return Enums.SendFileExportInRobotMethodPageNumber
}

func (p *SendFileExportInRobotMethodPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	exportIds := telSession.GetExportTempData().ExportIds
	if len(exportIds) > 1 {
		// zip
		panic("Zip not allow")
	}
	export, _ := Facades.ExportRepo().FindByID(exportIds[0])
	fileFormat := Enums.ConvertFileTypeToStr(export.FileType)
	//var newReplyMarkup = &tele.ReplyMarkup{}
	return &Page.PageContentOV{
		Message:     "File Exported:",
		ReplyMarkup: StaticBtns.GetBackStaticBtn(),
		File: &tele.Document{
			File:     tele.FromDisk(export.FilePath),
			FileName: "export." + fileFormat,
		},
	}
}

func (p *SendFileExportInRobotMethodPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *SendFileExportInRobotMethodPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return StaticBtns.HandleIfClickBackBtn(btnKey, Enums.MainExportPageNumber)
}

var _ Page.PageInterface = &SendFileExportInRobotMethodPage{}
