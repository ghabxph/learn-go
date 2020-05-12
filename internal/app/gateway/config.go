package gateway

type RealIpHeader struct {
	Enabled bool   `yaml:"enabled"`
	Value   string `yaml:"value"`
}

type Routes struct {
	From     string `yaml:"from"`
	To       string `yaml:"to"`
	Endpoint string `yaml:"endpoint"`
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
}

type Config struct {
	RealIpHeader RealIpHeader `yaml:"realIpHeader"`
	Routes       []Routes     `yaml:"routes"`
}

/** Config Instance **/
var __config_instance Config

func (t *Config) Configure(c *Config) {
	t = c
}

func (t *Config) LoadFile(path string) *Config {

	return t
}
