package database

type Project struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Category     string   `json:"category"`
	Url          string   `json:"url"`
	Technologies []string `json:"technologies"`
}
