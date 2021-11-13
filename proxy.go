package rest

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"time"
)

type ProxyClient struct {
	client *http.Client
}

func NewProxyClient(proxy func(*http.Request) (*url.URL, error)) *ProxyClient {
	return &ProxyClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: proxy,
			},
			Timeout: time.Second * 5,
		},
	}
}

func NewHTTP2ProxyClient(p string) *ProxyClient {
	dialer, err := proxy.SOCKS5("tcp", p, nil, proxy.Direct)
	if err != nil {
		panic(err)
	}
	proxyDialer := func(network string, addr string, cfg *tls.Config) (c net.Conn, e error) {
		c, e = dialer.Dial(network, addr)
		return
	}
	return &ProxyClient{
		client: &http.Client{
			Transport: &http2.Transport{
				DialTLS: proxyDialer,
			},
			Timeout: time.Second * 5,
		},
	}
}

func (c *ProxyClient) Do(r *http.Request) (*http.Response, error) {
	// Prevent proxy tunnelling (the proxy will upgrade this to https)
	r.URL.Scheme = "http"
	return c.client.Do(r)
}
