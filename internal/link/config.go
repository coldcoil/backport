package link

type Config struct {
	Targets []target `yaml:"targets"`
}

type target struct {
	Network string            `yaml:"network"`
	Addr    string            `yaml:"addr"`
	Labels  map[string]string `yaml:"labels"`
	Via     []via             `yaml:"via"`
}
