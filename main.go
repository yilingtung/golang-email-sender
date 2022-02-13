package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
)

func sendEmail() {
	// Sender data.
	from := "sender@email.com"
	password := "your_email_password"

	// Receiver email address.
	to := []string{
		"receiver@email.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Title         string
		LogoImgUrl    string
		BannerImgUrl  string
		ResetLink     string
	}{
		Title:        "Golang Email Sender",
		LogoImgUrl:   "https://fakeimg.pl/112x36/fff/000/?text=logo",
		BannerImgUrl: "https://fakeimg.pl/1280x900/eee/000/?text=Banner",
		ResetLink:    "https://www.google.com.tw",
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}

func devTemp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("template.html")
		t.Execute(w, map[string] string {
			"Title":        "Golang Email Sender",
			"LogoImgUrl":   "https://fakeimg.pl/112x36/fff/000/?text=logo",
			"BannerImgUrl": "https://fakeimg.pl/1280x900/eee/000/?text=Banner",
			"ResetLink":    "https://www.google.com.tw",
		})
	})

	http.ListenAndServe(":8000", nil)
}

func main() {
	// Define a bool flag
	devArgPtr := flag.Bool("dev", false, "Run HTML locally.")
	// Parse command line 
	// into the defined flags
	flag.Parse()

	if *devArgPtr {
		// Launch dev server for email template.
		devTemp()
	} else {
		// Sending email.
		sendEmail()
	}
}
