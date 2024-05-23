package outlay

import (
	"net/url"
	"strconv"
	"time"

	outlay_base "financialrecord-backend/bases/outlay"
	"financialrecord-backend/dtos/pagedlist"
	"financialrecord-backend/models/outlay"
	"financialrecord-backend/utils"

	"github.com/google/uuid"
)

// 支出Dto
type Dto = outlay.Outlay

// 创建支出Dto
type CreateDto = outlay_base.Props

// 更新支出Dto
type UpdateDto = outlay_base.Base

// 获取支出列表Dto
type ListDto struct {
	CatId    *uuid.UUID `json:"catId"`    // 类型Id
	LowMoney *float32   `json:"lowMoney"` // 最低金额
	TopMoney *float32   `json:"topMoney"` // 最高金额
	STime    *time.Time `json:"sTime"`    // 开始时间
	ETime    *time.Time `json:"eTime"`    // 结束时间
	UserId   *uuid.UUID `json:"userId"`   // 用户Id
}

// 解析获取支出列表Dto
func ParseListDto(values url.Values) ListDto {
	dto := ListDto{}
	catId, err := uuid.Parse(values.Get("catId"))
	if err == nil {
		dto.CatId = &catId
	}
	lowMoney, err := strconv.ParseFloat(values.Get("lowMoney"), 32)
	if err == nil {
		temp := float32(lowMoney)
		dto.LowMoney = &temp
	}
	topMoney, err := strconv.ParseFloat(values.Get("topMoney"), 32)
	if err == nil {
		temp := float32(topMoney)
		dto.TopMoney = &temp
	}
	sTime, err := utils.ParseTime(values.Get("sTime"))
	if err == nil {
		dto.STime = sTime
	}
	eTime, err := utils.ParseTime(values.Get("eTime"))
	if err == nil {
		dto.ETime = eTime
	}
	userId, err := uuid.Parse(values.Get("userId"))
	if err == nil {
		dto.UserId = &userId
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
