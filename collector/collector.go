package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/williamclot/emby_exporter/emby"
)

var _ StatusSource = &emby.Client{}

type StatusSource interface {
	Status() (*emby.Status, error)
}

type Collector struct {
	Info *prometheus.Desc

	DeviceCount        *prometheus.Desc
	MovieCount         *prometheus.Desc
	SeriesCount        *prometheus.Desc
	EpisodeCount       *prometheus.Desc
	UserCount          *prometheus.Desc
	FailedLoginCount   *prometheus.Desc
	MediaStreamCount   *prometheus.Desc
	TranscodingCount   *prometheus.Desc
	TranscodingBitRate *prometheus.Desc
	VideoBitRate       *prometheus.Desc
	AudioBitRate       *prometheus.Desc
	AudioChannels      *prometheus.Desc

	ss StatusSource
}

func New(ss StatusSource) prometheus.Collector {
	labels := []string{"server"}
	namespace := "emby"

	return &Collector{
		Info: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "info"),
			"Metadata about a given Emby server.",
			[]string{"server", "hostname", "version"},
			nil,
		),

		DeviceCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "device", "count"),
			"Number of devices configured in Emby.",
			labels,
			nil,
		),

		MovieCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "movie", "count"),
			"Number of movies available in Emby.",
			labels,
			nil,
		),

		SeriesCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "series", "count"),
			"Number of tv shows available in Emby.",
			labels,
			nil,
		),

		EpisodeCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "episode", "count"),
			"Number of tv show episodes available in Emby.",
			labels,
			nil,
		),

		UserCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "user", "count"),
			"Number of users configured in Emby.",
			labels,
			nil,
		),

		FailedLoginCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "failed", "login"),
			"Failed login counts per user in Emby.",
			[]string{"server", "user"},
			nil,
		),

		MediaStreamCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "stream", "count"),
			"Number of media streams being handled by Emby.",
			labels,
			nil,
		),

		TranscodingCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "transcoding", "count"),
			"Number of media streams being transcoded by Emby.",
			labels,
			nil,
		),

		TranscodingBitRate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "transcoding", "bitrate"),
			"Bitrate of transcoded mediastream.",
			[]string{"server", "user"},
			nil,
		),

		VideoBitRate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "video", "bitrate"),
			"Bitrate of original video file.",
			[]string{"server", "user"},
			nil,
		),

		AudioBitRate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "audio", "bitrate"),
			"Bitrate of original audio file.",
			[]string{"server", "user"},
			nil,
		),

		AudioChannels: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "audio", "channels"),
			"Number of channels in original audio file.",
			[]string{"server", "user"},
			nil,
		),

		ss: ss,
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ds := []*prometheus.Desc{
		c.Info,
		c.DeviceCount,
		c.MovieCount,
		c.SeriesCount,
		c.EpisodeCount,
		c.UserCount,
		c.FailedLoginCount,
		c.MediaStreamCount,
		c.TranscodingCount,
		c.TranscodingBitRate,
		c.VideoBitRate,
		c.AudioBitRate,
		c.AudioChannels,
	}
	for _, d := range ds {
		ch <- d
	}
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	s, err := c.ss.Status()
	if err != nil {
		log.Printf("failed collecting emby metrics: %v", err)
		ch <- prometheus.NewInvalidMetric(c.Info, err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.Info,
		prometheus.GaugeValue,
		1,
		s.ServerName, s.Hostname, s.Version,
	)

	ch <- prometheus.MustNewConstMetric(
		c.DeviceCount,
		prometheus.GaugeValue,
		float64(s.DeviceCount),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.MovieCount,
		prometheus.GaugeValue,
		float64(s.MovieCount),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.SeriesCount,
		prometheus.GaugeValue,
		float64(s.SeriesCount),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.EpisodeCount,
		prometheus.GaugeValue,
		float64(s.EpisodeCount),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.UserCount,
		prometheus.GaugeValue,
		float64(s.UserCount),
		s.ServerName,
	)

	for k, v := range s.FailedLoginCount {
		ch <- prometheus.MustNewConstMetric(
			c.FailedLoginCount,
			prometheus.GaugeValue,
			float64(v),
			s.ServerName, k,
		)
	}

	ch <- prometheus.MustNewConstMetric(
		c.MediaStreamCount,
		prometheus.GaugeValue,
		float64(s.MediaStreamCount),
		s.ServerName,
	)

	ch <- prometheus.MustNewConstMetric(
		c.TranscodingCount,
		prometheus.GaugeValue,
		float64(s.TranscodingCount),
		s.ServerName,
	)

	for k, v := range s.TranscodingBitRate {
		ch <- prometheus.MustNewConstMetric(
			c.TranscodingBitRate,
			prometheus.GaugeValue,
			float64(v),
			s.ServerName, k,
		)
	}

	for k, v := range s.VideoBitRate {
		ch <- prometheus.MustNewConstMetric(
			c.VideoBitRate,
			prometheus.GaugeValue,
			float64(v),
			s.ServerName, k,
		)
	}

	for k, v := range s.AudioBitRate {
		ch <- prometheus.MustNewConstMetric(
			c.AudioBitRate,
			prometheus.GaugeValue,
			float64(v),
			s.ServerName, k,
		)
	}

	for k, v := range s.AudioChannels {
		ch <- prometheus.MustNewConstMetric(
			c.AudioChannels,
			prometheus.GaugeValue,
			float64(v),
			s.ServerName, k,
		)
	}
}
