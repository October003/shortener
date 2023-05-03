package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	ShortUrlDB struct {
		DSN string
	}
	Sequence struct {
		DSN string
	}
	BaseString string // base62指定基础字符串

	ShortUrlBlackList []string

	ShortDomain string

	CacheRedis cache.CacheConf
}

// type ShortUrlDB struct{
// 	DSN string
// }
