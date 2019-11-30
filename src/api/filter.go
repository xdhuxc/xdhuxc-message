package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/xdhuxc/xdhuxc-message/src/conf"
	"strconv"
	"time"

	"github.com/emicklei/go-restful"

	"github.com/xdhuxc/xdhuxc-message/src/model"
)

func (b *BaseController) Page(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	if req.Request.Method != "GET" {
		chain.ProcessFilter(req, resp)
		return
	}

	var pageSize int64 = 10
	var cpage int64 = 1
	if ps, err := strconv.ParseInt(req.QueryParameter("limit"), 10, 64); err == nil &&
		ps > 0 {
		pageSize = ps
	}
	if p, err := strconv.ParseInt(req.QueryParameter("page"), 10, 64); err == nil &&
		p > 0 {
		cpage = p
	}

	offset := (cpage - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	page := model.Page{
		PageSize: pageSize,
		Offset:   offset,
		Page:     cpage,
		Query:    req.QueryParameter("query"),
	}

	switch req.QueryParameter("sort") {
	case "asc":
		page.Sort = "asc"
	default:
		page.Sort = "desc"
	}

	switch req.QueryParameter("order_by") {
	case "name":
		page.OrderBy = "name"
	default:
		page.OrderBy = "update_time"
	}

	req.SetAttribute("page", page)

	chain.ProcessFilter(req, resp)
}

func (b *BaseController) metrics(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	start := time.Now()
	duration := float64(time.Since(start)) / float64(time.Second)

	httpRequestTotal.With(prometheus.Labels{
		"method":   req.Request.Method,
		"endpoint": req.Request.URL.Path,
		"code":     strconv.Itoa(resp.StatusCode()),
		"env":      conf.GetConfiguration().Env,
	}).Inc()

	httpRequestDuration.With(prometheus.Labels{
		"method":   req.Request.Method,
		"endpoint": req.Request.URL.Path,
		"code":     strconv.Itoa(resp.StatusCode()),
		"env":      conf.GetConfiguration().Env,
	}).Observe(duration)

	chain.ProcessFilter(req, resp)
}
