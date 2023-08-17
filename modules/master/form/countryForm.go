package form

import (
	"quince/internal/copier"
	"quince/modules/master/models"
)

//CountryForm
type CountryForm struct {
	Id      int    `form:"id"`
	Name    string `form:"name"`
	Alpha_2 string `form:"alpha_2"`
	Alpha_3 string `form:"alpha_3"`
}

func (c *CountryForm) Validate() (*models.Country, error) {
	var m models.Country
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {

		return &m, m.Validate()
	}
}
