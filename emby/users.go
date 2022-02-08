package emby

import (
	"fmt"
	"net/http"
)

type UserList struct {
	TotalCount int    `json:"TotalRecordCount"`
	Users      []User `json:"Items"`
}

type User struct {
	Name                      string            `json:"Name"`
	ConnectUserName           string            `json:"ConnectUserName"`
	ConnectLinkType           string            `json:"ConnectLinkType"`
	Id                        string            `json:"Id"`
	HasPassword               bool              `json:"HasPassword"`
	HasConfiguredPassword     bool              `json:"HasConfiguredPassword"`
	HasConfiguredEasyPassword bool              `json:"HasConfiguredEasyPassword"`
	EnableAutoLogin           bool              `json:"EnableAutoLogin"`
	LastLoginDate             string            `json:"LastLoginDate"`
	LastActivityDate          string            `json:"LastActivityDate"`
	Configuration             UserConfiguration `json:"Configuration"`
	Policy                    UserPolicy        `json:"Policy"`
}

type UserConfiguration struct {
	AudioLanguagePreference    string   `json:"AudioLanguagePreference"`
	PlayDefaultAudioTrack      bool     `json:"PlayDefaultAudioTrack"`
	SubtitleLanguagePreference string   `json:"SubtitleLanguagePreference"`
	DisplayMissingEpisodes     bool     `json:"DisplayMissingEpisodes"`
	SubtitleMode               string   `json:"SubtitleMode"`
	EnableLocalPassword        bool     `json:"EnableLocalPassword"`
	OrderedViews               []string `json:"OrderedViews"`
	LatestItemsExcludes        []string `json:"LatestItemsExcludes"`
	MyMediaExcludes            []string `json:"MyMediaExcludes"`
	HidePlayedInLatest         bool     `json:"HidePlayedInLatest"`
	RememberAudioSelections    bool     `json:"RememberAudioSelections"`
	RememberSubtitleSelections bool     `json:"RememberSubtitleSelections"`
	EnableNextEpisodeAutoPlay  bool     `json:"EnableNextEpisodeAutoPlay"`
}

type UserPolicy struct {
	IsAdministrator                  bool     `json:"IsAdministrator"`
	IsHidden                         bool     `json:"IsHidden"`
	IsHiddenRemotely                 bool     `json:"IsHiddenRemotely"`
	IsHiddenFromUnusedDevices        bool     `json:"IsHiddenFromUnusedDevices"`
	IsDisabled                       bool     `json:"IsDisabled"`
	MaxParentalRating                int      `json:"MaxParentalRating"`
	BlockedTags                      []string `json:"BlockedTags"`
	IsTagBlockingModeInclusive       bool     `json:"IsTagBlockingModeInclusive"`
	EnableUserPreferenceAccess       bool     `json:"EnableUserPreferenceAccess"`
	BlockUnratedItems                []string `json:"BlockUnratedItems"`
	EnableRemoteControlOfOtherUsers  bool     `json:"EnableRemoteControlOfOtherUsers"`
	EnableSharedDeviceControl        bool     `json:"EnableSharedDeviceControl"`
	EnableRemoteAccess               bool     `json:"EnableRemoteAccess"`
	EnableLiveTvManagement           bool     `json:"EnableLiveTvManagement"`
	EnableLiveTvAccess               bool     `json:"EnableLiveTvAccess"`
	EnableMediaPlayback              bool     `json:"EnableMediaPlayback"`
	EnableAudioPlaybackTranscoding   bool     `json:"EnableAudioPlaybackTranscoding"`
	EnableVideoPlaybackTranscoding   bool     `json:"EnableVideoPlaybackTranscoding"`
	EnablePlaybackRemuxing           bool     `json:"EnablePlaybackRemuxing"`
	EnableContentDeletion            bool     `json:"EnableContentDeletion"`
	EnableContentDeletionFromFolders []string `json:"EnableContentDeletionFromFolders"`
	EnableContentDownloading         bool     `json:"EnableContentDownloading"`
	EnableSubtitleDownloading        bool     `json:"EnableSubtitleDownloading"`
	EnableSubtitleManagement         bool     `json:"EnableSubtitleManagement"`
	EnableSyncTranscoding            bool     `json:"EnableSyncTranscoding"`
	EnableMediaConversion            bool     `json:"EnableMediaConversion"`
	EnabledChannels                  []string `json:"EnabledChannels"`
	EnableAllChannels                bool     `json:"EnableAllChannels"`
	EnabledFolders                   []string `json:"EnabledFolders"`
	EnableAllFolders                 bool     `json:"EnableAllFolders"`
	InvalidLoginAttemptCount         int      `json:"InvalidLoginAttemptCount"`
	EnablePublicSharing              bool     `json:"EnablePublicSharing"`
	BlockedMediaFolders              []string `json:"BlockedMediaFolders"`
	BlockedChannels                  []string `json:"BlockedChannels"`
	RemoteClientBitrateLimit         int      `json:"RemoteClientBitrateLimit"`
	AuthenticationProviderId         string   `json:"AuthenticationProviderId"`
	ExcludedSubFolders               []string `json:"ExcludedSubFolders"`
	SimultaneousStreamLimit          int      `json:"SimultaneousStreamLimit"`
	EnabledDevices                   []string `json:"EnabledDevices"`
	EnableAllDevices                 bool     `json:"EnableAllDevices"`
}

func (c *Client) GetUsers() (*UserList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Users/Query", c.URL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(c.ctx)

	res := UserList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
