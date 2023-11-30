package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

type Email struct {
	Name    string
	Email   string
	Subject string
	Message string
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unkown from server")
		}
	}
	return nil, nil
}

func email(c *gin.Context) {
	email := Email{}
	if err := c.ShouldBind(&email); err != nil {
		c.String(http.StatusAccepted, "input error "+err.Error())
		return
	}

	pass, ok := os.LookupEnv("emailpass")
	if !ok {
		c.String(http.StatusAccepted, "smtp credentails invalid")
		return
	}
	auth := LoginAuth("mkasun@gmail.com", pass)
	to := []string{"mkasun@nusak.ca"}
	msg := []byte("To: mkasun@nusak.ca\r\n" +
		fmt.Sprintf("Subject:  %s\r\n", email.Subject) +
		"\r\n" +
		fmt.Sprintf("email from %s<%s>\r\n%s", email.Name, email.Email, email.Message))
	if err := smtp.SendMail("smtp.gmail.com:587", auth, email.Email, to, msg); err != nil {
		c.String(http.StatusAccepted, "sendmail fail "+err.Error())
		return
	}
	c.String(http.StatusOK, "email sent")
}
