package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"scd-golang-trial/models"
	"scd-golang-trial/scd"
)

func main() {
	db, err := gorm.Open(sqlite.Open("scd.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Job{}, &models.Timelog{}, &models.PaymentLineItem{})

	// ✅ Seed data (only once)
	// ⚠️ You can comment out after first run to avoid UNIQUE constraint errors
	/*
		db.Create(&models.Job{ID: "job1", Version: 1, UID: "job_uid_v1", Status: "extended", Rate: 20, Title: "SDE", CompanyID: "comp1", ContractorID: "cont1"})
		db.Create(&models.Job{ID: "job1", Version: 2, UID: "job_uid_v2", Status: "active", Rate: 22, Title: "SDE", CompanyID: "comp1", ContractorID: "cont1"})

		now := time.Now().Unix()
		db.Create(&models.Timelog{ID: "tl1", Version: 1, UID: "tl_uid_1", Duration: 1200, TimestampStart: now - 3600, TimestampEnd: now, Type: "captured", JobUID: "job_uid_v2"})

		db.Create(&models.PaymentLineItem{ID: "pl1", Version: 1, UID: "pl_uid_1", JobUID: "job_uid_v2", TimelogUID: "tl_uid_1", Amount: 264, Status: "pending"})
	*/

	// ✅ Call all 3 functions
	jobs, err := scd.GetLatestJobsForCompany(db, "comp1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ Latest Jobs for Company 'comp1':")
	for _, job := range jobs {
		fmt.Printf("%+v\n", job)

		timelogs, err := scd.GetLatestTimelogsForJob(db, job.UID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("🔹 Latest Timelogs for Job:", job.UID)
		for _, t := range timelogs {
			fmt.Printf("%+v\n", t)
		}

		items, err := scd.GetLatestPaymentLineItemsForJob(db, job.UID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("💰 Latest Payment Line Items for Job:", job.UID)
		for _, i := range items {
			fmt.Printf("%+v\n", i)
		}
	}
}
