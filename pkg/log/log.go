package log

import (
	"fmt"
	"log"
	"os"
)

const (
	PATH    = "/tmp/"
	TELEBOT = "telebot"
	WEBC    = "webcrawler"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

var (
	WarningTelLog  *log.Logger
	InfoTelLog     *log.Logger
	ErrorTelLog    *log.Logger
	WarningWebCLog *log.Logger
	InfoWebCLog    *log.Logger
	ErrorWebCLog   *log.Logger
)

func initLog(path string) error {
	var (
		teleLogFile, webCLogFile *os.File
		err                      error
	)

	teleLogFile, err = os.OpenFile(fmt.Sprintf("%stelegrambot.log", path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	InfoTelLog = log.New(teleLogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningTelLog = log.New(teleLogFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorTelLog = log.New(teleLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	webCLogFile, err = os.OpenFile(fmt.Sprintf("%swebcrawler.log", path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	InfoWebCLog = log.New(webCLogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningWebCLog = log.New(webCLogFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorWebCLog = log.New(webCLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return err
}

func Log(message, service, status string) error {
	var (
		err error
	)

	err = initLog(PATH)
	if err != nil {
		return err
	}

	if service == TELEBOT {
		switch status {
		case INFO:
			InfoTelLog.Println(message)
		case WARNING:
			WarningTelLog.Println(message)
		case ERROR:
			ErrorTelLog.Println(message)
		}
	} else {
		switch status {
		case INFO:
			InfoWebCLog.Println(message)
		case WARNING:
			WarningWebCLog.Println(message)
		case ERROR:
			ErrorWebCLog.Println(message)
		}
	}
	return err
}
