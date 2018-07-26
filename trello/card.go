package trello

// Card is the most granular unit of trello boards.
type Card struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Pos         float32       `json:"pos"`
	Desc        string        `json:"desc"`
	Cover       string        `json:"idAttachmentCover"`
	Labels      []string      `json:"idLabels"`
	List        string        `json:"idList"`
	Badges      *Badges       `json:"badges"`
	Attachments []*Attachment `json:"attachments"`
}

type Badges struct {
	CheckItems        uint16 `json:"checkItems"`
	CheckItemsChecked uint16 `json:"checkItemsChecked"`
	DueComplete       bool   `json:"dueComplete"`
}

type Attachment struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	URL      string     `json:"url"`
	Previews []*Preview `json:"previews"`
}

type Preview struct {
	ID     string `json:"_id"`
	Height uint16 `json:"height"`
	Width  uint16 `json:"width"`
	URL    string `json:"url"`
	Scaled bool   `json:"scaled"`
}
