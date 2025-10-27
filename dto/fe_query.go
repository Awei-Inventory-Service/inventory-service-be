package dto

type Filter struct {
	Key      string   `json:"key"`
	Values   []string `json:"values"`
	Wildcard string   `json:"wildcard"`
}

type Order struct {
	Key   string `json:"key"`
	IsAsc bool   `json:"is_asc"`
}

type GetListRequest struct {
	Filter []Filter `json:"filter"`
	Order  []Order  `json:"order"`
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
}
