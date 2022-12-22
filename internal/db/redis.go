package db

import (
	"fmt"
	"strconv"
	"surge/internal/config"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type RedisBackend struct {
	conn *redis.Client
}

var windowSizeSeconds = time.Duration(time.Minute * 10).Seconds() // in seconds
var rb *RedisBackend
var onceRedis sync.Once

func GetRedisBackend() *RedisBackend {
	onceRedis.Do(func() {
		cfg := config.GetConfig()
		rb = &RedisBackend{redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
			Password: cfg.RedisPassword,
			DB:       0, // use default DB
		})}
		_, err := rb.conn.Ping().Result()
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Infoln("Connection established to Redis")
	})

	return rb
}

func (rb *RedisBackend) AddNewRequestWithPrefix(prefix string) int64 {
	now := time.Now().Unix()
	// get current window count
	currentWindow := strconv.FormatInt(now/int64(windowSizeSeconds), 10)
	key := prefix + ":" + currentWindow
	value, _ := rb.conn.Get(key).Result()
	requestCountCurrentWindow, _ := strconv.ParseInt(value, 10, 64)

	// get last window reminder
	lastWindow := strconv.FormatInt(((now - int64(windowSizeSeconds)) / int64(windowSizeSeconds)), 10)
	keyLast := prefix + ":" + lastWindow
	value, _ = rb.conn.Get(keyLast).Result()
	requestCountlastWindow, _ := strconv.ParseInt(value, 10, 64)

	elapsedTimePercentage := (now % int64(windowSizeSeconds) * 100) / int64(windowSizeSeconds)
	lastWindowReminder := requestCountlastWindow * elapsedTimePercentage / 100

	totalRequests := lastWindowReminder + requestCountCurrentWindow

	rb.conn.Incr(key)

	return totalRequests
}
