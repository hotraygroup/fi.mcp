package logic

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/sync/singleflight"
)

type Cookie struct {
	cookies []*http.Cookie
	stamp   time.Time
}

var globalCookie = &Cookie{}

var globalSF = new(singleflight.Group)

func (c *Cookie) Get() []*http.Cookie {
	return c.cookies
}

func (c *Cookie) Set(cookies []*http.Cookie) {
	c.cookies = cookies
	c.stamp = time.Now()
}

func (c *Cookie) Expired() bool {
	return c.cookies == nil || c.stamp.Add(time.Duration(60+rand.Intn(120))*time.Second).Before(time.Now())
}

func setHeader(userAgent, indexURL string, _client *resty.Client) {

	if globalCookie.Expired() {
		_, err, _ := globalSF.Do("getCookie", func() (interface{}, error) {
			req := resty.New().R()
			resp, err := req.SetHeader("User-Agent", userAgent).Get(indexURL)
			if err != nil {
				return nil, err
			} else {
				globalCookie.Set(resp.Cookies())
			}
			return nil, nil
		})

		if err != nil {
			logx.Errorf("get cookie err is %s", err)
		}
	}

	cookies := globalCookie.Get()

	if cookies == nil {
		panic("cookie is null")
	}

	_client.SetCookies(cookies)
	_client.SetHeader("User-Agent", userAgent)
	_client.SetHeader("Referer", indexURL)

}
