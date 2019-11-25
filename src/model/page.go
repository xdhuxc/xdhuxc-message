package model

type Page struct {
	PageSize   int64  `json:"limit"`
	Offset     int64  `json:"offset"`
	Page       int64  `json:"page"`
	TotalCount int64  `json:"-"`
	Query      string `json:"-"`
	OrderBy    string `json:"orderBy"`
	Sort       string `json:"sort"`
}

type Result struct {
	TotalCount  *int64      `json:"totalCount,omitempty"`
	TotalPage   *int64      `json:"totalPage,omitempty"`
	CurrentPage *int64      `json:"currentPage,omitempty"`
	PageSize    *int64      `json:"pageSize,omitempty"`
	Results     interface{} `json:"result"`
	Code        int64       `json:"code"`
}

func NewResult(count int64, page *Page, results interface{}) Result {
	var result Result
	var totalPage int64
	if page != nil {
		result = Result{
			TotalCount:  &count,
			CurrentPage: &page.Page,
			PageSize:    &page.PageSize,
			TotalPage:   &totalPage,
			Results:     results,
			Code:        0,
		}

		pc := count / page.PageSize
		result.TotalPage = &pc
		if count%page.PageSize > 0 {
			*(result.TotalPage) += 1
		}
	} else {
		result = Result{
			Results: results,
			Code:    0,
		}
	}

	return result
}
