package main

import (
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	Product     string
	Acquired    uint
	Available   int
	Safety      uint
	Price       float32
	Overbooking bool
}

func (Stock *Stock) create() {
	db.Create(Stock)
}

func (Stock *Stock) findById(id uint) {
	db.First(Stock, id)
}

func (Stock *Stock) update() {
	db.Save(Stock)
}
