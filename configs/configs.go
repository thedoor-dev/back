package configs

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name           string `mapstructure:"name"`
	Mode           string `mapstructure:"mode"`
	Version        string `mapstructure:"version"`
	Port           int    `mapstructure:"port"`
	*AdminConfig   `mapstructure:"admin"`
	*LogConfig     `mapstructure:"log"`
	*MariaDBConfig `mapstructure:"mysql"`
	*RedisConfig   `mapstructure:"redis"`
}

type AdminConfig struct {
	AdminName   string `mapstructure:"admin_name"`
	AdminPasswd string `mapstructure:"admin_passwd"`
	Img         string `mapstructure:"img"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MariaDBConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// Init 初始化配置
// 读取文件
// 反序列化配置文件
//
//
func Init(path string) (err error) {
	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("配置文件读取失败viper.ReadInConfig failed, err:%v\n", err)
		return
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		log.Printf("配置文件反序列化失败viper.Unmarshal failed, err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置文件被修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			log.Printf("配置文件反序列化失败viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
