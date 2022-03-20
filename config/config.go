package config

type Config struct {
	Database Database `mapstructure:"database" json:"database"`
}

var NewConfig Config
