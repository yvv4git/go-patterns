package sharding

import (
	"sync"
)

// Shard - some realisation one segment
type Shard struct {
	sync.RWMutex
	bucket          map[string]interface{}
	indexCalculator Calculator
}

// ShardedMap - works with shards
type ShardedMap struct {
	shards          []*Shard
	shardCalculator Calculator
}

// NewSharedMap - simple fabric for create instance of SharedMap
func NewSharedMap(count int, indexCalculator Calculator) *ShardedMap {
	shards := make([]*Shard, count)

	for i := 0; i < count; i++ {
		shards[i] = &Shard{
			bucket: make(map[string]interface{}),
		}
	}

	return &ShardedMap{
		shards:          shards,
		shardCalculator: indexCalculator,
	}
}

// Get - used for get value by key concurrently
func (m ShardedMap) Get(key string) interface{} {
	shardIndex := m.shardCalculator.CalculateShardIndex(key, len(m.shards))
	shard := m.shards[shardIndex]

	shard.RLock()
	defer shard.RUnlock()

	return shard.bucket[key]
}

// Set - used for set value for key
func (m ShardedMap) Set(key string, value interface{}) {
	shardIndex := m.shardCalculator.CalculateShardIndex(key, len(m.shards))
	shard := m.shards[shardIndex]

	shard.Lock()
	defer shard.Unlock()

	shard.bucket[key] = value
}

// Keys - used for get keys slice
func (m ShardedMap) Keys() []string {
	keys := make([]string, 0)
	mutex := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(len(m.shards))

	for _, shard := range m.shards {
		go func(shard *Shard) {
			shard.RLock()

			for key := range shard.bucket {
				mutex.Lock()
				keys = append(keys, key)
				mutex.Unlock()
			}

			shard.RUnlock()
			wg.Done()
		}(shard)
	}

	wg.Wait()
	return keys
}
