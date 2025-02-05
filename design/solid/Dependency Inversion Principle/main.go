package main

import "fmt"

// 1. Define an abstraction (interface) for notification
type Notification interface {
	Send(message string) // Define a method for sending notifications
}

// 2. Low-level module: Email Notification implements the Notification interface
type EmailNotification struct{}

func (e EmailNotification) Send(message string) {
	fmt.Println("Sending Email: ", message)
}

// 3. Low-level module: SMS Notification implements the Notification interface
type SMSNotification struct{}

func (s SMSNotification) Send(message string) {
	fmt.Println("Sending SMS: ", message)
}

// 4. New low-level module: Push Notification implements the Notification interface
type PushNotification struct{}

func (p PushNotification) Send(message string) {
	fmt.Println("Sending Push Notification: ", message)
}

// 5. High-level module: NotificationService depends on the abstraction (Notification interface)
type NotificationService struct {
	notification Notification // Notification is an abstraction
}

func (n NotificationService) SendNotification(message string) {
	n.notification.Send(message) // Uses the abstraction to send the notification
}

func main() {
	// 6. We can now use any notification type that implements the Notification interface

	// Using Email Notification
	emailNotificationService := NotificationService{
		notification: EmailNotification{}, // Injecting EmailNotification
	}
	emailNotificationService.SendNotification("Hello, this is an email!")

	// Using SMS Notification
	smsNotificationService := NotificationService{
		notification: SMSNotification{}, // Injecting SMSNotification
	}
	smsNotificationService.SendNotification("Hello, this is an SMS!")

	// Using Push Notification
	pushNotificationService := NotificationService{
		notification: PushNotification{}, // Injecting PushNotification
	}
	pushNotificationService.SendNotification("Hello, this is a push notification!")
}
