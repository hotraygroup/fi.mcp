package config

import "github.com/zeromicro/go-zero/mcp"

type Config struct {
	mcp.McpConf
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
	}
}
