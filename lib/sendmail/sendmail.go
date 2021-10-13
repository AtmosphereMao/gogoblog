package sendmail

import (
	"encoding/base64"
	"fmt"
	"myblog/app/models/password_resets"
	"myblog/core"
	"myblog/lib/config"
	"myblog/lib/helper"
	"net/mail"
	"net/smtp"
)

func Send(pr password_resets.PasswordResets)(error){
	host := helper.ToString(config.Env("MAIL_HOST"))
	email := helper.ToString(config.Env("MAIL_USERNAME"))

	from := mail.Address{Name: helper.ToString(config.Env("APP_NAME")), Address: email}
	to := mail.Address{Name: "接收人", Address: pr.Email}

	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte("重置密码")))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"
	body := "点击链接修改您的账号的密码" + core.Name2URL("auth.forget.reset")
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))

	auth := smtp.PlainAuth(
		"",
		email,
		helper.ToString(config.Env("MAIL_PASSWORD")),
		host,
	)

	err := smtp.SendMail(
		host+":"+helper.ToString(config.Env("MAIL_PORT")),
		auth,
		email,
		[]string{to.Address},
		[]byte(message),
	)
	return err
}