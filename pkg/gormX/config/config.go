package config

import (
	"encoding/json"
	"os"
)

type Database struct { // 需要定义json tag, 用于config.json和client.go之间的数据转换
	Type        string `json:"type"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}

var DatabaseSetting = &Database{}

type Config struct { // 这里可以看config.json的格式
	Database *Database `json:"database"`
}

var GlobalConfigSetting = &Config{}

func init() { // 这个config是需要在import package的时候就初始化
	file, err := os.Open("/home/netlab/xlc_interview/project/test_go_basic/gormX/config/config.json")
	if err != nil {
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(GlobalConfigSetting)
	if err != nil {
		return
	}
	DatabaseSetting = GlobalConfigSetting.Database
}
