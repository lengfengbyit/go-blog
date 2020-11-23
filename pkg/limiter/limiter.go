package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

// Limiter 限流器
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// LimiterBucketRule 限流器令牌桶规则
type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64 // 每次到达间隔后所放的具体令牌数量
}



