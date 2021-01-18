package credential

// Credential that is issued by the University
type Credential struct {
	Name          string `json:"name"`
	RollNumber    string `json:"roll_number"`
	CompletedDate string `json:"completed_date"`
	Department    string `json:"department"`
	Address       string `json:"address"`
	IssuedDate    string `json:"issued_date"`
}
