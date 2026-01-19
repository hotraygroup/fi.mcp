package snowball

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/proxy"
)

// ProxyConfig 代理配置接口
type ProxyConfig interface {
	IsProxyEnabled() bool
	GetSocks5Addr() string
}

// NewClient 创建一个新的 resty 客户端
// 如果启用代理，则配置 SOCKS5 代理
func NewClient(proxyEnabled bool, socks5Addr string) *resty.Client {
	client := resty.New()

	if proxyEnabled && socks5Addr != "" {
		// 创建 SOCKS5 代理拨号器
		dialer, err := proxy.SOCKS5("tcp", socks5Addr, nil, proxy.Direct)
		if err != nil {
			// 如果代理配置失败，记录错误但继续使用直连
			// 这样不会因为代理问题导致整个服务不可用
			return client
		}

		// 配置 HTTP 传输使用 SOCKS5 代理
		httpTransport := &http.Transport{
			Dial: dialer.Dial,
		}

		client.SetTransport(httpTransport)
	}

	return client
}

// NewClientWithConfig 使用配置对象创建 resty 客户端
func NewClientWithConfig(cfg ProxyConfig) *resty.Client {
	return NewClient(cfg.IsProxyEnabled(), cfg.GetSocks5Addr())
}
