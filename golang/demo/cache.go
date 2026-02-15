package main

import (
	"hash"
	"hash/fnv"
	"runtime"
	"sync"
)

type shard struct {
	lock sync.RWMutex
	data map[string]string
}

type shardCache struct {
	shards   []*shard
	pool     *sync.Pool
	shardNum int
}

func (s *shardCache) Set(k, v string) {
	shard := s.shards[s.shardPosition(k)]
	shard.lock.Lock()
	defer shard.lock.Unlock()
	shard.data[k] = v
}

func (s *shardCache) Get(k string) (string, bool) {
	shard := s.shards[s.shardPosition(k)]
	shard.lock.RLock()
	defer shard.lock.RUnlock()
	if v, ok := shard.data[k]; ok {
		return v, ok
	}
	return "", false
}

var singleton struct {
	once sync.Once
	sc   *shardCache
}

func NewCache() *shardCache {

	singleton.once.Do(func() {
		sc := &shardCache{shardNum: runtime.NumCPU()}
		sc.shards = make([]*shard, sc.shardNum)
		for i := 0; i < sc.shardNum; i++ {
			sc.shards[i] = &shard{
				data: make(map[string]string),
			}
		}
		sc.pool = &sync.Pool{
			New: func() any {
				return fnv.New32a()
			},
		}
		singleton.sc = sc
	})
	return singleton.sc
}

func (sc *shardCache) shardPosition(k string) uint32 {
	hash := sc.pool.Get().(hash.Hash32)
	hash.Reset()
	hash.Write([]byte(k))
	ret := hash.Sum32() % uint32(sc.shardNum)
	sc.pool.Put(hash)
	return ret
}
