package emby

import (
	"fmt"
	"net/http"
)

type DeviceList struct {
	TotalCount int      `json:"TotalRecordCount"`
	Devices    []Device `json:"Items"`
}

type Device struct {
	Name             string `json:"Name"`
	ID               string `json:"Id"`
	LastUserName     string `json:"LastUserName"`
	AppName          string `json:"AppName"`
	AppVersion       string `json:"AppVersion"`
	LastUserId       string `json:"LastUserId"`
	DateLastActivity string `json:"DateLastActivity"`
}

func (c *Client) GetDevices() (*DeviceList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Devices", c.URL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(c.ctx)

	res := DeviceList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
