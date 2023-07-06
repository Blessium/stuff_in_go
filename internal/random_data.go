package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"math/rand"
	"time"
)

type metrics struct {
	cpuTemp prometheus.Gauge
	hddTemp prometheus.Gauge
}

func HandlerRandomMetrics() *prometheus.Registry {
	reg := prometheus.NewRegistry()
	m := newMetric(reg)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	bruh := func() {
		for {
			m.cpuTemp.Set(r.Float64())
			m.hddTemp.Set(r.Float64())
		}
	}

    go bruh()

	return reg
}

func goRandomicNumberGen() {

}

func newMetric(reg prometheus.Registerer) *metrics {
	m := &metrics{
		cpuTemp: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Name: "cpu_temperature_celsius",
			Help: "Current temperature of the CPU",
		}),
		hddTemp: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Name: "hdd_temperature_celsius",
			Help: "Current temperature of the HDD",
		}),
	}
	return m
}
