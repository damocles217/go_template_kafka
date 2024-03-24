package services

import "gorm.io/gorm"

type Services struct {
	DB *gorm.DB
}

func NewServicesWithDb(db *gorm.DB) *Services {
	return &Services{
		DB: db,
	}
}

func (s *Services) GetUser() int {
	return 0
}
