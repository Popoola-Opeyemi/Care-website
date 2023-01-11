package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Faq ...
type Faq struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// FaqListArray ...
type FaqListArray struct {
	Items []*Faq `json:"items"`
	Total int    `json:"count"`
}

// All Faq ...
func (pg *Faq) All(db *gorm.DB, limit, offset int) (ss []*Faq, count int, err error) {
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

//Create ...
func (pg *Faq) Create(db *gorm.DB) error {
	return db.Create(pg).Error
}

// GET ...
func (pg *Faq) GET(db *gorm.DB, id int) error {
	return db.First(pg, "id = ?", id).Error
}

// DELETE ...
func (pg *Faq) DELETE(db *gorm.DB) error {
	return db.Delete(pg).Error
}

// UPDATE ...
func (pg *Faq) UPDATE(db *gorm.DB) error {
	return db.Model(pg).Updates(*pg).Error
}
