package rateLimiter

import (
	"rateLimiter/bucket"
)

var clientBucketMap = make(map[string]*bucket.TokenBucket)

type Rule struct {
	MaxTokens int64
	Rate      int64
}

func GetBucket(identifier string, usertype string) *bucket.TokenBucket {
	//TODO: Not thread safe
	if clientBucketMap[identifier] == nil {
		clientBucketMap[identifier] = bucket.NewTokenBucket(rulesMap[usertype].MaxTokens, rulesMap[usertype].Rate)

	}
	return clientBucketMap[identifier]
}

/*
* These rules map can be fetched from a database rather than
* hardcoding at application layer and also can be periodically
* updated in a background job in case some dynamic changes are needed
 */
var rulesMap = map[string]Rule{
	"gen-user": {MaxTokens: 1, Rate: 5},
}
