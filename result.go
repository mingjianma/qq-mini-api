package qqapi

type qqResult interface {
	GetErr() int
	GetErrMessage() string
}

type baseResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type loginResult struct {
	baseResult
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

func (l *loginResult) GetErr() int {
	return l.ErrCode
}

func (l *loginResult) GetErrMessage() string {
	return l.ErrMsg
}

type accessTokenResult struct {
	baseResult
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (l *accessTokenResult) GetErr() int {
	return l.ErrCode
}

func (l *accessTokenResult) GetErrMessage() string {
	return l.ErrMsg
}

type msgCheckResult struct {
	baseResult
}

func (l *msgCheckResult) GetErr() int {
	return l.ErrCode
}

func (l *msgCheckResult) GetErrMessage() string {
	return l.ErrMsg
}

type imgCheckResult struct {
	baseResult
}

func (l *imgCheckResult) GetErr() int {
	return l.ErrCode
}

func (l *imgCheckResult) GetErrMessage() string {
	return l.ErrMsg
}

type mediaCheckAsyncResult struct {
	baseResult
	TraceId string `json:"trace_id"`
}

func (l *mediaCheckAsyncResult) GetErr() int {
	return l.ErrCode
}

func (l *mediaCheckAsyncResult) GetErrMessage() string {
	return l.ErrMsg
}
