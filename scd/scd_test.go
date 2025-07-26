package scd

import (
	"scd-golang-trial/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetLatestJobsForCompany(t *testing.T) {
	// Setup in-memory SQLite DB for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}

	// Migrate schema
	err = db.AutoMigrate(&models.Job{})
	if err != nil {
		t.Fatalf("Failed to migrate test DB: %v", err)
	}

	// Seed jobs with two versions
	db.Create(&models.Job{
		ID:           "job1",
		Version:      1,
		UID:          "job_uid_v1",
		Status:       "extended",
		Rate:         20,
		Title:        "SDE",
		CompanyID:    "comp1",
		ContractorID: "cont1",
	})
	db.Create(&models.Job{
		ID:           "job1",
		Version:      2,
		UID:          "job_uid_v2",
		Status:       "active",
		Rate:         22,
		Title:        "SDE",
		CompanyID:    "comp1",
		ContractorID: "cont1",
	})

	// Call the function
	jobs, err := GetLatestJobsForCompany(db, "comp1")
	if err != nil {
		t.Fatalf("Error fetching jobs: %v", err)
	}

	if len(jobs) != 1 {
		t.Errorf("Expected 1 job, got %d", len(jobs))
	}
	if jobs[0].Version != 2 {
		t.Errorf("Expected version 2, got %d", jobs[0].Version)
	}
}
