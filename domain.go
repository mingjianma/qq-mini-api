package qqapi

type msgCheckPost struct {
	AccessToken string `json:"access_token"`
	AppID       string `json:"appid"`
	Content     string `json:"content"`
}

func NewMsgCheckPost(accessToken string, appID string, content string) *msgCheckPost {
	return &msgCheckPost{
		AccessToken: accessToken,
		AppID:       appID,
		Content:     content,
	}
}

type mediaAsync struct {
	AccessToken string `json:"access_token"`
	AppID       string `json:"appid"`
	MediaUrl    string `json:"media_url"`
	MediaType   string `json:"media_type"`
}

func NewMediaAsync(accessToken string, appID string, mediaUrl string, mediaType string) *mediaAsync {
	return &mediaAsync{
		AccessToken: accessToken,
		AppID:       appID,
		MediaUrl:    mediaUrl,
		MediaType:   mediaType,
	}
}
