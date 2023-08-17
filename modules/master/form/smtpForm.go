package form

import (
	"quince/internal/copier"
	"quince/modules/master/models"
)

// CountyForm
type SMTPForm struct {
	Id        int64  `form:"id"`
	Server    string `form:"server"`
	Port      string `form:"port"`
	User      string `form:"user"`
	Password  string `form:"password"`
	UseTLS    string `form:"useTLS"`
	CompanyID int64  `form:"company_id"`
}

func (c *SMTPForm) Validate() (*models.SMTP, error) {
	var m models.SMTP
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		m.Company = &models.Company{}
		m.Company.Id = c.CompanyID
		return &m, m.Validate()
	}
}
