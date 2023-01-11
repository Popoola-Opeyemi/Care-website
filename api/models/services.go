package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

// Services ...
type Services struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	URL       string          `json:"url"`
	UUID      string          `json:"uuid"`
	Content   string          `json:"content"`
	Image     json.RawMessage `json:"image"`
	CreatedAt time.Time       `json:"created_at"`
}

type ServicesListArray struct {
	Items []*Services `json:"items"`
	Total int         `json:"count"`
}

//  All...
func (service *Services) All(db *gorm.DB, limit, offset int) (ss []*Services, count int, err error) {
	err = db.
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&ss).
		Limit(-1).
		Offset(-1).
		Count(&count).Error
	return
}

// GETALL
func (service *Services) GETALL(db *gorm.DB) (serviceList []*Services, err error) {
	err = db.Find(&serviceList).Error
	return
}

// Create a new record
func (service *Services) Create(db *gorm.DB) error {
	return db.Create(service).Error
}

// GET ...
func (service *Services) GETBYSTRING(db *gorm.DB, id string) error {
	return db.First(service, "url = ?", id).Error
}

// GET ...
func (service *Services) GETBYUUID(db *gorm.DB, id string) error {
	return db.First(service, "uuid = ?", id).Error
}

// GET ...
func (service *Services) GET(db *gorm.DB, id int) error {
	return db.First(service, "id = ?", id).Error
}

// DELETE ...
func (service *Services) DELETE(db *gorm.DB) error {
	return db.Delete(service).Error
}

// UPDATE ...
func (service *Services) UPDATE(db *gorm.DB) error {
	return db.Model(service).Updates(*service).Error
}
