package lru

import (
	"time"

	"github.com/bluele/gcache"
	"github.com/karlseguin/ccache"
)

type RealPhone string

func (p *RealPhone) Phone() string {
	return string(*p)
}

type VirtualPhone string

func (p *VirtualPhone) Phone() string {
	return string(*p)
}

type Interface interface {
	GetCalltracking(phones []RealPhone) (phonesInCache map[RealPhone]VirtualPhone, phonesNotFoundInCache []RealPhone)
	SetCalltracking(realPhone RealPhone, virtualPhone VirtualPhone) error

	GetPhone(phoneIDs []int64) (phonesInCache map[int64]RealPhone, phonesNotFoundInCache []int64)
	SetPhone(phoneID int64, realPhone RealPhone) error

	Stat() float64
	Len() int
}

type cacheConfigRow struct {
	maxRowsCount int
	ttl          time.Duration
	enabled      bool
}

type cacheConfig struct {
	calltracking cacheConfigRow
	phones       cacheConfigRow
}

var config = cacheConfig{
	calltracking: cacheConfigRow{
		maxRowsCount: 10000,
		ttl:          time.Minute * 10,
		enabled:      true,
	},
	phones: cacheConfigRow{
		maxRowsCount: 5000000,
		ttl:          time.Minute,
		enabled:      true,
	},
}

// In-memory кэш
type Cache struct {
	calltracking gcache.Cache
	phones       gcache.Cache
}

// Возвращает новый in-memory кэш
func NewLruCache() *Cache {
	return &Cache{
		calltracking: ccache.Configure().Buckets(8).MaxSize(5000000),
		phones:       gcache.New(config.phones.maxRowsCount).LRU().Build(),
	}
}
