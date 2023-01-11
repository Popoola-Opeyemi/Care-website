package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Testimonials ...
type Testimonials struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// FaqListArray ...
type TestimonialsListArray struct {
	Items []*Testimonials `json:"items"`
	Total int             `json:"count"`
}

// All Testimonials ...
func (pg *Testimonials) All(db *gorm.DB, limit, offset int) (ss []*Testimonials, count int, err error) {
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
func (pg *Testimonials) Create(db *gorm.DB) error {
	return db.Create(pg).Error
}

// GET ...
func (pg *Testimonials) GET(db *gorm.DB, id int) error {
	return db.First(pg, "id = ?", id).Error
}

// DELETE ...
func (pg *Testimonials) DELETE(db *gorm.DB) error {
	return db.Delete(pg).Error
}

// UPDATE ...
func (pg *Testimonials) UPDATE(db *gorm.DB) error {
	return db.Model(pg).Updates(*pg).Error
}
