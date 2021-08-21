package config

// DBConfig store  db config from env
type DBConfig struct {
	Engine            string `mapstructure:"DATABASE_ENGINE"`
	Host              string `mapstructure:"DATABASE_HOST"`
	Port              int    `mapstructure:"DATABASE_PORT"`
	Username          string `mapstructure:"DATABASE_USERNAME"`
	Password          string `mapstructure:"DATABASE_PASSWORD"`
	Schema            string `mapstructure:"DATABASE_SCHEMA"`
	ReconnectRetry    int    `mapstructure:"DATABASE_RECONNECT_RETRY"`
	ReconnectInterval int64  `mapstructure:"DATABASE_RECONNECT_INTERVAL"`
	DebugMode         bool   `mapstructure:"DATABASE_DEBUG_MODE"`
	PoolMaxOpen       int    `mapstructure:"DATABASE_POOL_MAX_OPEN_CONNS"`
	PoolMaxIdle       int    `mapstructure:"DATABASE_POOL_MAX_IDLE_CONNS"`
}

// RedisStorageConfig store redis config from env
type RedisStorageConfig struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     int    `mapstructure:"REDIS_PORT"`
	DB       int    `mapstructure:"REDIS_DB"`
	Username string `mapstructure:"REDIS_AUTH_USERNAME"`
	Password string `mapstructure:"REDIS_AUTH_PASSWORD"`
	Debug    bool   `mapstructure:"REDIS_DEBUG_MODE"`
}

// LogConfig store log config
type LogConfig struct {
	MaxAge int  `mapstructure:"LOG_MAX_AGE"`
	Debug  bool `mapstructure:"LOG_DEBUG"`
}

type KafkaConfig struct {
	Host         string `mapstructure:"KAFKA_HOST"`
	Port         int    `mapstructure:"KAFKA_PORT"`
	GroupID      string `mapstructure:"KAFKA_GROUP_ID"`
	MultipleHost string `mapstructure:"KAFKA_MULTIPLE_HOST"`
}

// Config store all configuration from env
type Config struct {
	DB        DBConfig           `mapstructure:",squash"`
	Redis     RedisStorageConfig `mapstructure:",squash"`
	Port      int                `mapstructure:"PORT"`
	Logger    LogConfig          `mapstructure:",squash"`
	Kafka     KafkaConfig        `mapstructure:",squash"`
	JWTSecret string             `mapstructure:"JWT_SECRET"`
}
