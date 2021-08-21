package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/Aldiwildan77/backend-skeleton/libs/files"
	"github.com/Aldiwildan77/backend-skeleton/libs/logger"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// Cfg config variable
var Cfg Config

func Initialize() {
	loadEnvVars()
	loadDatabases()
	loadLogger()
	loadMessageBroker()
}

func loadEnvVars() {
	os.Setenv("TZ", "Asia/Jakarta")

	cdir, _ := os.Getwd()
	ef := fmt.Sprintf("%s/%s", cdir, ".env")

	var err error
	if files.FileExists(ef) {
		err = loadConfigFile(ef)
	} else {
		err = loadConfigEnvVar()
	}

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err.Error()))
	}
}

func loadDatabases() {
	DBs = Databases{
		DB: DBInstance("name"),
	}
	RedisClient = RedisInstance()
}

func loadLogger() {
	logger.InitLogger(Cfg.Logger.MaxAge, Cfg.Logger.Debug)
}

func loadConfigFile(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Cfg)
	return
}

func loadConfigEnvVar() (err error) {
	viper.AutomaticEnv()

	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &Cfg,
	}

	n := make(map[string]interface{})
	for _, e := range os.Environ() {
		data := strings.Split(e, "=")
		n[data[0]] = data[1]
	}

	decoder, _ := mapstructure.NewDecoder(config)
	decoder.Decode(n)

	return
}

func loadMessageBroker() {
	KafkaConsumer = KafkaConsumerInstance()
	KafkaProducer = KafkaProducerInstance()
}
