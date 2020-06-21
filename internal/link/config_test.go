package link

import (
	"io/ioutil"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	in, err := ioutil.ReadFile("./fixtures/anchors.yaml")
	require.Nil(t, err)
	var c Config
	err = yaml.Unmarshal(in, &c)
	require.Nil(t, err)
	assert.Equal(t, Config{
		Targets: []target{
			{
				Network: "tcp",
				Addr:    "localhost:22",
				Labels:  map[string]string{"service": "ssh"},
				Via: []via{
					{
						URL: "https://a.example.com/",
						Authentication: authentication{
							ClientCertificate: clientCertificate{
								Path: "/etc/identity.pem",
							},
						},
					},
					{
						URL: "https://b.example.com",
						Authentication: authentication{
							ClientCertificate: clientCertificate{
								Path: "/etc/identity.pem",
							},
						}},
				},
			},
		},
	}, c)
}
