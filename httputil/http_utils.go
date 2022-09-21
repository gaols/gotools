package httputil

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var gTransport *http.Transport

func init() {
	gTransport = &http.Transport{
		MaxIdleConns:        2000,
		IdleConnTimeout:     2 * time.Minute,
		TLSHandshakeTimeout: 10 * time.Second,
		MaxIdleConnsPerHost: 50,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 15 * time.Second,
		}).DialContext,
	}
}

type HttpCli struct {
	Client *http.Client
}

func NewHttpCli() *HttpCli {
	return NewHttpCli2(time.Second * 30)
}

func NewHttpCli2(d time.Duration) *HttpCli {
	cli := &HttpCli{
		Client: &http.Client{
			Transport: gTransport,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return nil
			},
			Timeout: d,
		},
	}
	return cli
}

func (cli *HttpCli) GetBytes(_url string, headers string) ([]byte, error) {
	return GetBytes(cli.Client, _url, parseHeaders(headers))
}

func (cli *HttpCli) GetStr(_url string, headers string) (string, error) {
	return GetStr(cli.Client, _url, parseHeaders(headers))
}

func (cli *HttpCli) PostForm(_url string, headers string, form url.Values) ([]byte, error) {
	return PostForm(cli.Client, _url, parseHeaders(headers), form)
}

func (cli *HttpCli) PostJson(_url string, headers string, data string) ([]byte, error) {
	return PostJson(cli.Client, _url, parseHeaders(headers), data)
}

func parseHeaders(headers string) map[string]string {
	sep := regexp.MustCompile(`[\r\n]{1,2}`)
	headersArr := sep.Split(headers, -1)
	ret := make(map[string]string)
	for _, v := range headersArr {
		if item := strings.TrimSpace(v); item != "" {
			if pair := strings.SplitN(item, ":", 2); len(pair) > 1 {
				key := strings.TrimSpace(pair[0])
				val := strings.TrimSpace(pair[1])
				if key != "" && val != "" {
					ret[key] = val
				}
			}
		}
	}
	delete(ret, "Accept-Encoding")
	delete(ret, "accept-encoding")
	return ret
}

func GetBytes(cli *http.Client, _url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, _url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request (%s) error: %w", _url, err)
	}
	return doReq(cli, _url, headers, req)
}

func doReq(cli *http.Client, _url string, headers map[string]string, req *http.Request) ([]byte, error) {
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get request (%s) error: %w", _url, err)
	}
	body := resp.Body
	defer func() {
		_ = body.Close()
	}()
	ret, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("read response body of (%s) error: %w", _url, err)
	}
	return ret, nil
}

func GetStr(cli *http.Client, _url string, headers map[string]string) (string, error) {
	body, err := GetBytes(cli, _url, headers)
	return string(body), err
}

func PostForm(cli *http.Client, _url string, headers map[string]string, form url.Values) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, _url, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("create request (%s) error: %w", _url, err)
	}
	return doReq(cli, _url, headers, req)
}

func PostJson(cli *http.Client, _url string, headers map[string]string, data string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, _url, strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("create request (%s) error: %w", _url, err)
	}
	return doReq(cli, _url, headers, req)
}
