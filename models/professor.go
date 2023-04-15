package models

type Professor struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SubjectID    int    `json:"subjectId"`
	DepartmentID int    `json:"departmentId"`
}
