package structs

// Submission Struct
type Submission struct {
	ID      int64   `json:"id"`
	Lang    string  `json:"programmingLanguage"`
	Verdict string  `json:"verdict"`
	Problem Problem `json:"problem"`
}
