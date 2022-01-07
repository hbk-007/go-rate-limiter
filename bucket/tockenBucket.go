package bucket

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	rate               int64
	maxTokens          int64
	currentTokens      int64
	lastRefillTmestamp time.Time
	mutex              sync.Mutex
}

func NewTokenBucket(Rate int64, MaxTokens int64) *TokenBucket {
	return &TokenBucket{
		rate:               Rate,
		maxTokens:          MaxTokens,
		lastRefillTmestamp: time.Now(),
		currentTokens:      MaxTokens,
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	end := time.Since(tb.lastRefillTmestamp)
	tokensTobeAdded := (end.Nanoseconds() * tb.rate) / 1000000000
	tb.currentTokens = int64(math.Min(float64(tb.currentTokens+tokensTobeAdded), float64(tb.maxTokens)))
	tb.lastRefillTmestamp = now
}

func (tb *TokenBucket) IsRequestAllowed(tokens int64) bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	tb.refill()
	if tb.currentTokens >= tokens {
		tb.currentTokens = tb.currentTokens - tokens
		return true
	}
	return false
}
