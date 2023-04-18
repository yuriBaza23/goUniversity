package models

type Department struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Acronym      string `json:"acronym"`
	CordinatorID int    `json:"cordinatorId"`
	SchoolID     int    `json:"schoolId"`
}
