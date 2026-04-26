package model

// PageQuery 分页请求参数
type PageQuery struct {
	Page    int    `form:"page" binding:"omitempty,min=1"`
	Size    int    `form:"size" binding:"omitempty,min=1,max=100"`
	Keyword string `form:"keyword"`
}

// PageResult 分页响应
type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
}

func (q *PageQuery) Normalize() {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.Size <= 0 {
		q.Size = 10
	}
}

func (q *PageQuery) Offset() int {
	return (q.Page - 1) * q.Size
}
