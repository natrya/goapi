package mailer

import (
	"net/http"
	"os"
	"strconv"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type sendMail struct{}

type SendMailer interface {
	SendResetPassword(string, string, string, string) (*EmailResponse, error)
}

var (
	SendMail SendMailer = &sendMail{} //this is useful when we start testing
)

type EmailResponse struct {
	Status   int
	RespBody string
}

func (s *sendMail) SendResetPassword(ToUser string, FromAdmin string, Token string, AppEnv string) (*EmailResponse, error) {
	mailer := gomail.NewMessage()
	h := hermes.Hermes{
		Product: hermes.Product{
			Name:      "Natrya Forum",
			Link:      "https://natrya.co.id",
			Copyright: "Copyright © 2021 Natrya Kreasi Mandiri. All rights reserved.",
		},
	}
	var forgotUrl string
	if os.Getenv("APP_ENV") == "production" {
		forgotUrl = "https://natrya.co.id/resetpassword/" + Token //this is the url of the frontend app
	} else {
		forgotUrl = "http://127.0.0.1:8888/resetpassword/" + Token //this is the url of the local frontend app
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: ToUser,
			Intros: []string{
				"Welcome to Natrya! Good to have you here.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Click this link to reset your password",
					Button: hermes.Button{
						Color: "#FFFFFF",
						Text:  "Reset Password",
						Link:  forgotUrl,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}
	smtp, err := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	if err != nil {
		return nil, err
	}
	mailer.SetHeader("From", FromAdmin)
	mailer.SetHeader("To", ToUser)
	mailer.SetHeader("Subject", "Reset Password")
	mailer.SetBody("text/html", emailBody)
	dialer := gomail.NewDialer(
		os.Getenv("CONFIG_SMTP_HOST"),
		smtp,
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
	)
	go func() {
		errd := dialer.DialAndSend(mailer)
		if errd != nil {
			return
		}
	}()

	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success, Please click on the link provided in your email",
	}, nil
}
