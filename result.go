package qqapi

type QQResult interface {
	GetErr() int
	GetErrMessage() string
}

type BaseResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type LoginResult struct {
	BaseResult
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

func (l *LoginResult) GetErr() int {
	return l.ErrCode
}

func (l *LoginResult) GetErrMessage() string {
	return l.ErrMsg
}

type AccessTokenResult struct {
	BaseResult
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (l *AccessTokenResult) GetErr() int {
	return l.ErrCode
}

func (l *AccessTokenResult) GetErrMessage() string {
	return l.ErrMsg
}

type MsgCheckResult struct {
	BaseResult
}

func (l *MsgCheckResult) GetErr() int {
	return l.ErrCode
}

func (l *MsgCheckResult) GetErrMessage() string {
	return l.ErrMsg
}

type ImgCheckResult struct {
	BaseResult
}

func (l *ImgCheckResult) GetErr() int {
	return l.ErrCode
}

func (l *ImgCheckResult) GetErrMessage() string {
	return l.ErrMsg
}

type MediaCheckAsyncResult struct {
	BaseResult
	TraceId string `json:"trace_id"`
}

func (l *MediaCheckAsyncResult) GetErr() int {
	return l.ErrCode
}

func (l *MediaCheckAsyncResult) GetErrMessage() string {
	return l.ErrMsg
}
