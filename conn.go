package coscloud

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Conn struct {
	config    *Config
	bucketUrl string
}

func (conn Conn) Do(method, path, urlParams string, headers map[string]string, data io.Reader) ([]byte, error) {

	urlStr := conn.getURL(path, urlParams)
	return conn.doRequest(method, urlStr, headers, data)
}

// Build URL
func (conn Conn) getURL(path, params string) string {

	path = conn.bucketUrl + path
	if params != "" {
		path += "?" + params
	}
	return path
}

//Private
func (conn Conn) doRequest(method, urlStr string, headers map[string]string, body io.Reader) (respData []byte, err error) {

	req, err := http.NewRequest(method, urlStr, body)

	if nil != err {
		return
	}

	req.Header.Set("HOST", COSCLOUD_DOMAIN)
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("User-Agent", conn.config.UserAgent)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	//set timeOut
	tr := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			deadline := time.Now().Add(conn.config.Timeout)
			c, err := net.DialTimeout(netw, addr, conn.config.Timeout)
			if err != nil {
				return nil, err
			}
			c.SetDeadline(deadline)
			return c, nil
		},
	}
	//https
	if "https" == req.URL.Scheme {
		if nil != conn.config.Pool { // cert
			tr.TLSClientConfig = &tls.Config{RootCAs: conn.config.Pool}
		} else { // no cert
			tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}

	}

	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if nil != err {
		return
	}
	defer resp.Body.Close()

	respData, err = ioutil.ReadAll(resp.Body)
	if nil != err {
		return
	}

	return
}
