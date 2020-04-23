package structs

// UserStatus is a struct for user status
type UserStatus struct {
	Status  string       `json:"status"`
	Result  []Submission `json:"result"`
	Comment string       `json:"comment"`
}
