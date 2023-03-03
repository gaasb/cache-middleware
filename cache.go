package cache_middleware

import (
	"flag"
	"sync"
	"time"
)

const (
	DEFAULT_TTL            = 3600
	DEFAULT_FLUSH_ON_START = true
)

type CacheManager struct {
	sync.RWMutex
	storage *Storage
	options *CacheOpt
}
type CacheOpt struct {
	FlushOnStart bool
	TTL          time.Duration
}

func NewCache(storage *StorageProvider, options *CacheOpt) CacheManager {
	flag.String("addr", "localhost:9001", "address for distribute cache.")
	cache := CacheManager{}
	cache.Lock()
	cache.Unlock()
	if options != nil {
		cache.options = options
	} else {
		cache.options = &CacheOpt{
			FlushOnStart: DEFAULT_FLUSH_ON_START,
			TTL:          DEFAULT_TTL,
		}
	}
	if storage != nil {
		cache.storage.Provider = storage
	} else {
		cache.storage = &Storage{
			Provider: storage, //TODO add sync.Map type
			TTL:      cache.options.TTL,
		}
	}
	return cache
}
