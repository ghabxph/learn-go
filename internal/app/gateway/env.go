package gateway

type Env struct {
	REAL_IP_HEADER_ENABLED bool   `yaml:"REAL_IP_HEADER_ENABLED"`
	REAL_IP_HEADER_VALUE   string `yaml:"REAL_IP_HEADER_VALUE"`
	ROUTES_BLOG_ENDPOINT   string `yaml:"ROUTES_BLOG_ENDPOINT"`
	ROUTES_BLOG_HOSTNAME   string `yaml:"ROUTES_BLOG_HOSTNAME"`
	ROUTES_BLOG_PORT       string `yaml:"ROUTES_BLOG_PORT"`
}
