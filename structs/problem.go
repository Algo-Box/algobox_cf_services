package structs

// Problem struct represents a problem type returned bu cf
type Problem struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Rating    string   `json:"rating"`
	Tags      []string `json:"tags"`
}
