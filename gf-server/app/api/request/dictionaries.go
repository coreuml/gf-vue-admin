package request

// TODO 添加p与vTag标签
type CreateDictionary struct {
	Name   string `orm:"name" json:"name"`     // 字典名（中）
	Type   string `orm:"type" json:"type"`     // 字典名（英）
	Status int    `orm:"status" json:"status"` // 状态
	Desc   string `orm:"desc" json:"desc"`     // 描述
}

// TODO 添加p与vTag标签
type DeleteDictionary struct {
	ID float64 `json:"id"`
}

// TODO 添加p与vTag标签
type UpdateDictionary struct {
	ID float64 `json:"id"`
}

// TODO 添加p与vTag标签
type GetDictionary struct {
	ID   float64 `orm:"id" json:"id"`
	Type string  `orm:"type" json:"type"` // 字典名（英）
}

type DictionarySearch struct {
	Id     uint   `orm:"id,primary" json:"id"`     // 自增ID
	Name   string `orm:"name"       json:"name"`   // 字典名（中）
	Type   string `orm:"type"       json:"type"`   // 字典名（英）
	Status int    `orm:"status"     json:"status"` // 状态
	Desc   string `orm:"desc"       json:"desc"`   // 描述
	PageInfo
}
