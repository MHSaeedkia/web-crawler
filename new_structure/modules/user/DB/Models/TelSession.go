package Models

import "encoding/json"

type TelSession struct {
	ID           int                    `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	LoggedUserID *int                   `gorm:"column:logged_users_id;type:mediumint;default:null"`
	LoggedUser   *User                  `gorm:"foreignKey:LoggedUserID;references:ID"`
	ChatID       int64                  `gorm:"column:chat_id;type:bigint;unique"`
	CurrentPage  int                    `gorm:"column:current_page_num;type:mediumint"`
	TempData     map[string]interface{} `gorm:"column:temp_data;serializer:json"`
}

// ---------- Temp Data -----------

// ------- [General]
type GeneralTempData struct {
	LastMessage string `json:"last_message"`
}

func (ts *TelSession) GetGeneralTempData() *GeneralTempData {
	saveKey := "general"
	if existItem, ok := ts.TempData[saveKey].(*GeneralTempData); ok {
		return existItem
	}

	// deserialize
	if itemMap, ok := ts.TempData[saveKey].(map[string]interface{}); ok {
		jsonData, err := json.Marshal(itemMap)
		if err == nil {
			var item GeneralTempData
			err := json.Unmarshal(jsonData, &item)
			if err == nil {
				ts.TempData[saveKey] = &item
				return &item
			}
		}
	}

	// default
	defaultValue := &GeneralTempData{
		LastMessage: "",
	}
	ts.TempData[saveKey] = defaultValue
	return defaultValue
}

// ------- [Auth]
type AuthTempData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (ts *TelSession) GetAuthTempData() *AuthTempData {
	saveKey := "auth"
	if existItem, ok := ts.TempData[saveKey].(*AuthTempData); ok {
		return existItem
	}

	// deserialize
	if itemMap, ok := ts.TempData[saveKey].(map[string]interface{}); ok {
		jsonData, err := json.Marshal(itemMap)
		if err == nil {
			var item AuthTempData
			err := json.Unmarshal(jsonData, &item)
			if err == nil {
				ts.TempData[saveKey] = &item
				return &item
			}
		}
	}

	// default
	defaultAuth := &AuthTempData{
		Username: "",
		Password: "",
	}
	ts.TempData[saveKey] = defaultAuth
	return defaultAuth
}

// ------- [Report]
type ReportTempData struct {
	Title            string `json:"title"`
	ReportId         int    `json:"report_id"`
	FilterId         int    `json:"filter_id"`
	LastPageNumber   int    `json:"last_page_number"`
	ReportIdSelected int    `json:"report_id_selected"`
	IsNotification   int    `json:"is_notification"`
}

func (ts *TelSession) GetReportTempData() *ReportTempData {
	saveKey := "report"
	if existItem, ok := ts.TempData[saveKey].(*ReportTempData); ok {
		return existItem
	}

	// deserialize
	if itemMap, ok := ts.TempData[saveKey].(map[string]interface{}); ok {
		jsonData, err := json.Marshal(itemMap)
		if err == nil {
			var item ReportTempData
			err := json.Unmarshal(jsonData, &item)
			if err == nil {
				ts.TempData[saveKey] = &item
				return &item
			}
		}
	}

	// default
	defaultAuth := &ReportTempData{
		Title:            "",
		ReportId:         0,
		FilterId:         0,
		LastPageNumber:   1,
		ReportIdSelected: 0,
		IsNotification:   0,
	}
	ts.TempData[saveKey] = defaultAuth
	return defaultAuth
}
