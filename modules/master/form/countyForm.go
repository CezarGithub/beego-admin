package form

import (
	"quince/internal/copier"
	"quince/modules/master/models"
)

// CountyForm
type CountyForm struct {
	Id        int64  `form:"id"`
	Name      string `form:"name"`
	Code      string `form:"code"`
	CountryID int64  `form:"country_id"`
}

func (c *CountyForm) Validate() (*models.County, error) {
	var m models.County
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		m.Country = &models.Country{}
		m.Country.Id = c.CountryID
		return &m, m.Validate()
	}
}
