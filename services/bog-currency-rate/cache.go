package bogcurrencyrate

import (
	"sync"
	"time"
)

var (
	cache = make(map[string]CacheItem)
	mutex sync.Mutex
)

const cacheDuration = 10 * time.Minute

type CacheItem struct {
	RateInfo  RateResponse
	Timestamp time.Time
}
