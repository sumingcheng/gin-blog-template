package util

import (
	"reflect"
	"testing"
)

func TestCamel2Snake(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"UserId", "user_id"},
		{"ID", "i_d"},
		{"createdAt", "created_at"},
		{"Name", "name"},
		{"HTMLParser", "h_t_m_l_parser"},
		{"", ""},
		{"a", "a"},
		{"A", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Camel2Snake(tt.input)
			if got != tt.want {
				t.Errorf("Camel2Snake(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestGetGormFields(t *testing.T) {
	type Sample struct {
		Id       int    `gorm:"column:id;primaryKey"`
		Name     string `gorm:"column:name"`
		PassWd   string `gorm:"column:password"`
		NoTag    string
		Ignored  string `gorm:"-"`
		unexport string //nolint
	}

	got := GetGormFields(Sample{})
	want := []string{"id", "name", "password", "no_tag"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetGormFields = %v, want %v", got, want)
	}
	_ = Sample{}.unexport
}

func TestGetGormFields_Pointer(t *testing.T) {
	type S struct {
		Id int `gorm:"column:id"`
	}
	got := GetGormFields(&S{})
	if len(got) != 1 || got[0] != "id" {
		t.Errorf("GetGormFields pointer = %v, want [id]", got)
	}
}

func TestGetGormFields_NonStruct(t *testing.T) {
	got := GetGormFields("not a struct")
	if got != nil {
		t.Errorf("GetGormFields non-struct = %v, want nil", got)
	}
}
