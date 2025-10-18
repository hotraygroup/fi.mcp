package akshare

import (
	"context"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// stock_fund_flow_big_deal
func NewStockFundFlowBigDealTool(_mcp types.MCPProvider) mcp.Tool {
	var tool = mcp.Tool{
		Name:        "stock_fund_flow_big_deal",
		Description: "同花顺-数据中心-资金流向-大单追踪",

		Handler: func(ctx context.Context, params map[string]any) (any, error) {

			client := resty.New()
			url := fmt.Sprintf("%s/api/public/stock_fund_flow_big_deal",
				_mcp.GetServiceContext().Config.DataSource.Akshare.Host,
			)

			_mcp.GetLogger().Infof("url: %s", url)

			var result []map[string]any

			_, err := client.R().SetResult(&result).Get(url)
			if err != nil {
				return nil, err
			}
			return map[string]any{
				"items": result,
			}, nil
		},
	}
	return tool
}
