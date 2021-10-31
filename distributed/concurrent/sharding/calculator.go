package sharding

// Calculator - a common contract for finding the shard index
type Calculator interface {
	CalculateShardIndex(key string, shardsCount int) int
}
