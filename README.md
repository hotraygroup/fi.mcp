# fi.mcp

金融数据 MCP (Model Context Protocol) 服务，提供中国A股和加密货币市场数据查询工具。

## 功能

### Snowball (雪球) 工具

| 工具名 | 功能描述 | 必需参数 |
|--------|----------|----------|
| `suggest` | 根据公司名称或简称搜索股票代码 | `company` |
| `kline` | 获取K线数据 | `symbol` |
| `indicator` | 获取主要财务指标 | `symbol` |
| `income` | 获取利润表数据 | `symbol` |
| `balance` | 获取资产负债表数据 | `symbol` |
| `cashflow` | 获取现金流量表数据 | `symbol` |
| `quote` | 获取实时行情数据 | `symbol` |

### OKX 工具 (已禁用)

| 工具名 | 功能描述 |
|--------|----------|
| `candles` | 获取加密货币K线数据 |

### Akshare 工具 (已禁用)

| 工具名 | 功能描述 |
|--------|----------|
| `stock_zh_a_hist` | 东方财富A股历史数据 |
| `stock_cyq_em` | 筹码分布数据 |
| `stock_individual_fund_flow` | 个股资金流向 |
| 更多工具... | 见 `internal/logic/akshare/` 目录 |

## 安装

```bash
# 克隆项目
git clone https://github.com/hotraygroup/fi.mcp.git
cd fi.mcp

# 编译
go build -o fi.mcp main.go
```

## 配置

配置文件位于 `etc/fi.mcp.yaml`：

```yaml
name: fi.mcp
host: localhost
port: 8080
mcp:
  name: fi.mcp
  messageTimeout: 30s

Proxy:
  Enable: true
  Socks5: "100.115.184.183:1080"  # SOCKS5代理地址

DataSource:
  UserAgent: "Mozilla/5.0..."
  Snowball:
    IndexURL: "https://xueqiu.com/"
    QuoteURL: "https://stock.xueqiu.com/v5/stock/quote.json?symbol=%s&extend=detail"
    # 更多配置...
```

### 代理配置

项目支持SOCKS5代理访问雪球API，详见 [docs/PROXY_CONFIG.md](docs/PROXY_CONFIG.md)。

## 运行

```bash
./fi.mcp -f etc/fi.mcp.yaml
```

## 工具详细说明

### suggest - 股票代码搜索

**参数：**
- `company` (必需): 公司名称或简称

**示例：**
```json
{
  "company": "茅台"
}
```

**返回：**
- 匹配的股票代码列表

---

### quote - 实时行情

**参数：**
- `symbol` (必需): 股票代码，格式如 `SH600519` 或 `SZ000001`

**返回字段：**
| 字段 | 说明 |
|------|------|
| current | 当前价 |
| chg | 涨跌额 |
| percent | 涨跌幅 |
| open | 开盘价 |
| high | 最高价 |
| low | 最低价 |
| volume | 成交量 |
| amount | 成交额 |
| pe_ttm | 市盈率TTM |
| pb | 市净率 |
| market_capital | 总市值 |
| 更多字段... | |

---

### kline - K线数据

**参数：**
- `symbol` (必需): 股票代码
- `period` (可选): 周期，默认 `day`
- `count` (可选): 每次获取数量，默认 `284`
- `days` (可选): 最近多少天，默认 `360`

**返回：**
- K线数据列表，包含 timestamp、open、close、high、low、volume、amount

---

### indicator - 财务指标

**参数：**
- `symbol` (必需): 股票代码
- `count` (可选): 报告期数，默认 `5`
- `timestamp` (可选): 时间戳

**返回：**
- ROE、EPS、毛利率、资产负债率等财务指标

---

### income - 利润表

**参数：**
- `symbol` (必需): 股票代码

**返回：**
- 营业收入、净利润、营业成本等利润表数据

---

### balance - 资产负债表

**参数：**
- `symbol` (必需): 股票代码

**返回：**
- 总资产、总负债、股东权益等资产负债表数据

---

### cashflow - 现金流量表

**参数：**
- `symbol` (必需): 股票代码

**返回：**
- 经营活动、投资活动、筹资活动现金流数据

## 项目结构

```
fi.mcp/
├── main.go                 # 入口
├── etc/
│   └── fi.mcp.yaml        # 配置文件
├── internal/
│   ├── config/            # 配置结构
│   ├── svc/               # 服务上下文
│   ├── types/             # 数据类型定义
│   └── logic/
│       ├── logic.mcp.go   # MCP服务器和工具注册
│       ├── snowball/      # 雪球工具实现
│       ├── okx/           # OKX工具实现
│       └── akshare/       # Akshare工具实现
└── docs/
    └── PROXY_CONFIG.md    # 代理配置文档
```

## 依赖

- [go-zero](https://github.com/zeromicro/go-zero) - 微服务框架，提供MCP支持
- [resty](https://github.com/go-resty/resty) - HTTP客户端

## License

MIT