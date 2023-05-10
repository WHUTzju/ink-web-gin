package config

// 数据库配置
type DatabaseSettingS struct {
	Uri         string `mapstructure:"uri" json:"uri"`
	TablePrefix string `mapstructure:"table_prefix" json:"tablePrefix"`
}
