package sharding

import (
	"crypto/md5"
)

type CalculatorMD5 struct{}

// CalculateShardIndex - used for calculate index for
func (i CalculatorMD5) CalculateShardIndex(key string, shardCount int) int {
	md5hash := md5.Sum([]byte(key))

	var sum int
	for _, value := range md5hash {
		sum = sum + int(value)
	}

	return sum % shardCount
}
