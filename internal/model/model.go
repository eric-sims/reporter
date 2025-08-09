// Package model represents the model of summaries
package model

import "time"

// Summary - structure of summary info - also GORM model
type Summary struct {
	ID        uint      `gorm:"primaryKey"`
	Date      time.Time `gorm:"index;not null"`
	Text      string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}
