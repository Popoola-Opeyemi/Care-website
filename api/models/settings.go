package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

//Settings ...
type Settings struct {
	ID         int             `json:"id"`
	UUID       string          `json:"uuid"`
	Name       string          `json:"name"`
	URL        string          `json:"route"`
	Attributes json.RawMessage `json:"attributes"`
	CreatedAt  time.Time       `json:"created_at"`
}

//Create ...
func (pg *Settings) Create(db *gorm.DB) error {
	return db.Create(pg).Error
}

// All ... gets all new record
func (pg *Settings) All(db *gorm.DB) (Settings []*Settings, err error) {
	err = db.Find(&Settings).Error
	return
}

// GET ...
func (pg *Settings) GET(db *gorm.DB, id int) error {
	return db.First(pg, "id = ?", id).Error
}

func (pg *Settings) GETBYUUID(db *gorm.DB, id string) error {
	return db.First(pg, "uuid = ?", id).Error
}

// DELETE ...
func (pg *Settings) DELETE(db *gorm.DB) error {
	return db.Delete(pg).Error
}

// UPDATE ...
func (pg *Settings) UPDATE(db *gorm.DB) error {
	return db.Model(pg).Updates(*pg).Error
}
