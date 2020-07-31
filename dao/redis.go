package dao

import (
	"blogweb_gin/config"
	logger "blogweb_gin/gb"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var (
	Client *redis.Client
)

const (
	KeyArticleCountPerDay = "article:count:%s"
)

// 初始化连接
func InitRedis(cfg *config.RedisConfig) (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	_, err = Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func AddArticleCount(articleId int) (err error) {
	data := time.Now().Format("20160102")
	redisKey := fmt.Sprintf(KeyArticleCountPerDay, data)
	err = Client.ZIncrBy(redisKey, 1, fmt.Sprintf("%d", articleId)).Err()
	return
}
func ArticleTopN(n int64) ([]int64, error) {
	data := time.Now().Format("20160102")
	redisKey := fmt.Sprintf(KeyArticleCountPerDay, data)
	idStrs, err := Client.ZRevRange(redisKey, 0, n-1).Result()
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(idStrs))
	for _, value := range idStrs {
		id, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			logger.Warn("ArticleTopN:strconv.ParseInt failed", zap.Any("error", err))
			continue
		}
		ids = append(ids, id)
	}
	return ids, nil
}
