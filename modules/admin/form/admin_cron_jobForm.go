package form

import (
	"quince/internal/copier"
	"quince/modules/admin/models"
)

// AdminCronJobForm
type AdminCronJobForm struct {
	Id          int    `form:"id"`
	Name        string `form:"name"`
	Module      string `form:"module"`
	Description string `form:"description"`
	Status      int8   `form:"status"`
}

func (c *AdminCronJobForm) Validate() (*models.AdminCronJob, error) {
	var m models.AdminCronJob
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		return &m, m.Validate()
	}
}
