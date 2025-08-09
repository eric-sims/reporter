// Package model represents the model of summaries
package model

import "time"

// Summary - structure of summary info - also GORM model
type Summary struct {
	ID        uint      `gorm:"primaryKey"`
	Date      time.Time `gorm:"uniqueIndex"`
	Text      string    `gorm:"type:text;not null"`
	Tags      []Tag     `gorm:"many2many:summary_tags;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Projects  []Project `gorm:"many2many:summary_projects;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Tag - custom tag for summary
// like "Innovation Week" or "Company Training"
type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;size:64;not null"`
}

// Project - focus of the task
type Project struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;size:128;not null"`
}
