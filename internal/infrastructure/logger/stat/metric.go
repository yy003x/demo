package stat

import (
	kratosPrometheus "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	UrlDurationMsHistogram    metrics.Observer
	ServerRequestErrorCounter metrics.Counter
	SqlDurationMsHistogram    metrics.Observer
}

func NewMetrics() *Metrics {
	//服务请求耗时/总请求数
	_metricUrlDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "lib_url_duration_ms",
		Help:    "server requests duration(ms).",
		Buckets: []float64{5, 10, 15, 20, 25, 50, 100, 200, 500, 1000, 2000},
	}, []string{"action"})

	//服务异常数
	_metricServerErrorRequest := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "lib_server_request_error_total",
		Help: "The total number of http or grpc request error",
	}, []string{"kind", "action"})

	//sql请求耗时
	_metricSqlDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "lib_sql_duration_ms",
		Help:    "sql requests duration(ms).",
		Buckets: []float64{5, 10, 15, 20, 25, 50, 100, 200, 500},
	}, []string{})

	prometheus.MustRegister(_metricServerErrorRequest, _metricUrlDuration, _metricSqlDuration)

	return &Metrics{
		UrlDurationMsHistogram:    kratosPrometheus.NewHistogram(_metricUrlDuration),
		ServerRequestErrorCounter: kratosPrometheus.NewCounter(_metricServerErrorRequest),
		SqlDurationMsHistogram:    kratosPrometheus.NewHistogram(_metricSqlDuration),
	}
}
