package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/rclone/rclone/myali/utils"
)

var (
	conf = Config{}
)

type AliWebDavConfig struct {
	// flags
	AutoIndex string `json:"auto_index"`
	NoTrash   string `json:"no_trash"`
	ReadOnly  string `json:"read_only"`
	Version   string `json:"version"`

	// option
	AuthPassword     string `json:"auth_password"`
	AuthUser         string `json:"auth_user"`
	CacheSize        string `json:"cache_size"`
	CacheTtl         string `json:"cache_ttl"`
	DomainId         string `json:"domain_id"`
	Host             string `json:"host"`
	Port             string `json:"port"`
	ReadBuffSize     string `json:"read_buff_size"`
	SyncRefreshToken string `json:"sync_refresh_token"`
	RefreshToken     string `json:"refresh_token"`
	Root             string `json:"root"`
	WorkDir          string `json:"omitempty"`
}

type RcloneMountConfig struct {
	Remote      string `json:"remote"`
	Mountpoint  string `json:"mountpoint"`
	CacheDir    string `json:"cache_dir"`
	CacheMaxAge int    `json:"cache_max_age"`
}

type Config struct {
	Webdav AliWebDavConfig   `json:"aliwebdav"`
	Mount  RcloneMountConfig `json:"mount"`
}

func LoadConfig() {
	b, _ := ioutil.ReadFile(utils.ConfigPath())
	err := json.Unmarshal(b, &conf)
	if err != nil {
		log.Fatalln("加载配置文件失败", err)
	}
}

func GetConfig() *Config {
	return &conf
}
