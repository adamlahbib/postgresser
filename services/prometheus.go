package services

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Prometheus struct {
	counters         map[string]*prometheus.CounterVec
	histograms       map[string]*prometheus.HistogramVec
	gauges           map[string]*prometheus.GaugeVec
	numCounterLabels map[string]int
}

type MetricType int

const (
	Counter MetricType = iota
	Histogram
	Gauge
)

type Metric struct {
	Type        MetricType
	Name        string
	Description string
	Buckets     []float64
	Labels      []string
}

const (
	duplicatedMetricLabelError   = "duplicate metrics collector registration attempted"
	counterMetricNotFoundError   = "counter metric not found: %s"
	histogramMetricNotFoundError = "histogram metric not found: %s"
	gaugeMetricNotFoundError     = "gauge metric not found: %s"
)

func NewPrometheusService(metrics []Metric) (*Prometheus, error) {
	counters := make(map[string]*prometheus.CounterVec)
	histograms := make(map[string]*prometheus.HistogramVec)
	gauges := make(map[string]*prometheus.GaugeVec)
	numCounterLabels := make(map[string]int)
	for _, m := range metrics {
		var err error
		switch m.Type {
		case Counter:
			numCounterLabels[m.Name] = len(m.Labels)
			counters[m.Name] = prometheus.NewCounterVec(prometheus.CounterOpts{
				Name: m.Name,
				Help: m.Description,
			}, m.Labels)
			err = prometheus.Register(counters[m.Name])
		case Histogram:
			histograms[m.Name] = prometheus.NewHistogramVec(prometheus.HistogramOpts{
				Name:    m.Name,
				Help:    m.Description,
				Buckets: m.Buckets,
			}, m.Labels)
			err = prometheus.Register(histograms[m.Name])
		case Gauge:
			gauges[m.Name] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: m.Name,
				Help: m.Description,
			}, m.Labels)
			err = prometheus.Register(gauges[m.Name])
		}
		if err != nil {
			return nil, err
		}
	}
	return &Prometheus{
		counters:         counters,
		histograms:       histograms,
		gauges:           gauges,
		numCounterLabels: numCounterLabels,
	}, nil
}

func (p *Prometheus) IncCounterMetric(metric string, value float64, labels map[string]string) error {
	if _, ok := p.counters[metric]; !ok {
		return fmt.Errorf(counterMetricNotFoundError, metric)
	}
	p.counters[metric].With(labels).Add(value)
	return nil
}

func (p *Prometheus) ObserveHistogramMetric(metric string, value float64, labels map[string]string) error {
	if _, ok := p.histograms[metric]; !ok {
		return fmt.Errorf(histogramMetricNotFoundError, metric)
	}
	p.histograms[metric].With(labels).Observe(value)
	return nil
}

func (p *Prometheus) SetGaugeMetric(metric string, value float64, labels map[string]string) error {
	if _, ok := p.gauges[metric]; !ok {
		return fmt.Errorf(gaugeMetricNotFoundError, metric)
	}
	p.gauges[metric].With(labels).Set(value)
	return nil
}

// Describe sends the descriptors of each metric over to the provided channel.
func (p *Prometheus) Describe(descs chan<- *prometheus.Desc) {
	for _, c := range p.counters {
		c.Describe(descs)
	}
	for _, h := range p.histograms {
		h.Describe(descs)
	}
	for _, g := range p.gauges {
		g.Describe(descs)
	}
}

// collect is called by the Prometheus handler serving the /metrics endpoint.
func (p *Prometheus) Collect(metrics chan<- prometheus.Metric) {
	for _, c := range p.counters {
		c.Collect(metrics)
	}
	for _, h := range p.histograms {
		h.Collect(metrics)
	}
	for _, g := range p.gauges {
		g.Collect(metrics)
	}
}

const (
	MetricDeploymentAccessTotal       = "deployment_access_total"
	MetricDeploymentAccessFailedTotal = "deployment_access_failed_total"
	LabelID                           = "id"
	LabelOperation                    = "operation"
)

func GetMetricsDefinition() []Metric {
	return []Metric{
		{
			Type:        Counter,
			Name:        MetricDeploymentAccessTotal,
			Description: "Total number of deployments accessed",
			Labels:      []string{LabelID, LabelOperation},
		},
		{
			Type:        Counter,
			Name:        MetricDeploymentAccessFailedTotal,
			Description: "Total number of failed deployments accessed",
			Labels:      []string{LabelID, LabelOperation},
		},
	}
}
