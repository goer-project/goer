package config

type Dir struct {
	Cmd        string `mapstructure:"cmd" json:"cmd" yaml:"cmd"`
	Controller string `mapstructure:"controller" json:"controller" yaml:"controller"`
	Middleware string `mapstructure:"middleware" json:"middleware" yaml:"middleware"`
	Migration  string `mapstructure:"migration" json:"migration" yaml:"migration"`
	Model      string `mapstructure:"model" json:"model" yaml:"model"`
	Request    string `mapstructure:"request" json:"request" yaml:"request"`
	Upload     string `mapstructure:"upload" json:"upload" yaml:"upload"`
}

var NewDir = Dir{
	Cmd:        "cmd",
	Controller: "app/http/controllers",
	Middleware: "app/http/middleware",
	Migration:  "database/migrations",
	Model:      "app/models",
	Request:    "app/http/requests",
	Upload:     "storage/public",
}
