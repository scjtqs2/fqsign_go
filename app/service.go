/**
* @Author: scjtqs
* @Date: 2022/7/18 11:51
* @Email: scjtqs@qq.com
 */
package app

import (
	"github.com/robfig/cron/v3"
	"github.com/scjtqs2/fqsign_go/config"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// RunServer 服务方式运行，crontab 方式触发
func RunServer(ct *dig.Container) {
	err := ct.Invoke(func(c *config.Option) {
		cfg = c
	})
	if err != nil {
		log.Fatalf("解析配置信息失败 err=%v", err)
	}
	Cron := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))
	if cfg.Cron == "" {
		cfg.Cron = "0 8 * * *" // 不填的话，默认每天8点开始签到
	}
	_, err = Cron.AddFunc(cfg.Cron, func() {
		RunClient(ct)
	})
	if err != nil {
		log.Fatalf("faild init cron err=%v", err)
	}
	Cron.Run()
}
