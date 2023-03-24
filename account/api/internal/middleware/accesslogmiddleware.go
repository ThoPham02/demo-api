package middleware

import (
	"crypto/sha256"
	"demo-api/account/redisCli"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type AccessLogMiddleware struct {
	redisCache *redis.Redis
}

func NewAccessLogMiddleware() *AccessLogMiddleware {
	return &AccessLogMiddleware{
		redisCache: redisCli.NewRedis("localhost", "6379"),
	}
}

func (m *AccessLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		accessToken := r.Header.Get("Authentication")
		shortKey := m.GetShortKey(accessToken)
		lastTimeActive := time.Now().Format("2006-01-02 15:04:05")

		m.redisCache.Set(shortKey, lastTimeActive)
		next(w, r)
	}
}

// Gen a short string from access token
func (m *AccessLogMiddleware) GetShortKey(accessToken string) string {
	hash := sha256.Sum256([]byte(accessToken))
	shortKey := hex.EncodeToString(hash[:])[:8]
	return shortKey
}
