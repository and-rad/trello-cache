package trello

// List contains cards, representing individual trello tasks.
type List struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Pos    float32 `json:"pos"`
	Closed bool    `json:"closed"`

	Cards []*Card `json:"-"`
}
