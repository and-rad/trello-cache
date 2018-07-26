package trello

// Board contains lists. A Board represents a project.
type Board struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Lists []*List `json:"lists"`
	Cards []*Card `json:"cards"`
}
