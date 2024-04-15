package serializer

type PageAo struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
}

type PageVo struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

// BuildPage 分页
func BuildPage(data interface{}, total int64) Response {
	return Response{
		Data: PageVo{
			Total: int(total),
			Data:  data,
		},
		Code: 0,
		Msg:  "查询成功",
	}
}
