package http

import (
	"net/http"
	"testing"
)

//	func Test_InitProxyClient(t *testing.T) {
//		InitProxyClient("127.0.0.1:1080")
//		t.Log("init proxy success")
//	}
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

func TestHttpGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Request httpbin",
			args: args{
				url: "https://httpbin.org/base64/SFRUUEJJTiBpcyBhd2Vzb21l",
			},
			want: "HTTPBIN is awesome",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HttpGet(tt.args.url); got != tt.want {
				t.Errorf("HttpGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestHttpJSONResponse struct {
	Slideshow struct {
		Date   string `json:"date" `
		Slides []struct {
			Title string `json:"title" `
			Type  string `json:"type" `
		} `json:"slides" `
		Author string `json:"author" `
		Title  string `json:"title" `
	} `json:"slideshow" `
}

func TestHttpJSON(t *testing.T) {
	type args struct {
		url     string
		jsonObj TestHttpJSONResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Request httpbin",
			args: args{
				url:     "https://httpbin.org/json",
				jsonObj: TestHttpJSONResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpJSON(tt.args.url, &tt.args.jsonObj)
			t.Log("TestHttpJSON() = ", tt.args.jsonObj)
		})
	}
}

type TestHttpWithReqJSONResponse struct {
	Args struct {
	} `json:"args" `
	Headers struct {
		Origin          string `json:"Origin" `
		SecChUa         string `json:"Sec-Ch-Ua" `
		Accept          string `json:"Accept" `
		SecChUaPlatform string `json:"Sec-Ch-Ua-Platform" `
		Referer         string `json:"Referer" `
		UserAgent       string `json:"User-Agent" `
		SecFetchDest    string `json:"Sec-Fetch-Dest" `
		SecFetchSite    string `json:"Sec-Fetch-Site" `
		Host            string `json:"Host" `
		AcceptEncoding  string `json:"Accept-Encoding" `
		XAmznTraceID    string `json:"X-Amzn-Trace-Id" `
		SecFetchMode    string `json:"Sec-Fetch-Mode" `
		AcceptLanguage  string `json:"Accept-Language" `
		ContentLength   string `json:"Content-Length" `
		SecChUaMobile   string `json:"Sec-Ch-Ua-Mobile" `
	} `json:"headers" `
	Data string `json:"data" `
	Form struct {
	} `json:"form" `
	Method string `json:"method" `
	Origin string `json:"origin" `
	Files  struct {
	} `json:"files" `
	Json interface{} `json:"json" `
	Url  string      `json:"url" `
}

func TestHttpWithReqJSON(t *testing.T) {
	type args struct {
		url         string
		requestBody string
		jsonObj     TestHttpWithReqJSONResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Request httpbin",
			args: args{
				url:         "https://httpbin.org/anything/abc",
				requestBody: "{\"value\":\"Hello World!!\"}",
				jsonObj:     TestHttpWithReqJSONResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpWithReqJSON(tt.args.url, tt.args.requestBody, &tt.args.jsonObj)
			t.Log("HttpWithReqJSON() = ", tt.args.jsonObj)
		})
	}
}
