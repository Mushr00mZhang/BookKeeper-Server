package result

import (
	"encoding/json"
	"net/http"
)

type Dto[T any] struct {
	Code   int8   `json:"code"`   // 状态码
	Result T      `json:"result"` // 结果
	Tip    string `json:"tip"`    // 提示
	Error  error  `json:"error"`  // 错误
}

// 返回结果
func (dto Dto[T]) Return(w *http.ResponseWriter, statusCode int) {
	bytes := dto.Marshal()
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(statusCode)
	(*w).Write(bytes)
}

// 序列化
func (dto Dto[T]) Marshal() []byte {
	bytes, err := json.Marshal(dto)
	if err != nil {
		dto = Dto[T]{
			Code:  -1,
			Tip:   "序列化JSON失败",
			Error: err,
		}
		bytes, _ = json.Marshal(dto)
	}
	return bytes
}
