package model

// PaginationRequest 分页请求
type PaginationRequest struct {
	PageNo   int `json:"pageNo" form:"pageNo" binding:"min=1" default:"1"`
	PageSize int `json:"pageSize" form:"pageSize" binding:"min=1,max=100" default:"10"`
}

// SetDefault 设置默认值
func (p *PaginationRequest) SetDefault() {
	if p.PageNo < 1 {
		p.PageNo = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
}

// GetOffset 计算分页偏移量
func (p *PaginationRequest) GetOffset() int {
	return (p.PageNo - 1) * p.PageSize
}

// PaginationResponse 分页响应
type PaginationResponse struct {
	Total    int64       `json:"total"`    // 总记录数
	PageNo   int         `json:"pageNo"`   // 当前页码
	PageSize int         `json:"pageSize"` // 每页记录数
	Pages    int64       `json:"pages"`    // 总页数
	Data     interface{} `json:"data"`
}
