package http

import (
	"crypto/tls"
	"encoding/json"
	"github.com/antlabs/pcurl"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	clinetTransport *http.Transport
	Client          *http.Client
)

func init() {
	clinetTransport = &http.Transport{DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
		ResponseHeaderTimeout: 30 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ExpectContinueTimeout: 5 * time.Second,
		MaxIdleConns:          0,
		MaxIdleConnsPerHost:   100,
	}
	Client = &http.Client{Transport: clinetTransport}

}

func InitProxyClient(proxy string) {
	if len(proxy) > 0 {
		proxyUrl, err := url.Parse("socks5://" + proxy)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		clinetTransport.Proxy = http.ProxyURL(proxyUrl)
	}
}

func HttpDefaultDo(req *http.Request) (*http.Response, error) {
	req.Header.Set("Connection", "keep-alive")
	resp, err := Client.Do(req)
	return resp, err
}

func HttpDefaultBytes(req *http.Request) (code int, bodystr []byte) {
	resp, err := HttpDefaultDo(req)
	if err != nil {
		return -1, []byte(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, []byte(err.Error())
	}
	return resp.StatusCode, body
}

func HttpDefaultString(req *http.Request) (code int, bodystr string) {
	code, body := HttpDefaultBytes(req)
	return code, string(body)
}

func CurlToRequest(curlcmd string) (error, *http.Request) {
	req, err := pcurl.ParseAndRequest(curlcmd)
	return err, req
}

func HttpGet(url string) string {
	resp, err := Client.Get(url)
	if err != nil {
		logrus.Errorln(err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func HttpJSON(url string, jsonObj interface{}) {
	response := HttpGet(url)
	json.Unmarshal([]byte(response), &jsonObj)
}

func HttpWithReqJSON(url string, requestBody string, jsonObj interface{}) {
	if len(strings.TrimSpace(requestBody)) == 0 {
		requestBody = ""
	}
	request := strings.NewReader(requestBody)
	resp, err := Client.Post(url, "application/json", request)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	response := string(body)
	json.Unmarshal([]byte(response), &jsonObj)
}
