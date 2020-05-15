package m_cache

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

type RedisCache struct {
	Key      string `json:"key"`
	Conn     string `json:"conn"`
	Password string `json:"password"`
	DbNum    int    `json:"dbNum"`
}

type FileCache struct {
	CachePath      string `json:"CachePath"`
	FileSuffix     string `json:"FileSuffix"`
	DirectoryLevel int    `json:"DirectoryLevel"`
	EmbedExpiry    int    `json:"EmbedExpiry"`
}

func GetCache(driver string) (cache.Cache, error) {
	if driver == "redis" {
		return newRedisCache(), nil
	}
	if driver == "file" {
		return newFileCache(), nil
	}
	return nil, fmt.Errorf("invalid cache driver")
}

func newRedisCache() cache.Cache {
	redis := RedisCache{
		Key:      beego.AppConfig.String("rediskey"),
		Conn:     beego.AppConfig.String("redisurls"),
		Password: beego.AppConfig.String("redispass"),
		DbNum:    beego.AppConfig.DefaultInt("redisdb", 0),
	}
	res, _ := json.Marshal(redis)
	cache, err := cache.NewCache("redis", string(res))
	if err != nil {
		log.Fatalf("redis conn error:%v", err)
	}
	return cache
}

func newFileCache() cache.Cache {
	fi := FileCache{
		CachePath:      "./runtime/cache",
		FileSuffix:     ".cache",
		DirectoryLevel: 2,
		EmbedExpiry:    120,
	}
	res, _ := json.Marshal(fi)
	cache, err := cache.NewCache("file", string(res))
	if err != nil {
		log.Fatalf("file cache error:%v", err)
	}
	return cache
}
