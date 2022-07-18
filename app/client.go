/**
* @Author: scjtqs
* @Date: 2022/7/18 11:54
* @Email: scjtqs@qq.com
 */
package app

import (
	"github.com/scjtqs2/fqsign_go/app/adapter"
	"github.com/scjtqs2/fqsign_go/config"
	"github.com/scjtqs2/fqsign_go/instance"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var cfg *config.Option

// RunClient 一次性脚本运行
func RunClient(ct *dig.Container) {
	err := ct.Invoke(func(c *config.Option) {
		cfg = c
	})
	if err != nil {
		log.Fatalf("解析配置信息失败 err=%v", err)
	}
	for s, options := range cfg.Users {
		var app instance.AppInterface
		for _, option := range options {
			switch s {
			case config.ConfigNameVpork: // 初始化 vpork
				app = adapter.NewVpork(option, cfg.QqPush)
			case config.ConfigNameVporkVip: // 初始化 vpork_vip
				app = adapter.NewVporkVip(option, cfg.QqPush)
			}
			err = app.Login()
			if err != nil {
				log.Println(err)
				continue
			}
			err = app.Checkin()
			if err != nil {
				log.Println(err)
				continue
			}
			err = app.Logout()
			if err != nil {
				continue
			}
		}

	}
}
