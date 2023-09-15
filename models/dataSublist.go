package models

type DataSublist struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	DataListID  int    `json:"datalist_id"`
	Title       string `json:"title" gorm:"type: varchar(255)" form:"title"`
	Description string `json:"description" gorm:"type: varchar(255)" form:"description"`
	File        string `json:"file" gorm:"type: varchar(255)" form:"file"`
}

func (DataSublist) TableName() string {
	return "data_sublists"
}