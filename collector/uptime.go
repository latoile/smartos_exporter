// Uptime collector
// this will :
//  - call uptime
//  - gather load average
//  - feed the collector

package collector

import (
	"strconv"
	// Psutil Go
	"github.com/shirou/gopsutil/host"
	// Prometheus Go toolset
	"github.com/prometheus/client_golang/prometheus"
)

// UpTimeCollector declares the data type within the prometheus metrics
// package.
type UpTimeCollector struct {
	UpTime  prometheus.Gauge
}

// NewUpTimeExporter returns a newly allocated exporter UpTimeCollector.
// It exposes the UpTime Server in Second.
func NewUpTimeExporter() (*UpTimeCollector, error) {
        return &UpTimeCollector{
                UpTime: prometheus.NewGauge(prometheus.GaugeOpts{
                        Name: "smartos_up_time",
                        Help: "Up Time of server in second",
                }),
        }, nil
}

// Describe describes all the metrics.
func (e *UpTimeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.UpTime.Desc()
}

// Collect fetches the stats.
func (e *UpTimeCollector) Collect(ch chan<- prometheus.Metric) {
	e.uptime()
	ch <- e.UpTime
}

func (e *UpTimeCollector) uptime() error {
  // Uptime Server in second
	uptimesecond,_ := host.Uptime()
	uptimeconvert := strconv.FormatUint(uptimesecond, 10)
	uptime, err := strconv.ParseFloat(uptimeconvert, 64)
	if err != nil {
		return err
	}
	e.UpTime.Set(uptime)

	return nil
}
