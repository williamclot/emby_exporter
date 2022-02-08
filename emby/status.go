package emby

import (
	"strings"
)

type Status struct {
	ServerName string
	Hostname   string
	Version    string

	DeviceCount int

	MovieCount   int
	SeriesCount  int
	EpisodeCount int

	UserCount        int
	FailedLoginCount map[string]int

	MediaStreamCount   int
	TranscodingCount   int
	TranscodingBitRate map[string]int
	VideoBitRate       map[string]int
	AudioChannels      map[string]int
	AudioBitRate       map[string]int
}

func (c *Client) Status() (*Status, error) {
	s := &Status{
		FailedLoginCount:   make(map[string]int),
		TranscodingBitRate: make(map[string]int),
		VideoBitRate:       make(map[string]int),
		AudioChannels:      make(map[string]int),
		AudioBitRate:       make(map[string]int),
	}

	sys, err := c.GetSystemInfo()
	if err != nil {
		return nil, err
	}

	s.ServerName = strings.ToLower(sys.ServerName)
	s.Hostname = sys.WanAddress
	s.Version = sys.Version

	devices, err := c.GetDevices()
	if err != nil {
		return nil, err
	}

	s.DeviceCount = devices.TotalCount

	libary, err := c.GetLibraryItems()
	if err != nil {
		return nil, err
	}

	s.MovieCount = libary.MovieCount
	s.SeriesCount = libary.SeriesCount
	s.EpisodeCount = libary.EpisodeCount

	users, err := c.GetUsers()
	if err != nil {
		return nil, err
	}

	s.UserCount = users.TotalCount
	for _, user := range users.Users {
		s.FailedLoginCount[strings.ToLower(user.Name)] = user.Policy.InvalidLoginAttemptCount
	}

	sessions, err := c.GetSessions()
	if err != nil {
		return nil, err
	}

	for _, session := range *sessions {
		if session.TranscodingInfo != nil {
			s.TranscodingCount++
			s.VideoBitRate[strings.ToLower(session.UserName)] += int(session.TranscodingInfo.Bitrate)
		}

		if len(session.NowPlayingItem.MediaStreams) > 0 {
			for _, stream := range session.NowPlayingItem.MediaStreams {
				if stream.Type == "Video" {
					s.MediaStreamCount++
					s.VideoBitRate[strings.ToLower(session.UserName)] += int(stream.BitRate)
				}

				if stream.Type == "Audio" {
					s.AudioChannels[strings.ToLower(session.UserName)] += int(stream.Channels)
					s.AudioBitRate[strings.ToLower(session.UserName)] += int(stream.BitRate)
				}
			}
		}
	}

	return s, nil
}
