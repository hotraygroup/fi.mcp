# HTTP 代理配置说明

## 功能概述

fi.mcp 现已支持通过 SOCKS5 代理访问外部 API。所有 HTTP 请求都会根据配置自动使用代理。

## 配置方法

在 `etc/fi.mcp.yaml` 配置文件中添加以下配置：

```yaml
Proxy:
  Enable: true # 是否启用代理
  Socks5: "100.115.184.183:1080" # SOCKS5 代理地址
```

### 配置参数说明

- **Enable**: 布尔值，控制是否启用代理
  - `true`: 启用代理，所有 HTTP 请求将通过 SOCKS5 代理
  - `false`: 禁用代理，直接连接

- **Socks5**: 字符串，SOCKS5 代理服务器地址
  - 格式: `host:port`
  - 示例: `"127.0.0.1:1080"` 或 `"proxy.example.com:1080"`

## 使用示例

### 启用代理

```yaml
Proxy:
  Enable: true
  Socks5: "127.0.0.1:1080"
```

### 禁用代理

```yaml
Proxy:
  Enable: false
  Socks5: ""
```

## 技术实现

### 核心组件

1. **代理工具** (`internal/logic/snowball/utils.proxy.go`)
   - 提供统一的 HTTP 客户端创建接口
   - 自动根据配置决定是否使用代理
   - 支持 SOCKS5 代理协议

2. **配置接口** (`internal/config/config.go`)
   - `IsProxyEnabled()`: 检查代理是否启用
   - `GetSocks5Addr()`: 获取代理地址

### 使用方式

在 Snowball 模块内部，创建 HTTP 客户端时使用以下方式：

```go
// 使用配置创建客户端
client := NewClientWithConfig(&config)

// 或者直接传递参数
client := NewClient(proxyEnabled, socks5Addr)
```

## 影响范围

以下模块的所有 HTTP 请求已支持代理：

- **Snowball (雪球)**: 所有股票数据查询
  - 股票建议 (suggest)
  - K线数据 (kline)
  - 财务指标 (indicator)
  - 利润表 (income)
  - 资产负债表 (balance)
  - 现金流量表 (cashflow)
  - Cookie 管理

**注意**: OKX 和 Akshare 模块暂不支持代理，仍使用直连方式。

## 故障处理

如果代理配置失败（例如代理服务器不可达），系统会：

1. 记录错误日志
2. 自动降级为直连模式
3. 继续提供服务（不会因代理问题导致服务不可用）

## 注意事项

1. 确保 SOCKS5 代理服务器可访问
2. 代理地址格式必须正确 (`host:port`)
3. 如果不需要代理，建议设置 `Enable: false` 而不是删除配置
4. 修改配置后需要重启服务才能生效

## 依赖

- `golang.org/x/net/proxy`: SOCKS5 代理支持
- `github.com/go-resty/resty/v2`: HTTP 客户端库
