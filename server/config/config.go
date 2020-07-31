package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//envconfig option for .env vars
// os exit codes 1,2,3

// Config struct
type Config struct {
	Security struct {
		Oauth2 struct {
			Github struct {
				ClientID     string `yaml:"clientId"`
				ClientSecret string `yaml:"clientSecret"`
				RedirectURL  string `yaml:"redirectURL"`
			}
		}
	}
	Server struct {
		Host string `yaml:"host"`
		Port int16  `yaml:"port"`
	}
}

//ReadYAMLFile func reads local yaml config and returns a Config struct
func ReadYAMLFile() (*Config, error) {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, fmt.Errorf("yamlFile file load failed #%v ", err)

	}
	c := &Config{}

	if err := yaml.UnmarshalStrict(yamlFile, c); err != nil {
		return nil, fmt.Errorf("Unmarshal: %v", err)

	}
	return c, nil

}
