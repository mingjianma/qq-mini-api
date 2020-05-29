package qqapi

import (
	"fmt"
	"go.uber.org/zap"
)

func (c *QQClient) GetUserInfo(jscode string) (*LoginResult, error) {
	endpoint := fmt.Sprintf("%s%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", BaseUrl, LoginUrl, c.appID, c.appSecret, jscode)
	c.logger.Debug("获取用户信息接口调用前：", zap.String("endpoint", endpoint))
	result := &LoginResult{}
	err := c.http.Get(endpoint, result)

	if err != nil {
		c.logger.Error("GET接口调用失败", zap.Any("err:", err), zap.String("endpoint", endpoint))
		return nil, err
	}

	return result, nil
}
