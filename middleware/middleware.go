package middleware

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"net/http"
	rateLimiter "rateLimiter/rules"

	"github.com/gin-gonic/gin"
)

func RateLimit(c *gin.Context) {
	var userType string
	if val, exists := c.Get("user-type"); exists {
		userType = val.(string)
	}
	if userType == "" {
		userType = "gen-user"
	}
	tokenBucket := rateLimiter.GetBucket(GetClientIdentifier(c), userType)
	if !tokenBucket.IsRequestAllowed(5) {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "Try Again after Sometime!",
		})
		return
	}
	c.Next()
}

func GetClientIdentifier(c *gin.Context) string {
	ip := c.ClientIP()
	url := c.Request.URL.Path
	data := fmt.Sprint("%s-%s", ip, url)
	h := md5.Sum([]byte(data))
	hash := new(big.Int).SetBytes(h[:]).Text(62)
	return hash
}
