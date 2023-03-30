package controller

import (
	"GoTools/model"
	"bytes"
	"fmt"
	"text/template"

	"github.com/jasonlvhit/gocron"

	gm "gopkg.in/gomail.v2"
)

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	if err := t.Execute(&buff, data); err != nil {
		return "", err
	}

	return buff.String(), nil
}

func sendEmail(user model.User) {
	mail := gm.NewMessage()

	template := "bin/template/mail.html"

	result, _ := parseTemplate(template, user)

	mail.SetHeader("From", "Kx5bPjry3gREqQiKkVJrM27f@gmail.com")
	mail.SetHeader("To", user.Email)
	mail.SetHeader("Subject", "Notifications")
	mail.SetBody("text/html", result)

	sender := gm.NewDialer("smtp.gmail.com", 25, "Kx5bPjry3gREqQiKkVJrM27f@gmail.com", "tvzyqdonjztrwsod")

	if err := sender.DialAndSend(mail); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent to: ", user.Email)
	}

}

var scheduler = gocron.NewScheduler()

func startScheduler(user model.User) {
	scheduler.Every(30).Second().Do(func() {
		sendEmail(user)
	})
	scheduler.Start()
}

func stopScheduler() {
	scheduler.Clear()
}
