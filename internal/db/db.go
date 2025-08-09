// Package db represents an interface for the SQLite database
package db

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/eric-sims/reporter/internal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// DB - SQLite database
type DB struct{ g *gorm.DB }

// Open - opends db connection
func Open() (*DB, error) {
	path, err := defaultPath()
	if err != nil {
		return nil, err
	}
	if mkErr := os.MkdirAll(filepath.Dir(path), 0o755); mkErr != nil {
		return nil, mkErr
	}
	g, gErr := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if gErr != nil {
		return nil, gErr
	}
	if miErr := g.AutoMigrate(&model.Summary{}); miErr != nil {
		return nil, miErr
	}
	return &DB{g: g}, nil
}

// Close - close database
func (d *DB) Close() error { return nil }

func defaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "reporter", "data.db"), nil
}

// UpsertSummary - upserts summary
func (d *DB) UpsertSummary(s *model.Summary) error {
	if s.Date.IsZero() {
		return errors.New("summary date required")
	}
	// Always create a new summary entry (allow multiple entries per day).
	return d.g.Create(s).Error
}

// ListSummaries - Lists summaries
func (d *DB) ListSummaries(start, end time.Time) ([]model.Summary, error) {
	var out []model.Summary
	q := d.g.Order("date asc, id asc").Where("date BETWEEN ? AND ?", start.Truncate(24*time.Hour), end)
	return out, q.Find(&out).Error
}
