# fqsign

自用免费梯子的签到脚本

qq推送 请到 https://wx.scjtqs.com/qq/user 里面获取

## 配置文件说明

```yaml
# 各网站的签到配置，每个网站都支持多账号配置（yaml数组方式）。也支持独立推送配置。不配置留空就行了
users:
  "vpork":
    - "user_name": "账号1"
      "user_password": "密码1"
    - "user_name": "账号2"
      "user_password": "密码2"
      "cqq": "账号2 的专属推送qq，不配置的话，走全局配置。"
      "qq_token": "账号2 的专属推送qq 对应的token"
      "domain": "https://forever.pork16.com"  # 设定自定义的配置域名，用于机场换地址后免编译更新覆盖地址。或者同样一套架构的新机场。
  "vpork_vip": [ ]
  "cdpc": [ ]
  "liuns": [ ]
  "pucloud": [ ]

# 全局推送配置，如果上面配置了定向推送配置，那么配置了定向配置的推送，走定向配置，不走全局。
qq_push:
  "cqq": "全局推送qq"
  "qq_token": "全局推送qq对应的token"

# 签到执行的定时配置。server方式启动的时候才有效。不配置（留空）的话，默认为 每天早上8点。
"cron": "0 8 * * *"
```

```shell
# -c的默认参数是bash当前目录下的config.yaml文件。可以自行注入绝对路径
./fqsign -c ./config.yaml
```

## 手动执行脚本编译

```shell
go build -o fqsign cmd/client/*
```

这种入口编译的程序，需要自行维护执行时间。程序在执行完成一次签到后，将自行退出。

## 手动编译服务方式程序

```shell
go build -o fqsign cmd/server/*
```

这种入口编译的程序，会持续运行，会根据配置的执行频率来触发执行。适合docker中运行。

### docker方式运行

```shell
docker run -d --name fqsign -v `pwd`/config.yaml:/app/config.yaml --restart=always scjtqs/fqsign
```

### docker-compose方式

放入配置文件 `docker-compose.yaml`：

```yaml
version: 3
services:
  fqsign:
    image: scjtqs/fqsign:latest
    restart: always
    volumes:
      - ./config.yaml:/app/config.yaml     #传入配置文件
```

```shell
# 启动容器
docker-compose up -d
# 停止容器
docker-compose down
# 查看日志
docker-compose logs -f
```