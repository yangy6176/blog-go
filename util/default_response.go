package util

/*
@TODO 通用返回数据类型
@Author 江南小魏晨
*/
type DefaultResponse struct {
	Data    interface{} `json:"data"`
	Total   uint64      `json:"total"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

// 自定义响应信息
func (res *DefaultResponse) Get(data DefaultResponse) DefaultResponse {
	return data
}
