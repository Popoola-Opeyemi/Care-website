package models

import (
	"encoding/json"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/jinzhu/gorm"
)

// Pages ...
type Pages struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Route      string          `json:"route"`
	UUID       string          `json:"uuid"`
	Attributes json.RawMessage `json:"attributes"`
	CreatedAt  time.Time       `json:"created_at"`
}

// PagesListArray ...
type PagesListArray struct {
	Items []*Pages `json:"items"`
	Total int      `json:"count"`
}

// Create a new record
func (pg *Pages) Create(db *gorm.DB) error {
	return db.Create(pg).Error
}

// All pages ...
func (s *Pages) All(db *gorm.DB, limit, offset int) (ss []*Pages, count int, err error) {
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

// GET ...
func (pg *Pages) GETBYUUID(db *gorm.DB, id string) error {
	return db.First(pg, "uuid = ?", id).Error
}

// GET ...
func (pg *Pages) GET(db *gorm.DB, id int) error {
	return db.First(pg, "id = ?", id).Error
}

// DELETE ...
func (pg *Pages) DELETE(db *gorm.DB) error {
	return db.Delete(pg).Error
}

// UPDATE ...
func (pg *Pages) UPDATE(db *gorm.DB) error {
	return db.Model(pg).Updates(*pg).Error
}

//AttributesImgHandler ...
func AttributesImgHandler(attr json.RawMessage) (e error) {
	var attrmap interface{}
	err := json.Unmarshal(attr, &attrmap)
	if err != nil {
		log.Debug(err)
	}
	// log.Info(&attrmap)
	return

}
