package api

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsController struct {
	*BaseController
}

/*
var messageTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "message_metric_messages_total",
		Help: "The total number of messages",
	},
)

var dingTalkMessageTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "message_metric_dingtalk_messages_total",
		Help: "The total number of dingtalk messages",
	},
)
*/

var totalMessages = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "message_metric_messages_total",
		Help: "The total number of messages",
	},
)

var totalFailedMessages = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "message_metric_failed_messages_total",
		Help: "The total number of failed messages",
	},
)

var totalSentMessages = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "message_metric_sent_messages_total",
		Help: "The total number of messages that has been sent",
	},
)

var dingTalkMessages = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "message_metric_dingtalk_messages_total",
		Help: "The total number of dingtalk messages",
	},
)

var dingTalkFailedMessages = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "message_metric_failed_dingtalk_messages_total",
		Help: "The total number of failed dingtalk messages",
	},
)

var dingTalkSentMessages = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "message_metric_sent_dingtalk_messages_total",
		Help: "The total number of dingtalk messages that has been sent",
	},
)

var emailMessages = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "message_metric_email_messages_total",
		Help: "The total number of email messages",
	},
)

var emailFailedMessages = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "message_metric_failed_email_messages_total",
		Help: "The total number of failed email messages",
	},
)

var emailSentMessages = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "message_metric_sent_email_messages_total",
		Help: "The total number of email messages that has been sent",
	},
)

var httpRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "The HTTP request latencies in seconds",
		Buckets: nil,
	}, []string{"method", "endpoint", "code", "env"},
)

var httpRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "The total number of http request to a route",
	}, []string{"method", "endpoint", "code", "env"},
)

func init() {
	// 注册监控指标
	prometheus.MustRegister(totalMessages)
	prometheus.MustRegister(totalFailedMessages)
	prometheus.MustRegister(totalSentMessages)

	prometheus.MustRegister(dingTalkMessages)
	prometheus.MustRegister(dingTalkFailedMessages)
	prometheus.MustRegister(dingTalkSentMessages)

	prometheus.MustRegister(emailMessages)
	prometheus.MustRegister(emailFailedMessages)
	prometheus.MustRegister(emailSentMessages)

	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(httpRequestTotal)
}

func newMetricsController(baseController *BaseController) *MetricsController {
	return &MetricsController{baseController}
}

func (mc *MetricsController) ServeHTTP(httpwriter http.ResponseWriter, httpRequest *http.Request) {
	// 添加各种指标
	mr := mc.bs.MetricsService.Count()

	totalMessages.Set(float64(mr.TotalMessages))
	totalFailedMessages.Set(float64(mr.TotalFailedMessages))
	totalSentMessages.Set(float64(mr.TotalSentMessages))

	dingTalkMessages.Set(float64(mr.DingTalkMessages))
	dingTalkFailedMessages.Set(float64(mr.DingTalkFailedMessages))
	dingTalkSentMessages.Set(float64(mr.DingTalkSentMessages))

	emailMessages.Set(float64(mr.EmailMessages))
	emailFailedMessages.Set(float64(mr.EmailFailedMessages))
	emailSentMessages.Set(float64(mr.EmailSentMessages))

	promhttp.Handler().ServeHTTP(httpwriter, httpRequest)
}

func metrics(c *restful.Container, controller *BaseController) {
	mc := newMetricsController(controller)

	c.Handle("/metrics", mc)
}
