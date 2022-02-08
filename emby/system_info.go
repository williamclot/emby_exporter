package emby

import (
	"fmt"
	"net/http"
)

type SystemInfo struct {
	SystemUpdateLevel                    string `json:"SystemUpdateLevel"`
	OperatingSystemDisplayName           string `json:"OperatingSystemDisplayName"`
	HasPendingRestart                    bool   `json:"HasPendingRestart"`
	IsShuttingDown                       bool   `json:"IsShuttingDown"`
	OperatingSystem                      string `json:"OperatingSystem"`
	SupportsLibraryMonitor               bool   `json:"SupportsLibraryMonitor"`
	SupportsLocalPortConfiguration       bool   `json:"SupportsLocalPortConfiguration"`
	WebSocketPortNumber                  int    `json:"WebSocketPortNumber"`
	CanSelfUpdate                        bool   `json:"CanSelfUpdate"`
	CanLaunchWebBrowser                  bool   `json:"CanLaunchWebBrowser"`
	ProgramDataPath                      string `json:"ProgramDataPath"`
	ItemsByNamePath                      string `json:"ItemsByNamePath"`
	CachePath                            string `json:"CachePath"`
	LogPath                              string `json:"LogPath"`
	InternalMetadataPath                 string `json:"InternalMetadataPath"`
	TranscodingTempPath                  string `json:"TranscodingTempPath"`
	HttpServerPortNumber                 int    `json:"HttpServerPortNumber"`
	SupportsHttps                        bool   `json:"SupportsHttps"`
	HttpsPortNumber                      int    `json:"HttpsPortNumber"`
	HasUpdateAvailable                   bool   `json:"HasUpdateAvailable"`
	SupportsAutoRunAtStartup             bool   `json:"SupportsAutoRunAtStartup"`
	HardwareAccelerationRequiresPremiere bool   `json:"HardwareAccelerationRequiresPremiere"`
	LocalAddress                         string `json:"LocalAddress"`
	WanAddress                           string `json:"WanAddress"`
	ServerName                           string `json:"ServerName"`
	Version                              string `json:"Version"`
	Id                                   string `json:"Id"`
}

func (c *Client) GetSystemInfo() (*SystemInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/System/Info", c.URL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(c.ctx)

	res := SystemInfo{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
