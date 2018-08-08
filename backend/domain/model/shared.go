package model

import "time"

// Category struct
type Category struct {
	Category string `gorm:"primary_key"`
	Type     string `gorm:"primary_key"`
	Name     string
	Desc     string
}

// Calendar struct
type Calendar struct {
	ID              string `gorm:"primary_key"`
	SameAs          string
	Context         string
	Type            string
	Date            *time.Time `gorm:"type:datetime"`
	Title           string
	CalendarTitleJa string
	CalendarTitleEn string
	Duration        string
	UpdatedAt       time.Time `gorm:"type:datetime" json:"-"`
}

// CalendarDay struct
type CalendarDay struct {
	CalendarID string    `gorm:"primary_key"`
	Day        string    `gorm:"primary_key"`
	UpdatedAt  time.Time `gorm:"type:datetime" json:"-"`
}

// Operator struct
type Operator struct {
	ID              string `gorm:"primary_key"`
	SameAs          string
	Context         string
	Type            string
	Date            *time.Time `gorm:"type:datetime"`
	Title           string
	OperatorTitleJa string
	OperatorTitleEn string
	UpdatedAt       time.Time `gorm:"type:datetime" json:"-"`
}
