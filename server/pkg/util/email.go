package util

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"server/pkg/setting"
)

func sendMail(mailTo []string, subject string, body string) error {
	alias := "PokeShare Official"

	m := gomail.NewMessage()
	m.SetHeader("From", alias+"<"+setting.EmailUser+">") //这种方式可以添加别名
	m.SetHeader("To", mailTo...)                            //发送给多个用户
	m.SetHeader("Subject", subject)                         //设置邮件主题
	m.SetBody("text/html", body)                            //设置邮件正文

	d := gomail.NewDialer(setting.EmailHost, setting.EmailPort, setting.EmailUser, setting.EmailPass)

	// cert, err := tls.LoadX509KeyPair(setting.HTTPS_CRT, setting.HTTPS_KEY)
	// if err != nil {
	// 	return err
	// }
	// d.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(m)
}

func SendNewPostMessage() error {
	mailTo := []string{
		"z.jiang@pokeshare.monster",
		"zeminjiang@mail.ru",
	}
	subject := "A new team has been updated. "
	body := "A new team has been updated just now. "
	return sendMail(mailTo, subject, body)
}

