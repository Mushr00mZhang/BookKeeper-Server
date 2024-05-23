package pagedlist

import (
	"net/url"
	"strconv"
)

// 获取分页列表Dto
type GetDto struct {
	Index int `json:"index"` // 序号
	Size  int `json:"size"`  // 大小
}

// 分页列表Dto
type Dto[T any] struct {
	Total int `json:"total"` // 总数
	Items []T `json:"items"` //列表
}

func ParseGetDto(values url.Values) GetDto {
	dto := GetDto{}
	size, err := strconv.ParseInt(values.Get("size"), 10, 32)
	if err == nil {
		dto.Size = int(size)
	}
	index, err := strconv.ParseInt(values.Get("index"), 10, 32)
	if err == nil {
		dto.Index = int(index)
	}
	return dto
}
func Empty[T any]() Dto[T] {
	return Dto[T]{
		Total: 0,
		Items: make([]T, 0),
	}
}
