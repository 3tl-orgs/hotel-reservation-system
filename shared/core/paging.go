package core

type Paging struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
	Total int `json:"total" form:"total"`
	// Support cursor with UID
	//FakeCursor string `json:"cursor" form:"cursor"`
	//NextCursor string `json:"next_cursor" form:"next_cursor"`
	SortBy  string `json:"sortBy,omitempty" form:"sortBy"`
	SortDir string `json:"sortDir,omitempty" form:"sortDir"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}

	if p.SortBy == "" {
		p.SortBy = "id"
	}

	if p.SortDir == "" {
		p.SortDir = "desc"
	}

}
