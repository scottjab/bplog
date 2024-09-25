package bplog

import (
	"encoding/csv"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

type BloodPressure struct {
	gorm.Model
	Systolic  int
	Diastolic int
	Pulse     int
	Date      time.Time
	Note      string
}

func (b *BPLog) DBMigrate() {
	err := b.db.AutoMigrate(&BloodPressure{})
	if err != nil {
		log.Fatalf("error migrating blood pressure table: %v", err)
	}
}

func (b *BPLog) LoadCSVIntoDB(csvFile string) {
	csvfile, err := os.Open(csvFile)
	if err != nil {
		log.Fatalf("error opening csv file: %v", err)
	}
	defer csvfile.Close()
	csvReader := csv.NewReader(csvfile)

	// Skip header
	_, err = csvReader.Read()
	if err != nil {
		log.Fatalf("error reading csv file: %v", err)
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("error reading csv file: %v", err)
	}
	for _, record := range records {
		log.Printf("loading blood pressure: %v", record)

		systolic, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalf("error parsing systolic: %v", err)
		}
		diastolic, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatalf("error parsing diastolic: %v", err)
		}
		pulse, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatalf("error parsing pulse: %v", err)
		}

		date, err := time.Parse("Jan 02 2006 03:04:05 PM -07:00", record[3])
		if err != nil {
			log.Fatalf("error parsing date: %v", err)
		}
		b.db.Create(&BloodPressure{
			Systolic:  systolic,
			Diastolic: diastolic,
			Pulse:     pulse,
			Date:      date,
			Note:      record[4],
		})
	}

}
