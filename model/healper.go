/**
 * 功能描述: 数据库操作帮助
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package model

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

const (
	KeySort   = "sort"
	KeyFilter = "filter"
)

type PageResult struct {
	Count   int64       `json:"count"`
	List    interface{} `json:"list"`
	Size    int64       `json:"size"`
	PageNum int64       `json:"page_num"`
}

type WhereMatch string

type CompareOperator string

type Filter struct {
	Field string
	Op    CompareOperator
	Value string
}

const (
	OpEqual        CompareOperator = "eq"
	OpNotEqual     CompareOperator = "not_eq"
	OpLessThan     CompareOperator = "lt"
	OpLessEqual    CompareOperator = "lte"
	OpGreaterThan  CompareOperator = "gt"
	OpGreaterEqual CompareOperator = "gte"
	OpContain      CompareOperator = "contains"
	OpNotContain   CompareOperator = "not_contains"
	OpStartWith    CompareOperator = "starts_with"
	OpEndWith      CompareOperator = "ends_with"
)

var (
	// the variable is used as read only (so no lock)
	operatorMap = map[CompareOperator]bool{
		OpEqual:        true,
		OpNotEqual:     true,
		OpLessThan:     true,
		OpLessEqual:    true,
		OpGreaterThan:  true,
		OpGreaterEqual: true,
		OpContain:      true,
		OpNotContain:   true,
		OpStartWith:    true,
		OpEndWith:      true,
	}
)

// 分页
type PaginationRequest struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

type Pagination struct {
	From int64 `json:"from"`
	Size int64 `json:"size"`
}

type DateSelector struct {
	Pagination *Pagination
	Sorter     *SortBy
	Filters    []Filter
}

type FilterQuery struct {
	FilterByList *[]Filter
}

func DateSelectFromContext(c *gin.Context) DateSelector {
	var result DateSelector
	result.Pagination = ParsePagination(c)
	result.Filters = ParseFilter(c)
	result.Sorter = ParseSorting(c, "")
	return result
}

// 解析请求参数中的翻页信息
// eg.  ?page=1&size=1
// 如未指定相关信息,则默认按照1页10条记录返回
func ParsePagination(c *gin.Context) *Pagination {
	paginationRequest := PaginationRequest{
		Page: 1,
		Size: 10,
	}

	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	if len(pageStr) > 0 && len(sizeStr) > 0 {
		page, err1 := strconv.ParseInt(pageStr, 10, 64)
		size, err2 := strconv.ParseInt(sizeStr, 10, 64)
		if err1 != nil || err2 != nil {
			page := &Pagination{
				From: 0,
				Size: 10,
			}
			return page
		}
		paginationRequest.Page = page
		// 每页数据允许 1-100条
		if size < 1 || size > 100 {
			paginationRequest.Size = 10
		} else {
			paginationRequest.Size = size
		}

	}
	page := &Pagination{
		From: (paginationRequest.Page - 1) * paginationRequest.Size,
		Size: paginationRequest.Size,
	}
	return page
}

// 解析请求参数中的过滤条件
// eg. filter=username admin,age eq 18
func ParseFilter(c *gin.Context) []Filter {
	queryStr := strings.Trim(c.Query(KeyFilter), " ")
	var filters []Filter
	strs := strings.Split(queryStr, ",")
	for _, str := range strs {
		// split str to 3 sub strings
		items := strings.SplitN(strings.Trim(str, " "), " ", 3)
		if len(items) < 2 { // length must be 2+
			continue
		}

		// default operation is OpContain
		if len(items) == 2 {
			items = append(items, items[1])
			items[1] = string(OpContain)
		}

		// check the second param
		op := CompareOperator(items[1])
		//items[2] = sqlutil.EscapeUnderlineInLikeStatement(items[2])
		if _, ok := operatorMap[op]; ok {
			filters = append(filters, Filter{items[0], op, items[2]})
		}
	}
	return filters
}

type SortOrder string

const (
	KeySortBy    = "sort_by"
	KeySortOrder = "sort_order"
)

const (
	SortAsc  SortOrder = "asc"
	SortDesc SortOrder = "desc"
)

// 排序
type SortBy struct {
	By    string    `json:"sort_by"`
	Order SortOrder `json:"sort_order"`
}

// 获取请求参数中的排序信息
// eg. ?sort_by=username&sort_order=desc   默认desc
func ParseSorting(c *gin.Context, defaultSortKey string) *SortBy {
	res := &SortBy{}
	queryStr := c.Query(KeySortBy)
	if len(queryStr) == 0 {
		if defaultSortKey == "" {
			return nil
		}
		res.By = defaultSortKey
	} else {
		res.By = queryStr
	}

	// sort order must be "asc" or "desc", default use desc
	queryStr = c.Query(KeySortOrder)
	if len(queryStr) == 0 {
		queryStr = string(SortDesc)
	}

	res.Order = SortOrder(queryStr)

	if res.Order != SortDesc && res.Order != SortAsc {
		return nil
	}
	return res
}
