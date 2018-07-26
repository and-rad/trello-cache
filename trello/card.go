package trello

// Card is the most granular unit of trello boards.
type Card struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Pos         float32  `json:"pos"`
	Desc        string   `json:"desc"`
	DueComplete bool     `json:"dueComplete"`
	Cover       string   `json:"idAttachmentCover"`
	Checklists  []string `json:idChecklists`
	Labels      []string `json:"idLabels"`
	List        string   `json:"idList"`

	ChecklistsComplete bool `json:"-"`
}
