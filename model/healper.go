/**
 * 功能描述: 数据库操作帮助
 * @Date: 2019-04-16
 * @author: lixiaoming
 */
package model

// 定义SQL条件
type SQLCondition struct {
	Where
	Pagination
	SortBy
}

type WhereMatch string

const (
	WhereLike             WhereMatch = "like"
	WhereIs               WhereMatch = "="
	WhereIsNot            WhereMatch = "!="
	WhereGreatThan        WhereMatch = ">"
	WhereLessThan         WhereMatch = "<"
	WhereGreatThanOrEqual WhereMatch = ">="
	WhereLessThanOrEqual  WhereMatch = "<="
)

type WhereConditionsGroup struct {
	WhereConditions []Where
	Relation        string
}

// where 条件
type Where struct {
	Field   string     `json:"field"`
	Match   WhereMatch `json:"match"`
	Keyword string     `json:"keyword"`
}

// 分页
type Pagination struct {
	From int `json:"from"`
	Size int `json:"size"`
}

type SortOrder string

const (
	SortAsc  SortOrder = "asc"
	SortDesc SortOrder = "desc"
)

// 排序
type SortBy struct {
	By    string    `json:"sort_by"`
	Order SortOrder `json:"sort_order"`
}
