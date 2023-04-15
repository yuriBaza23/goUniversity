package entities

type Subject struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Acronym      string `json:"acronym"`
	DepartmentID int    `json:"departmentId"`
	Obrigatory   bool   `json:"obrigatory"`
}
