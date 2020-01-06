package Database

import (
	"fmt"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Settings"
	"github.com/gomodule/redigo/redis"
	"time"
)

type CacheDB struct {
	IDatabase
	clientPool *redis.Pool
}

func ConnectWithCacheDB(config *Settings.CacheDbConf) (*CacheDB, error) {

	Logger.SysLog.Info("[CacheDatabase] Connecting to Cache Service")

	client := &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Millisecond,
		Wait:        config.Wait,
		Dial: func() (redis.Conn, error) {
			Logger.SysLog.Debug("[CacheDatabase] Dial Connects To The Cache Server")
			connectString := fmt.Sprintf("%s:%d", config.Host, config.Port)
			c, err := redis.Dial("tcp", connectString)
			if err != nil {
				return nil, err
			}
			if config.Password != "" {
				if _, err := c.Do("AUTH", config.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	CacheDbClient := &CacheDB{
		clientPool: client,
	}

	// Checking cache can be connected
	err := CacheDbClient.pingHealth()
	if err != nil {
		return nil, err
	}

	return CacheDbClient, nil
}

func (r *CacheDB) pingHealth() error {
	conn := r.GetClient()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		return err
	}
	Logger.SysLog.Info("[CacheDatabase] Connected to the Redis Database")
	return nil
}

func (r *CacheDB) GetClient() redis.Conn {
	return r.clientPool.Get()
}

func (r *CacheDB) SetString(key string, data string, time int) error {
	conn := r.clientPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, data)
	if err != nil {
		return err
	}

	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *CacheDB) Exists(key string) bool {
	conn := r.clientPool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func (r *CacheDB) Get(key string) ([]byte, error) {
	conn := r.clientPool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (r *CacheDB) Delete(key string) (bool, error) {
	conn := r.clientPool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func (r *CacheDB) LikeDeletes(key string) error {
	conn := r.clientPool.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = r.Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
