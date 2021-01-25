package filters

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"app/config"
	"time"
)

func RegisterSession() gin.HandlerFunc {
	store, _ := redis.NewStore(
		10,
		"tcp",
		config.GetEnv().RedisIp+":"+config.GetEnv().RedisPort,
		config.GetEnv().RedisPassword,
		[]byte(config.GetEnv().SessionSecret))
	return sessions.Sessions(config.GetEnv().SessionKey, store)
}

func RegisterCache() gin.HandlerFunc {
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(config.GetEnv().RedisIp+":"+config.GetEnv().RedisPort, config.GetEnv().RedisPassword, time.Minute)
	return cache.Cache(&cacheStore)
}


