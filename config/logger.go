package config

type logger struct {
	Level  string `yaml:"level"`
	Type   string `json:"type"`
	Path   string `yaml:"path"`
	MaxDay int    `yaml:"max-day"`
}
