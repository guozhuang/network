package config

import (
	"github.com/gin-gonic/gin" //新版本golang不能引入带@的包
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	RunMode  string
	PageSize int
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	//先根据整体应用进行模式加载，然后区分相应的不同环境再加载对应的文件
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("app", AppSetting)

	if AppSetting.RunMode == gin.DebugMode {
		cfg, err = ini.Load("config/dev/dev.ini")
		if err != nil {
			log.Fatalf("setting.Setup, fail to parse 'conf/dev/dev.ini': %v", err)
		}
	} else {
		cfg, err = ini.Load("config/production/production.ini")
		if err != nil {
			log.Fatalf("setting.Setup, fail to parse 'conf/production/production.ini': %v", err)
		}
	}

	mapTo("server", ServerSetting)
	mapTo("redis", RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
