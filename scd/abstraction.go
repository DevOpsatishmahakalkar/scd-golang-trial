package scd

import (
	"scd-golang-trial/models"

	"gorm.io/gorm"
)

// Get latest jobs for a company
func GetLatestJobsForCompany(db *gorm.DB, companyID string) ([]models.Job, error) {
	var jobs []models.Job

	sub := db.Model(&models.Job{}).
		Select("id, MAX(version) as version").
		Where("company_id = ?", companyID).
		Group("id")

	err := db.Model(&models.Job{}).
		Joins("JOIN (?) as latest ON jobs.id = latest.id AND jobs.version = latest.version", sub).
		Where("company_id = ?", companyID).
		Find(&jobs).Error

	return jobs, err
}

// Get latest timelogs for a job UID
func GetLatestTimelogsForJob(db *gorm.DB, jobUID string) ([]models.Timelog, error) {
	var timelogs []models.Timelog

	sub := db.Model(&models.Timelog{}).
		Select("id, MAX(version) as version").
		Where("job_uid = ?", jobUID).
		Group("id")

	err := db.Model(&models.Timelog{}).
		Joins("JOIN (?) as latest ON timelogs.id = latest.id AND timelogs.version = latest.version", sub).
		Where("job_uid = ?", jobUID).
		Find(&timelogs).Error

	return timelogs, err
}

// Get latest payment line items for a job UID
func GetLatestPaymentLineItemsForJob(db *gorm.DB, jobUID string) ([]models.PaymentLineItem, error) {
	var items []models.PaymentLineItem

	sub := db.Model(&models.PaymentLineItem{}).
		Select("id, MAX(version) as version").
		Where("job_uid = ?", jobUID).
		Group("id")

	err := db.Model(&models.PaymentLineItem{}).
		Joins("JOIN (?) as latest ON payment_line_items.id = latest.id AND payment_line_items.version = latest.version", sub).
		Where("job_uid = ?", jobUID).
		Find(&items).Error

	return items, err
}
