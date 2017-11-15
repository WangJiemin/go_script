package sendmail

import (
	"log"
	"strings"
	//"strconv"

	gomail "gopkg.in/gomail.v2"
)

var (
	mailFromUser string
	mailPwd      string
	mailHost     string
	//mailPort     string
	mailToUser string
	mailCcUser string
)

func SendMailInitialize(config map[string]string) {

	mailFromUser = config["mail_from_user"]
	mailPwd = config["mail_pwd"]
	mailHost = config["mail_host"]
	//mailPort = config["mail_port"]
	mailToUser = config["mail_to_user"]
	mailCcUser = config["mail_cc_user"]

}

func SendGoMail(config map[string]string, body, work string) {
	SendMailInitialize(config)
	mail := gomail.NewMessage()
	mail.SetAddressHeader("From", mailFromUser, "王杰民")      // 发件人
	mail.SetHeader("To", strings.Split(mailToUser, ",")...) // 收件人
	//mail.SetHeader("To", mail.FormatAddress(mailToUser, "收件人")) // 收件人
	mail.SetHeader("Cc", strings.Split(mailCcUser, ",")...) // 抄送
	//mail.SetHeader("Cc", mail.FormatAddress(mailCcUser, "抄送"))  // 抄送
	//mail.SetHeader("Bcc",mail.FormatAddress(bc_user,"暗送")) // 暗送
	mail.SetHeader("Subject", body) // 邮件主题
	//mail.SetBody("详情请看附件")          // 邮件正文
	mail.Attach(work) // 邮件附件

	send := gomail.NewPlainDialer(mailHost, 587, mailFromUser, mailPwd) // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := send.DialAndSend(mail); err != nil {
		log.Printf("SendGoMail failed to send\n", err)
		return
	} else {
		log.Print("Dond.SendGoMail sent successfully\n")
	}

}
