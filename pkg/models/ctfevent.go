package models

import (
	"time"

	"gorm.io/gorm"
)

type CTFEvent struct {
	gorm.Model
	Organizer   string
	CTFTimeURL  string
	CTFID       int
	Weight      float32
	Duration    time.Duration
	Logo        string
	Title       string
	Start       time.Time
	Finish      time.Time
	Description string
	Format      string
	URL         string
}

type CTFEvents []CTFEvent
