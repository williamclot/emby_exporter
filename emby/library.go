package emby

import (
	"fmt"
	"net/http"
)

type LibraryItems struct {
	MovieCount      int `json:"MovieCount"`
	SeriesCount     int `json:"SeriesCount"`
	EpisodeCount    int `json:"EpisodeCount"`
	GameCount       int `json:"GameCount"`
	ArtistCount     int `json:"ArtistCount"`
	ProgramCount    int `json:"ProgramCount"`
	GameSystemCount int `json:"GameSystemCount"`
	TrailerCount    int `json:"TrailerCount"`
	SongCount       int `json:"SongCount"`
	AlbumCount      int `json:"AlbumCount"`
	MusicVideoCount int `json:"MusicVideoCount"`
	BoxSetCount     int `json:"BoxSetCount"`
	BookCount       int `json:"BookCount"`
	ItemCount       int `json:"ItemCount"`
}

func (c *Client) GetLibraryItems() (*LibraryItems, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Items/Counts", c.URL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(c.ctx)

	res := LibraryItems{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
