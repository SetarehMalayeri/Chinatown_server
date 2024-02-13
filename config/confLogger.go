package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show-line"`      //行号显示？
	LogInConsole bool   `yaml:"log-in-console"` //显示打印路径？
}
