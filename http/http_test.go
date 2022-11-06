package http

import (
	"net/http"
	"testing"
)

func Test_InitProxyClient(t *testing.T) {
	InitProxyClient("127.0.0.1:1080")
	t.Log("init proxy success")
}
func Test_HttpDefaultString(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://myip.ipip.net/", nil)
	if err != nil {
		t.Error(err)
	}
	code, str := HttpDefaultString(req)
	t.Log(code, str)
}

func Test_pcurl(t *testing.T) {
	curlstr := `curl 'http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp' \
  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
  -H 'Accept-Language: zh-CN,zh;q=0.9' \
  -H 'Cache-Control: max-age=0' \
  -H 'Connection: keep-alive' \
  -H 'Upgrade-Insecure-Requests: 1' \
  -H 'User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1' \
  --compressed`
	err, req := CurlToRequest(curlstr)
	if err != nil {
		t.Error(err)
	}
	code, str := HttpDefaultString(req)
	t.Log(code, str)

}
