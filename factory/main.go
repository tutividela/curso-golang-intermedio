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

type EmailNotification struct {

}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notificaction via Email")
}

type EmailNotificationSender struct {

}

func (EmailNotificationSender) GetSenderMethod()  string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string{
	return "yes"
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func getNotificationFactory(notifcacionType string)( INotificationFactory,error) {
	if notifcacionType == "SMS"{
		return &SMSNotification{},nil
	}

	if notifcacionType == "Email"{
		return &EmailNotification{},nil
	}
	return nil,fmt.Errorf("No Notification Type")
}

func SendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory){
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main(){
	smsFactory , _ := getNotificationFactory("SMS")
	emailFactory , _ := getNotificationFactory("Email")

	SendNotification(smsFactory)
	SendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}