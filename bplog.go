package bplog

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type BPLog struct {
	db *gorm.DB
}

func NewBPLog() *BPLog {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return &BPLog{
		db: db,
	}
}
