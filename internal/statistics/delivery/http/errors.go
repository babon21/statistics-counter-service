package http

import "errors"

var (
	SortByParamIsEmpty  = errors.New("sort_by parameter is empty")
	OrderByParamIsEmpty = errors.New("order_by parameter is empty")
	WrongOrderByParam   = errors.New("order_by parameter is wrong")
	WrongSortField      = errors.New("Sort field is wrong in sort param")
)
