package model

import "testing"

func TestPageQuery_Normalize(t *testing.T) {
	tests := []struct {
		name     string
		input    PageQuery
		wantPage int
		wantSize int
	}{
		{"默认值", PageQuery{}, 1, 10},
		{"负数", PageQuery{Page: -1, Size: -5}, 1, 10},
		{"正常值", PageQuery{Page: 3, Size: 20}, 3, 20},
		{"零值", PageQuery{Page: 0, Size: 0}, 1, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.input.Normalize()
			if tt.input.Page != tt.wantPage {
				t.Errorf("Page = %d, want %d", tt.input.Page, tt.wantPage)
			}
			if tt.input.Size != tt.wantSize {
				t.Errorf("Size = %d, want %d", tt.input.Size, tt.wantSize)
			}
		})
	}
}

func TestPageQuery_Offset(t *testing.T) {
	tests := []struct {
		page, size, want int
	}{
		{1, 10, 0},
		{2, 10, 10},
		{3, 20, 40},
	}
	for _, tt := range tests {
		q := PageQuery{Page: tt.page, Size: tt.size}
		if got := q.Offset(); got != tt.want {
			t.Errorf("Offset() page=%d size=%d = %d, want %d", tt.page, tt.size, got, tt.want)
		}
	}
}
