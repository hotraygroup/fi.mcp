package types

type Suggest struct {
	Code int `json:"code"`
	Data []struct {
		Code      string `json:"code"`
		Label     string `json:"label"`
		Query     string `json:"query"`
		State     int    `json:"state"`
		StockType int    `json:"stock_type"`
		Type      int    `json:"type"`
	} `json:"data"`
	Message string `json:"message"`
	Meta    struct {
		Count       int   `json:"count"`
		Feedback    int   `json:"feedback"`
		HasNextPage bool  `json:"has_next_page"`
		MaxPage     int   `json:"maxPage"`
		Page        int   `json:"page"`
		QueryID     int64 `json:"query_id"`
		Size        int   `json:"size"`
	} `json:"meta"`
	Success bool `json:"success"`
}
