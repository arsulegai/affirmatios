package credential

// Credential that is issued by the University
type Credential struct {
	Name          string `json:"name"`
	EmployeeId    string `json:"employee_id"`
	JoiningDate   string `json:"joining_date"`
	Department    string `json:"department"`
	Role          string `json:"role"`
	IssuedDate    string `json:"issued_date"`
	RelievingDate string `json:"relieving_date"`
}
