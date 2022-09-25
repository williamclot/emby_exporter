package collector

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/williamclot/emby_exporter/emby"
)

func TestEmbyCollector(t *testing.T) {
	var tests = []struct {
		desc    string
		ss      *testStatusSource
		matches []*regexp.Regexp
	}{
		{
			desc: "empty",
			ss: &testStatusSource{
				s: &emby.Status{},
			},
		},
		{
			desc: "full",
			ss: &testStatusSource{
				s: &emby.Status{
					ServerName:   "mars",
					Hostname:     "mars.emby.com",
					Version:      "1.1.2",
					DeviceCount:  9,
					MovieCount:   123,
					SeriesCount:  12,
					EpisodeCount: 362,
					UserCount:    2,
					FailedLoginCount: map[string]int{
						"bill": 0,
						"joe":  2,
					},
					MediaStreamCount: 2,
					TranscodingCount: 1,
					TranscodingBitRate: map[string]int{
						"bill": 9614027,
					},
					VideoBitRate: map[string]int{
						"bill": 9614027,
						"joe":  950000,
					},
					AudioBitRate: map[string]int{
						"bill": 10200,
						"joe":  10300,
					},
					AudioChannels: map[string]int{
						"bill": 2,
						"joe":  5,
					},
				},
			},
			matches: []*regexp.Regexp{
				regexp.MustCompile(`emby_device_count{server="mars"} 9`),
				regexp.MustCompile(`emby_movie_count{server="mars"} 123`),
				regexp.MustCompile(`emby_series_count{server="mars"} 12`),
				regexp.MustCompile(`emby_episode_count{server="mars"} 362`),
				regexp.MustCompile(`emby_user_count{server="mars"} 2`),
				regexp.MustCompile(`emby_stream_count{server="mars"} 2`),
				regexp.MustCompile(`emby_transcoding_count{server="mars"} 1`),
				regexp.MustCompile(`emby_transcoding_bitrate{server="mars",user="bill"} 9.614027e\+06`),
				regexp.MustCompile(`emby_video_bitrate{server="mars",user="bill"} 9.614027e\+06`),
				regexp.MustCompile(`emby_video_bitrate{server="mars",user="joe"} 950000`),
				regexp.MustCompile(`emby_audio_bitrate{server="mars",user="bill"} 10200`),
				regexp.MustCompile(`emby_audio_bitrate{server="mars",user="joe"} 10300`),
				regexp.MustCompile(`emby_audio_channels{server="mars",user="bill"} 2`),
				regexp.MustCompile(`emby_audio_channels {server="mars",user="joe"} 5`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			out := testCollector(t, New(tt.ss))
			fmt.Println(string(out))
			for _, m := range tt.matches {
				if !m.Match(out) {
					t.Fatalf("output failed to match regex (regexp: %v)", m)
				}
			}
		})
	}
}

func testCollector(t *testing.T, collector prometheus.Collector) []byte {
	t.Helper()

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(collector)

	srv := httptest.NewServer(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	defer srv.Close()

	c := &http.Client{Timeout: 1 * time.Second}
	resp, err := c.Get(srv.URL)
	if err != nil {
		t.Fatalf("failed to HTTP GET data from prometheus: %v", err)
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read server response: %v", err)
	}

	return buf
}

var _ StatusSource = &testStatusSource{}

type testStatusSource struct {
	s *emby.Status
}

func (ss *testStatusSource) Status() (*emby.Status, error) {
	return ss.s, nil
}
