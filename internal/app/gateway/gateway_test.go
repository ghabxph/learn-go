package gateway

type test struct {
	Test []struct {
		Title      string `yaml:"title"`
		Env        Env    `yaml:"env"`
		Config     Config `yaml:"config"`
		ConfigFile string `yaml:"configFile"`
	} `yaml:"test"`
}
