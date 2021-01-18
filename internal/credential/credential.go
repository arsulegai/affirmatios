package credential

// Credential that is issued by the Hospital
type Credential struct {
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Place   string `json:"place"`
	Date    string `json:"date"`
}
