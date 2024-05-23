package outlaycat

import (
	outlaycat_base "financialrecord-backend/bases/outlaycat"
	"financialrecord-backend/dtos/pagedlist"
	"financialrecord-backend/models/outlaycat"
	"net/url"

	"github.com/google/uuid"
)

// 支出类型Dto
type Dto = outlaycat.OutlayCat

// 创建支出Dto
type CreateDto = outlaycat_base.Props

// 更新支出Dto
type UpdateDto = outlaycat_base.Base

// 获取支出列表Dto
type ListDto struct {
	ParentId *uuid.UUID `json:"parentId"` // 父级Id
}

// 解析获取支出列表Dto
func ParseListDto(values url.Values) ListDto {
	dto := ListDto{}
	parentId, err := uuid.Parse(values.Get("parentId"))
	if err == nil {
		dto.ParentId = &parentId
	}
	return dto
}

// 获取支出分页列表Dto
type PagedListDto struct {
	pagedlist.GetDto
	ListDto
}

// 解析获取支出分页列表Dto
func ParsePagedListDto(values url.Values) PagedListDto {
	return PagedListDto{
		GetDto:  pagedlist.ParseGetDto(values),
		ListDto: ParseListDto(values),
	}
}
