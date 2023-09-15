package models

import "time"

type DataList struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" gorm:"type: varchar(255)" form:"title"`
	Description	 string 	`json:"description" gorm:"type: varchar(255)" form:"description"`
	File string `json:"file" gorm:"type: varchar(255)" form:"file"`
	Sublist  []DataSublist `json:"sublist" gorm:"foreignKey:DataListID"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (DataList) TableName() string {
    return "data_lists"
}
