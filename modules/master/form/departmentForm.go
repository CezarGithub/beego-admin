                    package form


import (
	"quince/internal/copier"
	"quince/modules/master/models"
)

// DepartmentForm struct
type DepartmentForm struct {
		Id           	int64  `form:"id"`
		CompanyID		int64 `form:"company_id"`
		Name			string  `form:"name"`
		Code			string  `form:"code"`
}

func (c *DepartmentForm) Validate() (*models.Department, error) {
	var m models.Department
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
			m.Company = &models.Company{}
			m.Company.Id = c.CompanyID
		return &m, m.Validate()
	}
}
                