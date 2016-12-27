package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DatabaseRef struct {
	Active *gorm.DB
}

func initDatabase(filepath string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Report{})
	db.AutoMigrate(&Tag{})

	return db, nil
}

type Tag struct {
	ID   uint   `gorm:"primary_key"`
	Text string `gorm:"text;type:varchar(255);unique_index"`
}

type Report struct {
	gorm.Model
	Ulid            string    `gorm:"type:varchar(100);unique_index"`
	Name            string    `json:"name"`
	Command         string    `json:"-"`
	RawCommand      []string  `gorm:"-" json:"command"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	ElapsedSeconds  float32   `json:"elapsed_seconds"`
	ExitCode        int       `json:"exit_code"`
	ExitDescription string    `json:"exit_description"`
	CapturedOutput  string    `json:"captured_output"`
	Hostname        string    `json:"hostname"`
	Tags            []Tag     `json:"-" gorm:"many2many:report_tags;"`
	RawTags         []string  `gorm:"-" json:"tags"`
}
