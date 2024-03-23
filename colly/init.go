package colly

import (
	"net/http"
	"time"
)

type BaseInfo struct {
	Cookie    string
	Url       string
	userAgent string
}

func Init(cookie, url, userAgent string) *BaseInfo {
	return &BaseInfo{
		Cookie:    cookie,
		Url:       url,
		userAgent: userAgent,
	}
}

func (l *BaseInfo) InitCookie() *http.Cookie {
	return &http.Cookie{
		Name:    "cookie",
		Value:   l.Cookie,
		Domain:  l.Url,                          // 需要设置合适的域名
		Path:    "/",                            // 设置路径，默认为"/"
		Expires: time.Now().Add(24 * time.Hour), // 设置过期时间
	}
}
