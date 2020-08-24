package entity

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total    int  `json:"total"`
	Current  uint `json:"current"`
	PageSize uint `json:"page_size"`
}

// PaginationParam 分页查询条件
type PaginationParam struct {
	Pagination bool `form:"-"`                                      // 是否使用分页查询
	OnlyCount  bool `form:"-"`                                      // 是否仅查询count
	Current    uint `form:"current,default=1"`                      // 当前页
	PageSize   uint `form:"page_size,default=10" binding:"max=100"` // 页大小
}

// GetCurrent 获取当前页
func (a PaginationParam) GetCurrent() uint {
	return a.Current
}

// GetPageSize 获取页大小
func (a PaginationParam) GetPageSize() uint {
	pageSize := a.PageSize
	if a.PageSize == 0 {
		pageSize = 100
	}
	return pageSize
}

// OrderDirection 排序方向
type OrderDirection int

const (
	// OrderByASC 升序排序
	OrderByASC OrderDirection = 1
	// OrderByDESC 降序排序
	OrderByDESC OrderDirection = 2
)

// OrderField 排序字段
type OrderField struct {
	Key       string         // 字段名(字段名约束为小写蛇形)
	Direction OrderDirection // 排序方向
}

// NewOrderFieldWithKeys 创建排序字段(默认升序排序)，可指定不同key的排序规则
// keys 需要排序的key
// directions 排序规则，按照key的索引指定，索引默认从0开始
func NewOrderFieldWithKeys(keys []string, directions ...map[int]OrderDirection) []*OrderField {
	m := make(map[int]OrderDirection)
	if len(directions) > 0 {
		m = directions[0]
	}

	fields := make([]*OrderField, len(keys))
	for i, key := range keys {
		d := OrderByASC
		if v, ok := m[i]; ok {
			d = v
		}

		fields[i] = NewOrderField(key, d)
	}

	return fields
}

// NewOrderFields 创建排序字段列表
func NewOrderFields(orderFields ...*OrderField) []*OrderField {
	return orderFields
}

// NewOrderField 创建排序字段
func NewOrderField(key string, d OrderDirection) *OrderField {
	return &OrderField{
		Key:       key,
		Direction: d,
	}
}

// NewIDResult 创建响应唯一标识实例
func NewIDResult(id int64) *IDResult {
	return &IDResult{
		ID: id,
	}
}

// IDResult 响应唯一标识
type IDResult struct {
	ID int64 `json:"id"`
}
