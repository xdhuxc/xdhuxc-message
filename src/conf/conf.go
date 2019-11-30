package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/xdhuxc/xdhuxc-message/src/model"
)

var conf *model.Configuration

func InitConfiguration(path string) (*model.Configuration, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func GetConfiguration() *model.Configuration {
	return conf
}
