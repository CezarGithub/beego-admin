package form

import (
	"quince/internal/copier"
	"quince/modules/admin/models"
)

// AdminCronJobIntervalForm struct
type AdminCronJobIntervalForm struct {
	Id             int64  `form:"id"`
	AdminCronJobID int64  `form:"admin_cron_job_id"`
	DayOfWeek      int    `form:"day_of_week"`
	MonthOfYear    int    `form:"month_of_year"`
	DayOfMonth     int    `form:"day_of_month"`
	Hour           int    `form:"hour"`
	Status         int8   `form:"status"`
	RunOnce        int8   `form:"run_once"`
	Description    string `form:"description"`
}

func (c *AdminCronJobIntervalForm) Validate() (*models.AdminCronJobInterval, error) {
	var m models.AdminCronJobInterval
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		m.AdminCronJob = &models.AdminCronJob{}
		m.AdminCronJob.Id = c.AdminCronJobID
		return &m, m.Validate()
	}
}
