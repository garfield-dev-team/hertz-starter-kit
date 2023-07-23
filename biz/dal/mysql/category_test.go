package mysql

import (
	"encoding/json"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	type args struct {
		category *Category
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "Add category 1",
			args: args{category: &Category{Name: "Java 相关"}},
		},
		{
			name: "Add category 2",
			args: args{category: &Category{Name: "Golang 相关"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateCategory(tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("CreateCategory() success with result: %v", got)
		})
	}
}

func TestDeleteCategoryById(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Delete category 1",
			args: args{id: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteCategoryById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCategoryById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryCategories(t *testing.T) {
	tests := []struct {
		name    string
		want    []*Category
		wantErr bool
	}{
		{
			name: "Get all categories",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryCategories()
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			res, err := json.MarshalIndent(got, "", "  ")
			t.Logf("QueryCategories() success with result\n %s", string(res))
		})
	}
}

func TestQueryCategoryById(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    *Category
		wantErr bool
	}{
		{
			name: "Get category 3",
			args: args{id: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryCategoryById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCategoryById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			res, err := json.MarshalIndent(got, "", "  ")
			t.Logf("QueryCategoryById() success with result\n %s", string(res))
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	type args struct {
		category *Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Edit category 1",
			args: args{category: &Category{BaseModel: &BaseModel{ID: 1}, Name: "Java 相关分类"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateCategory(tt.args.category); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
