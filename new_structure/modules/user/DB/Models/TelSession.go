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
	defaultItem := &GeneralTempData{
		LastMessage: "",
	}
	ts.TempData[saveKey] = defaultItem
	return defaultItem
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
	defaultItem := &AuthTempData{
		Username: "",
		Password: "",
	}
	ts.TempData[saveKey] = defaultItem
	return defaultItem
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
	defaultItem := &ReportTempData{
		Title:            "",
		ReportId:         0,
		FilterId:         0,
		LastPageNumber:   1,
		ReportIdSelected: 0,
		IsNotification:   0,
	}
	ts.TempData[saveKey] = defaultItem
	return defaultItem
}

// ------- [Post]
type PostTempData struct {
	PostId         int `json:"post_id"`
	LastPageNumber int `json:"last_page_number"`
	FilterId       int `json:"filter_id"`
}

func (ts *TelSession) GetPostTempData() *PostTempData {
	saveKey := "post"
	if existItem, ok := ts.TempData[saveKey].(*PostTempData); ok {
		return existItem
	}

	// deserialize
	if itemMap, ok := ts.TempData[saveKey].(map[string]interface{}); ok {
		jsonData, err := json.Marshal(itemMap)
		if err == nil {
			var item PostTempData
			err := json.Unmarshal(jsonData, &item)
			if err == nil {
				ts.TempData[saveKey] = &item
				return &item
			}
		}
	}

	// default
	defaultItem := &PostTempData{
		PostId:         0,
		FilterId:       0,
		LastPageNumber: 1,
	}
	ts.TempData[saveKey] = defaultItem
	return defaultItem
}

// ------- [Monitoring]
type MonitoringTempData struct {
	LastPageNumber int `json:"last_page_number"`
}

func (ts *TelSession) GetMonitoringTempData() *MonitoringTempData {
	saveKey := "monitoring"
	if existItem, ok := ts.TempData[saveKey].(*MonitoringTempData); ok {
		return existItem
	}

	// deserialize
	if itemMap, ok := ts.TempData[saveKey].(map[string]interface{}); ok {
		jsonData, err := json.Marshal(itemMap)
		if err == nil {
			var item MonitoringTempData
			err := json.Unmarshal(jsonData, &item)
			if err == nil {
				ts.TempData[saveKey] = &item
				return &item
			}
		}
	}

	// default
	defaultItem := &MonitoringTempData{
		LastPageNumber: 1,
	}
	ts.TempData[saveKey] = defaultItem
	return defaultItem
}

// ------- [Export]
type ExportTempData struct {
	LastPageNumber int   `json:"last_page_number"`
	ExportIds      []int `json:"export_ids"`
}

func (ts *TelSession) GetExportTempData() *ExportTempData {
	saveKey := "export"
	if existItem, ok := ts.TempData[saveKey].(*ExportTempData); ok {
		return existItem
	}

	// deserialize
	if itemMap, ok := ts.TempData[saveKey].(map[string]interface{}); ok {
		jsonData, err := json.Marshal(itemMap)
		if err == nil {
			var item ExportTempData
			err := json.Unmarshal(jsonData, &item)
			if err == nil {
				ts.TempData[saveKey] = &item
				return &item
			}
		}
	}

	// default
	defaultItem := &ExportTempData{
		LastPageNumber: 1,
	}
	ts.TempData[saveKey] = defaultItem
	return defaultItem
}
