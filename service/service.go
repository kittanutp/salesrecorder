package service

import "gorm.io/gorm"

// create database controller
type DBController struct {
	Database *gorm.DB
}
