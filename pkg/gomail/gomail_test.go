// Package now
// @author moqi
// On 2023/07/24 16:52:35
// see: https://github.com/go-gomail/gomail
package gomail_test

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"testing"
)

func TestMailTest(t *testing.T) {

	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		fmt.Println("Email sent!")
	}
}
