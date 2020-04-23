package structs

// DashboardData struct
type DashboardData struct {
	ByIndex  map[string]int `json:"byIndex"`
	ByRating map[int]int    `json:"byRating"`
	ByTags   map[string]int `json:"byTags"`
}
