package common

import "github.com/asim/go-micro/v3/config"

type MysqlConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	DB   string `json:"db"`
}

func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	_ = config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
