package qqapi

import (
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"mime/multipart"
	"strconv"
)

//文本检查
func (c *QQClient) MsgCheck(msg string) (ok bool, err error) {
	//获取token
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s%s?access_token=%s", baseUrl, msgCheckUrl, accessToken)
	data := newMsgCheckPost(accessToken, c.appID, msg)
	result := &msgCheckResult{}
	//发起api请求
	err = c.http.Post(url, data, "application/json", result)
	if err != nil {
		c.logger.Error("MsgCheck", zap.Any("err", err), zap.String("url", url), zap.Any("data", data), zap.Any("result", result))
		return
	}
	//获取检测结果
	if result.ErrCode == 0 {
		ok = true
	} else {
		c.logger.Debug("ImgCheck", zap.Any("result", result))
		ok = false
	}
	return
}

//图片同步检测 imgUrl为检测图片的url
func (c *QQClient) ImgCheck(imgUrl string) (ok bool, err error) {
	//获取token
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s%s?access_token=%s", baseUrl, imgCheckUrl, accessToken)
	//获取图片文件内容
	bin, err := c.http.GetPic(imgUrl)
	if err != nil {
		return
	}
	//生成format_data
	data := new(bytes.Buffer)
	w := multipart.NewWriter(data)

	err = w.WriteField("access_token", c.accessToken)
	if err != nil {
		return
	}
	err = w.WriteField("appid", c.appID)
	if err != nil {
		return
	}
	fw, err := w.CreateFormFile("media", "test")
	if err != nil {
		return
	}

	_, err = fw.Write(bin)
	if err != nil {
		return
	}
	//close不能defer，close以后才能把数据写入Buffer
	err = w.Close()
	if err != nil {
		return
	}

	result := &imgCheckResult{}
	//发起api请求
	err = c.http.FormatDataPost(url, data, w.FormDataContentType(), result)
	if err != nil {
		c.logger.Error("ImgCheck", zap.Any("err", err), zap.String("url", url), zap.Any("data", data), zap.Any("result", result))
		return
	}
	//获取检测结果
	if result.ErrCode == 0 {
		ok = true
	} else {
		c.logger.Debug("ImgCheck", zap.Any("result", result))
		ok = false
	}
	return
}

//异步校验图片或音频
// mediaType => 1:音频;2:图片
func (c *QQClient) MediaCheckAsync(mediaUrl string, mediaType int) (traceID string, err error) {
	//获取token
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s%s?access_token=%s", baseUrl, mediaCheckUrl, accessToken)
	data := newMediaAsync(accessToken, c.appID, mediaUrl, strconv.Itoa(mediaType))
	result := &mediaCheckAsyncResult{}
	//发起api请求
	err = c.http.Post(url, data, "application/json", result)
	if err != nil {
		c.logger.Error("MediaCheckAsync", zap.Any("err", err), zap.String("url", url), zap.Any("data", data), zap.Any("result", result))
		return
	}
	//获取检测结果
	if result.ErrCode == 0 {
		traceID = result.TraceId
	} else {
		c.logger.Debug("MediaCheckAsync", zap.Any("result", result))
		err = errors.New("MediaCheckRequestErr:" + result.ErrMsg)
	}
	return
}
