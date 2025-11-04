package core

import (
	"strings"
)

type Paging struct {
	Page    int    `json:"page" form:"page"`
	Limit   int    `json:"limit" form:"limit"`
	Total   int    `json:"total" form:"total"`
	SortBy  string `json:"sortBy,omitempty" form:"sortBy"`
	SortDir string `json:"sortDir,omitempty" form:"sortDir"`
}

// Fulfill chuẩn hóa giá trị paging + sort
func (p *Paging) Fulfill() {
	// Mặc định trang đầu tiên
	if p.Page <= 0 {
		p.Page = 1
	}

	// Mặc định limit = 50, giới hạn tối đa 100
	if p.Limit <= 0 {
		p.Limit = 50
	} else if p.Limit > 100 {
		p.Limit = 100
	}

	// Mặc định sortBy
	if p.SortBy == "" {
		p.SortBy = "id"
	}

	// Chuẩn hóa sortDir
	p.SortDir = strings.TrimSpace(strings.ToLower(p.SortDir))
	if p.SortDir != "asc" && p.SortDir != "desc" {
		p.SortDir = "desc"
	}
}

// ValidateSortFields kiểm tra sortBy có hợp lệ không
func (p *Paging) ValidateSortFields(allowed []string) {
	valid := false
	for _, f := range allowed {
		if f == p.SortBy {
			valid = true
			break
		}
	}
	if !valid && len(allowed) > 0 {
		p.SortBy = allowed[0]
	}
}
