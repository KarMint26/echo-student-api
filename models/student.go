package models

import "gorm.io/gorm"

type Students struct {
	ID		uint		`gorm:"primary key;autoIncrement" json:"id"`
	Name	string		`gorm:"type:varchar(55)" json:"name"`
	Age		uint		`gorm:"type:smallint" json:"age"`
	Grade	string		`gorm:"type:char(2)" json:"grade"`
}

func MigrateStudents(DB *gorm.DB) error {
	err := DB.AutoMigrate(&Students{})
	return err
}