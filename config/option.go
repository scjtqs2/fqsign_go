/**
* @Author: scjtqs
* @Date: 2022/7/18 11:23
* @Email: scjtqs@qq.com
 */
package config

type Option struct {
	Users  map[string][]*UserOption `yaml:"users"`
	QqPush QqPush                   `yaml:"qq_push"`
	Cron   string                   `yaml:"cron"` //  服务方式运行，定时签到的 crontab 定时参数配置
}

// QqPush 统一推送配置，不填不推送。
type QqPush struct {
	Cqq     string `yaml:"cqq"`
	QqToken string `yaml:"qq_token"`
}

type UserOption struct {
	UserName     string `yaml:"user_name"`     // 账号名
	UserPassword string `yaml:"user_password"` // 账号密码
	Cqq          string `yaml:"cqq"`           // 单独配置的推送配置，选填
	QqToken      string `yaml:"qq_token"`      // 单独的推送配置，选填
	Domain       string `yaml:"domain"`        // 设定自定义的配置域名，用于机场换地址后免编译更新覆盖地址。或者同样一套架构的新机场。
}
