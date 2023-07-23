package mysql

import (
	"encoding/json"
	"testing"
)

func TestCreateArticle(t *testing.T) {
	type args struct {
		article *Article
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name:    "Add article 1",
			args:    args{article: &Article{Title: "测试文章2333", Desc: "文章描述2333", UserID: 1}},
			wantErr: false,
		},
		{
			name:    "Add article 2",
			args:    args{article: &Article{Title: "测试文章666", Desc: "文章描述666", UserID: 2}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateArticle(tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("CreateArticle() success with result: %v", got)
		})
	}
}

func TestDeleteArticleById(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Delete article 1",
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteArticleById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteArticleById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryArticleById(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    *Article
		wantErr bool
	}{
		{
			name:    "Get article",
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryArticleById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryArticleById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			res, err := json.MarshalIndent(got, "", "  ")
			t.Logf("QueryArticleById() success with result\n %s", string(res))
		})
	}
}

func TestQueryArticles(t *testing.T) {
	type args struct {
		pageNum  int
		pageSize int
	}
	tests := []struct {
		name    string
		args    args
		want    []*Article
		wantErr bool
	}{
		{
			name:    "Get article list",
			args:    args{pageNum: 1, pageSize: 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryArticles(tt.args.pageNum, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			res, err := json.MarshalIndent(got, "", "  ")
			t.Logf("QueryArticles() success with result\n %s", string(res))
		})
	}
}

func TestUpdateArticle(t *testing.T) {
	type args struct {
		article *Article
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Edit article 1",
			args: args{
				article: &Article{
					BaseModel: &BaseModel{ID: 1},
					Desc:      "更新234234文章描述666",
					// 关联新分类，会自动在 category 表插入数据
					Categories: []*Category{
						{Name: "TypeScript 相关"},
						{Name: "Python 相关"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Edit article 2",
			args: args{
				article: &Article{
					BaseModel: &BaseModel{ID: 2},
					Desc:      "更新文章描述234223234234",
					// 关联已有分类，只需要传入 category 表的 ID 字段
					Categories: []*Category{
						{BaseModel: &BaseModel{ID: 1}},
						{BaseModel: &BaseModel{ID: 2}},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateArticle(tt.args.article); (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
