package setting

type Config struct {
	Server ServerSetitng `mapstructure:"server"`
	Logger LoggerSetting `mapstructure:"logger"`
	MySql  MySQLSetting  `mapstructure:"mysql"`
	Redis  RedisSetting  `mapstructure:"redis"`
}

type ServerSetitng struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type MySQLSetting struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	ConnMaxLife  int    `mapstructure:"conn_max_life"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}
