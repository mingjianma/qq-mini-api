package qqapi

import (
	"go.uber.org/zap"
	"testing"
)

var (
	appID     = "your appid"
	appSecret = "your appSecret"
	logger    = zap.L()
	qqClient  = NewQQClient(appID, appSecret, logger)
)

func TestQQClient_MsgCheck(t *testing.T) {
	ok, err := qqClient.MsgCheck("特3456书yuuo莞6543李zxcz蒜7782法fgnv级")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if ok == true {
		t.Error("结果1异常")
		t.FailNow()
	}

	ok, err = qqClient.MsgCheck("完2347全dfji试3726测asad感3847知qwez到")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if ok == true {
		t.Error("结果2异常")
		t.FailNow()
	}

	ok, err = qqClient.MsgCheck("123")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if ok == false {
		t.Error("结果3异常")
		t.FailNow()
	}
}

func TestQQClient_ImgCheck(t *testing.T) {
	ok, err := qqClient.ImgCheck("http://n.sinaimg.cn/news/transform/700/w1000h500/20200527/6d68-iufmpmn0212173.jpg")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(ok)
}

func TestQQClient_MediaCheckAsync(t *testing.T) {
	traceID, err := qqClient.MediaCheckAsync("http://n.sinaimg.cn/news/transform/700/w1000h500/20200527/6d68-iufmpmn0212173.jpg", 2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(traceID)
}
