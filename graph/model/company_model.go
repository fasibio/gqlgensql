package model

import "context"

type Company struct {
	ID              int      `json:"id" gorm:"primaryKey"`
	Name            string   `json:"Name"`
	MotherCompanyID *int     `json:"motherCompanyID"`
	MotherCompany   *Company `json:"motherCompany"`
}

func (c *Company) Description(ctx context.Context) (*string, error) {
	res := "fetch from CMS"
	return &res, nil
}
