package internal

import (
    "github.com/prometheus/client_golang/prometheus"
    "math/rand"
)

type metrics struct {
    cpuTemp prometheus.Gauge
    hddTemp prometheus.Gauge
}

func HandlerRandomMetrics() *prometheus.Registry {
    reg := prometheus.NewRegistry()
    m := newMetric(reg)
    m.cpuTemp.Set(rand.Float64())
    m.hddTemp.Set(rand.Float64())
    return reg
}

func newMetric(reg prometheus.Registerer) *metrics {
    m := &metrics {
        cpuTemp: prometheus.NewGauge(prometheus.GaugeOpts{
            Name: "cpu_temperature_celsius",
            Help: "Current temperature of the CPU",
        }),
        hddTemp: prometheus.NewGauge(prometheus.GaugeOpts{
            Name: "hdd_temperature_celsius",
            Help: "Current temperature of the HDD",
        }),
    }
    reg.MustRegister(m.cpuTemp)
    reg.MustRegister(m.hddTemp)

    return m
}
