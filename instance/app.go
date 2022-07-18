/**
* @Author: scjtqs
* @Date: 2022/7/18 11:46
* @Email: scjtqs@qq.com
 */
package instance

// AppInterface 签到接口定义
type AppInterface interface {
	Login() error       // 登入
	Checkin() error     // 签到
	Logout() error      // 登出
	MsgPush(msg string) // 推送消息到qq
}
