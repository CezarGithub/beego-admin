package form

import (
	"quince/internal/copier"
	"quince/modules/master/models"
)

// CompanyForm
type CompanyForm struct {
	Id          int64  `form:"id"`
	Name        string `form:"name"`
	Code        string `form:"code"`
	VAT         string `form:"vat"`
	RegNo       string `form:"regno"`
	Mark        string `form:"mark"`
	Web         string `form:"web"`
	Email       string `form:"email"`
	Description string `form:"description"`
	Status      int8   `form:"status"`
	CountryID   int64  `form:"country_id"`
	GroupID     int64  `form:"group_id"`
}

func (c *CompanyForm) Validate() (*models.Company, error) {
	var m models.Company
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		m.Country = &models.Country{}
		m.Country.Id = c.CountryID
		m.Group = &models.Group{}
		m.Group.Id = c.GroupID
		return &m, m.Validate()
	}
}
