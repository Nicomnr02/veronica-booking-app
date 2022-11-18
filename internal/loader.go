package internal

import (
	"fmt"
	"log"
	"path/filepath"
	"pos/internal/config/db"
	"pos/internal/config/logging"
	"pos/internal/config/server"
	"runtime"

	"github.com/spf13/viper"
)

// Config is struct all of configuration
type Config struct {
	Server   server.ServerList
	Database db.DatabaseList
	Client   server.ClientList
	Logger   logging.LoggerConfig
}

var configuration Config

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b) + "/config"
)

func init() {
	var err error
	fmt.Printf("Reading Config %s\n", basepath+"/server")
	viper.AddConfigPath(basepath + "/server")
	viper.SetConfigType("yaml")
	viper.SetConfigName("server.yml")
	errConf := viper.ReadInConfig()
	if errConf != nil {
		log.Fatalf("Failed to load config: %v", errConf)
	}

	viper.AddConfigPath(basepath + "/server")
	viper.SetConfigName("client.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load server client config: %v", err))
	}

	fmt.Printf("Reading Config %s\n", basepath+"/db")
	viper.AddConfigPath(basepath + "/db")
	viper.SetConfigName("database.yml")
	err = viper.MergeInConfig()
	if err != nil {
		log.Println("Cannot read database config: %v", err)
	}

	fmt.Printf("Reading Config %s\n", basepath+"/logging")
	viper.AddConfigPath(basepath + "/logging")
	viper.SetConfigName("logger.yml")
	err = viper.MergeInConfig()
	if err != nil {
		log.Println("Failed to load log config: %v", err)
	}

	viper.Unmarshal(&configuration)
	viper.AutomaticEnv()
}

func GetConfig() *Config {
	return &configuration
}
