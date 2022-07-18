/**
* @Author: scjtqs
* @Date: 2022/7/18 11:55
* @Email: scjtqs@qq.com
 */
// Package adapter 签到服务详细配置
package adapter

import (
	"errors"
	"fmt"
	"github.com/scjtqs2/fqsign_go/config"
	"github.com/scjtqs2/fqsign_go/util"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

/**
https://pucloud.vip/ 的飞机场每日签到，多账户版，多余的账户，请于List中删除
注册即送100GB流量,充值一块即可成为会员获得2T流量,邀请返利更变态可无限白嫖,更有专属客户端加持！http://invcode.puyun.me/0Vi15Z
*/
// Pucloud 类
type Pucloud struct {
	User       *config.UserOption
	Qqpush     config.QqPush
	Domain     string
	ConfigName string
	Client     http.Client
}

// NewPucloud 初始化类
func NewPucloud(user *config.UserOption, qqpush config.QqPush) *Pucloud {
	gCurCookieJar, _ := cookiejar.New(nil) // 持久化 cookie
	if user.Domain == "" {
		user.Domain = "https://pucloud.vip"
	}
	return &Pucloud{
		User:       user,
		Qqpush:     qqpush,
		Domain:     user.Domain,
		Client:     http.Client{Timeout: time.Second * 6, Jar: gCurCookieJar},
		ConfigName: config.ConfigNamePucloud,
	}
}

// Login 登录接口
func (v *Pucloud) Login() error {
	posturl := fmt.Sprintf("%s/auth/login", v.Domain)
	postdata := url.Values{
		"email":  {v.User.UserName},
		"passwd": {v.User.UserPassword},
	}
	header := make(http.Header)
	header.Set("Users-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:32.0) Gecko/20100101 Firefox/32.0")
	header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	res, err := util.Post(v.Client, posturl, header, []byte(postdata.Encode()))
	if err != nil {
		return err
	}
	logincheck := gjson.ParseBytes(res)
	if logincheck.Get("ret").Int() != 1 {
		log.Errorf("登录失败 resp %s err=%v", string(res), err)
		errmsg := fmt.Sprintf("账号 %s 签到 %s 登录失败", v.User.UserName, v.ConfigName)
		v.MsgPush(errmsg)
		return errors.New(errmsg)
	}
	log.Infof("账号 %s %s 登录成功", v.User.UserName, v.ConfigName)
	return nil
}

// Checkin 签到接口
func (v *Pucloud) Checkin() error {
	posturl := fmt.Sprintf("%s/user/checkin", v.Domain)
	postdata := url.Values{}
	header := make(http.Header)
	header.Set("Users-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:32.0) Gecko/20100101 Firefox/32.0")
	header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	res, err := util.Post(v.Client, posturl, header, []byte(postdata.Encode()))
	if err != nil {
		return err
	}
	checkin := gjson.ParseBytes(res)
	if checkin.Get("ret").Int() != 1 {
		log.Errorf("签到失败 resp %s err=%v", string(res), err)
		errmsg := fmt.Sprintf("账号 %s  %s  签到失败 %s", v.User.UserName, v.ConfigName, checkin.Get("msg").String())
		v.MsgPush(errmsg)
		return errors.New(errmsg)
	}
	log.Infof("账号 %s %s 签到成功", v.User.UserName, v.ConfigName)
	msg := fmt.Sprintf("账号 %s %s 签到成功 %s ", v.User.UserName, v.ConfigName, checkin.Get("msg").String())
	v.MsgPush(msg)
	return nil
}

// Logout 登出接口
func (v *Pucloud) Logout() error {
	posturl := fmt.Sprintf("%s/user/logout")
	postdata := url.Values{}
	header := make(http.Header)
	header.Set("Users-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:32.0) Gecko/20100101 Firefox/32.0")
	header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	_, err := util.Post(v.Client, posturl, header, []byte(postdata.Encode()))
	if err != nil {
		return err
	}
	return nil
}

// MsgPush 消息推送处理接口
func (v *Pucloud) MsgPush(msg string) {
	// 优先推送 独立配置
	if v.User.Cqq != "" && v.User.QqToken != "" {
		res, err := util.Qqpush(msg, v.User.Cqq, v.User.QqToken)
		if err != nil {
			errmsg := fmt.Sprintf("账号 %s %s 推送消息失败 err=%v", v.User.UserName, v.ConfigName, err)
			log.Error(errmsg)
			return
		}
		log.Println(string(res))
		return
	}
	// 未配置独立推送，使用全局推送
	if v.Qqpush.Cqq != "" && v.Qqpush.QqToken != "" {
		res, err := util.Qqpush(msg, v.Qqpush.Cqq, v.Qqpush.QqToken)
		if err != nil {
			errmsg := fmt.Sprintf("账号 %s %s 推送消息失败 err=%v", v.User.UserName, v.ConfigName, err)
			log.Error(errmsg)
			return
		}
		log.Println(string(res))
	}
}
