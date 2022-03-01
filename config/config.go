package config

import (
	"github.com/spf13/viper"
	"time"
)

var (
	ServerConfig *serverConfig
	AppConfig    *appConfig
	RunMode      string
)

type Env struct {
	RunMode string
}

type serverConfig struct {
	RunMode       string
	HttpPort      string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	DBType        string
	UserName      string
	PassWord      string
	Host          string
	DBName        string
	TablePrefix   string
	Charset       string
	ParseTime     bool
	MaxIdleConns  int
	MaxOpenConns  int
	SeverHost     string
	RedisHost     string
	RedisPassword string
	DbUri         string
}

type appConfig struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

func init() {
	PickConfigFile()
	loadConfig()
}

func PickConfigFile() {
	vp := viper.New()
	vp.SetConfigName("runmode")
	vp.AddConfigPath("conf/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	Mode := new(Env)
	if err != nil {
		Mode.RunMode = "config"
	}
	err = vp.UnmarshalKey("Env", Mode)
	RunMode = Mode.RunMode
}

func loadConfig() {
	vp := viper.New()
	vp.SetConfigName(RunMode)
	vp.AddConfigPath("conf/")
	vp.SetConfigType("yaml")
	//fmt.Println(vp.AllKeys())
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	ServerConfig = new(serverConfig)
	AppConfig = new(appConfig)
	err = vp.UnmarshalKey("Server", ServerConfig)
	if err != nil {
		panic(err)
	}
	err = vp.UnmarshalKey("App", AppConfig)
	if err != nil {
		panic(err)
	}
}
