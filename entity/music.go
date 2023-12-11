package entity

import (
	"errors"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	ISRC     string
	ImageURL string
	Title    string
	Artists  pq.StringArray `gorm:"type:varchar(255)[]"`
}

func (m *Music) Validate() error {
	if m.ISRC == "" {
		return errors.New("ISRC cannot be empty")
	}

	if m.ImageURL == "" {
		return errors.New("ImageURL cannot be empty")
	}

	if m.Title == "" {
		return errors.New("title cannot be empty")
	}

	if len(m.Artists) == 0 {
		return errors.New("artists cannot be empty")
	}

	return nil
}
