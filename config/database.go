package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig represents config for connecting to Database
type DatabaseConfig struct {
	Engine            string
	Host              string
	User              string
	Password          string
	Schema            string
	Port              int
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
}

// RedisConfig represents config for connecting to Redis
type RedisConfig struct {
	Addr     string
	Username string
	Password string
	DB       int
	Debug    bool
}

// Databases store more db connection
type Databases struct {
	DB *gorm.DB
}

var (
	// DBs represent Databases struct
	DBs Databases

	// RedisClient represent redis client
	RedisClient *redis.Client

	ctx = context.Background()
)

func loadRedisConfiguration() RedisConfig {
	addr := fmt.Sprintf("%s:%d", Cfg.Redis.Host, Cfg.Redis.Port)

	conf := RedisConfig{
		Addr:     addr,
		Username: Cfg.Redis.Username,
		Password: Cfg.Redis.Password,
		DB:       Cfg.Redis.DB,
		Debug:    Cfg.Redis.Debug,
	}

	return conf
}

// RedisInstance to instantiation a redis client
func RedisInstance() *redis.Client {
	conf := loadRedisConfiguration()
	opt := &redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
		Username: conf.Username,
	}

	if !conf.Debug {
		opt.TLSConfig = &tls.Config{}
	}

	redis := redis.NewClient(opt)
	if _, err := redis.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return redis
}

// RedisPing to ping redis server
func RedisPing() error {
	redis := RedisInstance()
	if _, err := redis.Ping(ctx).Result(); err != nil {
		redis.Close()
		RedisInstance()
	}

	return nil
}

func loadDBConfiguration() DatabaseConfig {
	conf := DatabaseConfig{
		Engine:            Cfg.DB.Engine,
		Host:              Cfg.DB.Host,
		User:              Cfg.DB.Username,
		Password:          Cfg.DB.Password,
		Schema:            Cfg.DB.Schema,
		Port:              Cfg.DB.Port,
		ReconnectRetry:    Cfg.DB.ReconnectRetry,
		ReconnectInterval: Cfg.DB.ReconnectInterval,
		DebugMode:         Cfg.DB.DebugMode,
	}

	return conf
}

func loadDBLoggerConfiguration() logger.Interface {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	return logger
}

// DBInstance to instantiation a database client
func DBInstance(connectionType string) *gorm.DB {
	var conf DatabaseConfig

	switch connectionType {
	case "":
		conf = loadDBConfiguration()
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Schema,
	)

	gormConf := &gorm.Config{}

	if conf.DebugMode {
		gormConf.Logger = loadDBLoggerConfiguration()
	}

	instance, err := gorm.Open(mysql.Open(dsn), gormConf)
	if err != nil {
		panic(err)
	}

	db, _ := instance.DB()

	switch connectionType {
	case "":
		db.SetMaxIdleConns(Cfg.DB.PoolMaxIdle)
		db.SetMaxOpenConns(Cfg.DB.PoolMaxOpen)
	}

	if conf.DebugMode {
		return instance.Debug()
	}

	return instance
}
