package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "api_request_duration_seconds",
		Help:    "Request duration in seconds.",
		Buckets: prometheus.DefBuckets,
	},
		[]string{"method", "path"})

	RequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_request_count_total",
		Help: "Total number of API requests.",
	},
		[]string{"method", "path"})

	CacheDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cache_duration_seconds",
		Help:    "Cache duration in seconds.",
		Buckets: prometheus.DefBuckets,
	},
		[]string{"endpoint"})

	DBDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "db_duration_seconds",
		Help:    "DB duration in seconds.",
		Buckets: prometheus.DefBuckets,
	},
		[]string{"function"})

	ExternalAPIDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "external_api_duration_seconds",
		Help:    "External API duration in seconds.",
		Buckets: prometheus.DefBuckets,
	},
		[]string{"method", "path"})
)

func MustRegister() {
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(CacheDuration)
	prometheus.MustRegister(DBDuration)
	prometheus.MustRegister(ExternalAPIDuration)
}
