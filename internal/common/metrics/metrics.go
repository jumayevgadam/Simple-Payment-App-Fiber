package metrics

import (
	"net/http"
	"time"

	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	TotalRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of http(fiber) requests",
	})

	// ResponseTimes is.
	ResponseTimes = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_response_times_seconds",
		Help:    "Response times for http requests",
		Buckets: prometheus.LinearBuckets(0.01, 0.05, 20),
	})

	// ResponseStatus is.
	ResponseStatus = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_response_status",
		Help: "Status of http responses",
	}, []string{"status"})

	// ErrorRates is.
	ErrorRates = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_error_rates_total",
		Help: "Total number of error responses in http requests",
	})
)

// Listen method for connecting and using metrics for seeing app details.
func Listen(metricsPort string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:         metricsPort,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	errListen := server.ListenAndServe()
	if errListen != nil {
		return errlst.NewInternalServerError("error setting and serving metrics port")
	}

	return nil
}
