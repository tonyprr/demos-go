package main

import "fmt"

// SMS Email

type INotificationFactory interface {
	SendNotification()
	AddDestinataries(des ...string)
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
	destinataries []string
}

func (smsN SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS to ", smsN)
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

func (smsN *SMSNotification) AddDestinataries(des ...string) {
	smsN.destinataries = append(smsN.destinataries, des...)
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct {
	destinataries []string
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func (emailN *EmailNotification) AddDestinataries(des ...string) {
	emailN.destinataries = append(emailN.destinataries, des...)
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	} else if notificationType == "Email" {
		return &EmailNotification{}, nil
	} else {
		return nil, fmt.Errorf("No notification Type")
	}
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	smsFactory.AddDestinataries("Juan", "Manuel")
	smsFactory.AddDestinataries("Antonio", "Ángel")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
