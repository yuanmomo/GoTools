package json

import (
	"encoding/json"
	myhttp "github.com/bianzhifu/GoTools/http"
	"github.com/buger/jsonparser"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GetTimestampResp struct {
	API  string   `json:"api"`
	V    string   `json:"v"`
	Ret  []string `json:"ret"`
	Data struct {
		T string `json:"t"`
	} `json:"data"`
}

func GetTimestamp() (resp GetTimestampResp) {
	req, err := http.NewRequest(http.MethodGet, "http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp", nil)
	if err != nil {
		logrus.Errorln(req)
		return
	}
	code, body := myhttp.HttpDefaultBytes(req)
	if code == http.StatusOK {
		err = json.Unmarshal(body, &resp)
		if err != nil {
			logrus.Errorln(req)
			return
		}
	}
	return
}

func GetTimestampT() (t string) {
	req, err := http.NewRequest(http.MethodGet, "http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp", nil)
	if err != nil {
		logrus.Errorln(req)
		return
	}
	code, body := myhttp.HttpDefaultBytes(req)
	if code == http.StatusOK {
		t, err = jsonparser.GetString(body, "data", "t")
		if err != nil {
			logrus.Errorln(err)
		}
	}
	return
}

func GetTimestampMap() (resp map[string]interface{}) {
	req, err := http.NewRequest(http.MethodGet, "http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp", nil)
	if err != nil {
		logrus.Errorln(req)
		return
	}
	code, body := myhttp.HttpDefaultBytes(req)
	if code == http.StatusOK {
		err = json.Unmarshal(body, &resp)
		if err != nil {
			logrus.Errorln(req)
			return
		}
	}
	return
}
