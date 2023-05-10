package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
)

type Setting struct {
	vp  *viper.Viper
	cfg string
}

// 定义全局变量
var (
	ServerSetting     *ServerSettingS
	DatabaseSetting   *DatabaseSettingS
	LogsConfiguration *LogsConfigurationS
	RuntimeRoot       string
)

// map存储配置
var sections = make(map[string]interface{})

// 读取系统配置
// viper读取配置文件conf/config.yaml
func newSetting(cfg string) (*Setting, error) {

	vp := viper.New()
	if cfg != "" {
		vp.SetConfigFile(cfg)
		fmt.Println("load config cfg:", cfg)
	} else {
		vp.AddConfigPath("conf")
		vp.SetConfigName("config")
		fmt.Println("load defult config cfg:", "conf/config.yaml")
	}
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil { // viper read config
		return nil, err
	}
	s := &Setting{
		vp:  vp,
		cfg: cfg}
	return s, nil
}

// 读取指定段
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

// 重新加载配置
func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// 读取配置到全局变量
func Config(c context.Context, cfg string) error {
	s, err := newSetting(cfg)
	if err != nil {
		return err
	}

	//数据库
	err = s.ReadSection("db", &DatabaseSetting)
	if err != nil {
		return err
	}

	//系统
	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}

	//日志
	err = s.ReadSection("logs", &LogsConfiguration)
	if err != nil {
		return err
	}

	fmt.Println("setting:")
	fmt.Println(ServerSetting)
	fmt.Println(DatabaseSetting)
	fmt.Println(LogsConfiguration)
	return nil
}
