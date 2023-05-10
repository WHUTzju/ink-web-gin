package config

import "ink-web/src/util/log"

// 日志配置
type LogsConfigurationS struct {
	Category                 string                    `mapstructure:"category" json:"category"`
	Level                    log.Level                 `mapstructure:"level" json:"level"`
	Json                     bool                      `mapstructure:"json" json:"json"`
	LineNum                  LogsLineNumConfigurationS `mapstructure:"line-num" json:"lineNum"`
	OperationKey             string                    `mapstructure:"operation-key" json:"operationKey"`
	OperationAllowedToDelete bool                      `mapstructure:"operation-allowed-to-delete" json:"operationAllowedToDelete"`
}
type LogsLineNumConfigurationS struct {
	Disable bool `mapstructure:"disable" json:"disable"`
	Level   int  `mapstructure:"level" json:"level"`
	Version bool `mapstructure:"version" json:"version"`
	Source  bool `mapstructure:"source" json:"source"`
}
