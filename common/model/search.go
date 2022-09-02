package model


// SearchCounselorModel 諮商師搜尋Model
type SearchCounselorModel struct {
	Filter  string  // 搜尋方式 json型態
	Limit   int    // "資料量限制，預設20"
	Offset  int    // "資料量限制，預設0"
	SortBy  string //排序參照，預設ID
	OrderBy string //排序方式，預設desc
}

