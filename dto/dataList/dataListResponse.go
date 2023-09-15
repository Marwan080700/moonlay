package datalist

type DataListResponse struct {
	Title       string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Description string `json:"description" form:"description" gorm:"type: varchar(255)"`
	File        string `json:"file" form:"file" gorm:"type: varchar(255)"`
}