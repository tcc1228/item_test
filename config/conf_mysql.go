package config

type Mysql struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Db        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Log_level string `yaml:"log_level"` //日志等级，debug就是输出全部sql，dev，release
}
