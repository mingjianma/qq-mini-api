package qqapi

type MsgCheckPost struct {
	AccessToken string `json:"access_token"`
	AppID       string `json:"appid"`
	Content     string `json:"content"`
}

func NewMsgCheckPost(accessToken string, appID string, content string) *MsgCheckPost {
	return &MsgCheckPost{
		AccessToken: accessToken,
		AppID:       appID,
		Content:     content,
	}
}

type MediaAsync struct {
	AccessToken string `json:"access_token"`
	AppID       string `json:"appid"`
	MediaUrl    string `json:"media_url"`
	MediaType   string `json:"media_type"`
}

func NewMediaAsync(accessToken string, appID string, mediaUrl string, mediaType string) *MediaAsync {
	return &MediaAsync{
		AccessToken: accessToken,
		AppID:       appID,
		MediaUrl:    mediaUrl,
		MediaType:   mediaType,
	}
}
