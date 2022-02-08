package emby

import (
	"fmt"
	"net/http"
)

type SessionInfo struct {
	PlayState             *PlayerStateInfo `json:"PlayState"`
	PlayableMediaTypes    []string         `json:"PlayableMediaTypes"`
	PlaylistItemId        string           `json:"PlaylistItemId"`
	PlaylistIndex         int32            `json:"PlaylistIndex"`
	PlaylistLength        int32            `json:"PlaylistLength"`
	Id                    string           `json:"Id"`
	ServerId              string           `json:"ServerId"`
	UserId                string           `json:"UserId"`
	UserName              string           `json:"UserName"`
	UserPrimaryImageTag   string           `json:"UserPrimaryImageTag"`
	Client                string           `json:"Client"`
	LastActivityDate      string           `json:"LastActivityDate"`
	DeviceName            string           `json:"DeviceName"`
	DeviceType            string           `json:"DeviceType"`
	NowPlayingItem        BaseItemDto      `json:"NowPlayingItem"`
	DeviceId              string           `json:"DeviceId"`
	ApplicationVersion    string           `json:"ApplicationVersion"`
	AppIconUrl            string           `json:"AppIconUrl"`
	SupportedCommands     []string         `json:"SupportedCommands"`
	TranscodingInfo       *TranscodingInfo `json:"TranscodingInfo"`
	SupportsRemoteControl bool             `json:"SupportsRemoteControl"`
}

type BaseItemDto struct {
	MediaStreams []*Mediastream
}

type Mediastream struct {
	Codec                  string
	Language               string
	Comment                string
	TimeBase               string
	CodecTimeBase          string
	Title                  string
	Extradata              string
	VideoRange             string
	DisplayTitle           string
	DisplayLanguage        string
	NalLengthSize          string
	IsInterlaced           bool
	IsAVC                  bool
	BitRate                int32
	BitDepth               int32
	RefFrames              int32
	PacketLength           int32
	Channels               int32
	SampleRate             int32
	IsDefault              bool
	IsForced               bool
	Height                 int32
	Width                  int32
	AverageFrameRate       float32
	RealFrameRate          float32
	Profile                string
	Type                   string
	AspectRatio            string
	IsExternal             bool
	IsTextSubtitleStream   bool
	SupportsExternalStream bool
}

type PlayerStateInfo struct {
	PositionTicks       int64   `json:"PositionTicks"`
	CanSeek             bool    `json:"CanSeek"`
	IsPaused            bool    `json:"IsPaused"`
	IsMuted             bool    `json:"IsMuted"`
	VolumeLevel         int32   `json:"VolumeLevel"`
	AudioStreamIndex    int32   `json:"AudioStreamIndex"`
	SubtitleStreamIndex int32   `json:"SubtitleStreamIndex"`
	MediaSourceId       string  `json:"MediaSourceId"`
	PlayMethod          string  `json:"PlayMethod"`
	RepeatMode          string  `json:"RepeatMode"`
	SubtitleOffset      int32   `json:"SubtitleOffset"`
	PlaybackRate        float32 `json:"PlaybackRate"`
}

type TranscodingInfo struct {
	AudioCodec                    string   `json:"AudioCodec"`
	VideoCodec                    string   `json:"VideoCodec"`
	SubProtocol                   string   `json:"SubProtocol"`
	Container                     string   `json:"Container"`
	IsVideoDirect                 bool     `json:"IsVideoDirect"`
	IsAudioDirect                 bool     `json:"IsAudioDirect"`
	Bitrate                       int32    `json:"Bitrate"`
	Framerate                     float32  `json:"Framerate"`
	CompletionPercentage          float32  `json:"CompletionPercentage"`
	TranscodingPositionTicks      float32  `json:"TranscodingPositionTicks"`
	TranscodingStartPositionTicks float32  `json:"TranscodingStartPositionTicks"`
	Width                         int32    `json:"Width"`
	Height                        int32    `json:"Height"`
	AudioChannels                 int32    `json:"AudioChannels"`
	TranscodeReasons              []string `json:"TranscodeReasons"`
	CurrentCpuUsage               float32  `json:"CurrentCpuUsage"`
	AverageCpuUsage               float32  `json:"AverageCpuUsage"`
	CurrentThrottle               int32    `json:"CurrentThrottle"`
	VideoDecoder                  string   `json:"VideoDecoder"`
	VideoDecoderIsHardware        bool     `json:"VideoDecoderIsHardware"`
	VideoDecoderMediaType         string   `json:"VideoDecoderMediaType"`
	VideoDecoderHwAccel           string   `json:"VideoDecoderHwAccel"`
	VideoEncoder                  string   `json:"VideoEncoder"`
	VideoEncoderIsHardware        bool     `json:"VideoEncoderIsHardware"`
	VideoEncoderMediaType         string   `json:"VideoEncoderMediaType"`
	VideoEncoderHwAccel           string   `json:"VideoEncoderHwAccel"`
}

func (c *Client) GetSessions() (*[]SessionInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Sessions", c.URL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(c.ctx)

	res := []SessionInfo{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
