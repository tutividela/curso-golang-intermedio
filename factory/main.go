package main

import "fmt"

//Notificaciones por SMS , o Email

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {

}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification vis SMS")
}

type SMSNotificationSender struct {

}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}