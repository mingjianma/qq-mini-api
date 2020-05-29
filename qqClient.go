package qqapi

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type QQClient struct {
	appID       string
	appSecret   string
	http        *httpHelper
	logger      *zap.Logger
	accessToken string
	expired     int64
	accessUrl   string
}

func NewQQClient(appID string, appSecret string, logger *zap.Logger) *QQClient {
	client := &http.Client{Timeout: 5 * time.Second}
	url := fmt.Sprintf("%s%s?grant_type=client_credential&appid=%s&secret=%s", baseUrl, accessTokenUrl, appID, appSecret)
	httpHelper := NewHttpHelper(logger, client)
	return &QQClient{
		appID:     appID,
		appSecret: appSecret,
		http:      httpHelper,
		logger:    logger,
		accessUrl: url,
	}
}

func (c *QQClient) GetAccessToken() (accessToken string, err error) {
	now := time.Now().Unix()
	//判断accessToken是否存在和过期，不存在和过期需要重新获取
	if c.accessToken == "" || c.expired <= now {
		result := &accessTokenResult{}
		err = c.http.Get(c.accessUrl, result)
		if err != nil {
			return
		}
		accessToken, c.accessToken = result.AccessToken, result.AccessToken
		c.expired = now + result.ExpiresIn
	} else {
		accessToken = c.accessToken
	}
	return
}
