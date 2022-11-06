package http

import (
	"crypto/tls"
	"github.com/antlabs/pcurl"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
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
