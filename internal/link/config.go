package link

// https://medium.com/@kinghuang/docker-compose-anchors-aliases-extensions-a1e4105d70bd
// https://github.com/goccy/go-yaml

type Config struct {
	Targets []target `yaml:"targets"`
}

type target struct {
	Network string            `yaml:"network"`
	Addr    string            `yaml:"addr"`
	Labels  map[string]string `yaml:"labels"`
	Via     []via             `yaml:"via"`
}

type via struct {
	URL            string         `yaml:"url"`
	Authentication authentication `yaml:"authentication"`
}

type authentication struct {
	ClientCertificate clientCertificate `yaml:"client_certificate"`
}

type clientCertificate struct {
	Path string `yaml:"path"`
}
