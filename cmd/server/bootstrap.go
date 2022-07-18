/**
* @Author: scjtqs
* @Date: 2022/7/18 11:26
* @Email: scjtqs@qq.com
 */
package main

import (
	"github.com/scjtqs2/fqsign_go/config"
	"go.uber.org/dig"
	"gopkg.in/yaml.v2"
)

func bootstrap() (*dig.Container, error) {
	container := dig.New()
	err := container.Provide(func() (*config.Option, error) {
		var cfg *config.Option
		configData, err := config.LoadConf(c)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(configData, &cfg)
		return cfg, err
	})

	return container, err
}
