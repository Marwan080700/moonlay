package database

import (
	"fmt"
	"moonlay/models"
	"moonlay/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.DataList{},
		&models.DataSublist{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration success")
}