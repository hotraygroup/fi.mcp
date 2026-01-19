package config

import "github.com/zeromicro/go-zero/mcp"

type Config struct {
	mcp.McpConf
	Proxy struct {
		Enable bool
		Socks5 string
	}
	DataSource struct {
		UserAgent string
		Snowball  struct {
			IndexURL     string
			CookieURL    string
			SuggestURL   string
			KlineURL     string
			IndicatorURL string //ZYCWZB zhuyaocaiwuzhibiao
			IncomeURL    string //GSLRB gongsilirunbiao
			BalanceURL   string //ZCFZB zichanfuzhaibiao
			CashFlowURL  string //XJLLB xianjinliuliangbiao
			QuoteURL     string
			ListURL      string
			Symbol       struct {
				SH string
			}
		}
		OKX struct {
			Host    string
			Candles string
		}
		Akshare struct {
			Host string
		}
	}
}

// IsProxyEnabled 返回代理是否启用
func (c *Config) IsProxyEnabled() bool {
	return c.Proxy.Enable
}

// GetSocks5Addr 返回 SOCKS5 代理地址
func (c *Config) GetSocks5Addr() string {
	return c.Proxy.Socks5
}
