package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Contacts ...
type Contacts struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	MobileNumber string    `json:"mobile_number"`
	Message      string    `json:"message"`
	CreatedAt    time.Time `json:"created_at"`
}

// ContactsListArray ...
type ContactsListArray struct {
	Items []*Contacts `json:"items"`
	Total int         `json:"count"`
}

// All Contacts ...
func (pg *Contacts) All(db *gorm.DB, limit, offset int) (ss []*Contacts, count int, err error) {
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
func (pg *Contacts) Create(db *gorm.DB) error {
	return db.Create(pg).Error
}

// GET ...
func (pg *Contacts) GET(db *gorm.DB, id int) error {
	return db.First(pg, "id = ?", id).Error
}

// DELETE ...
func (pg *Contacts) DELETE(db *gorm.DB) error {
	return db.Delete(pg).Error
}

// UPDATE ...
func (pg *Contacts) UPDATE(db *gorm.DB) error {
	return db.Model(pg).Updates(*pg).Error
}
