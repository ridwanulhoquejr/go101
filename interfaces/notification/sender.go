package notification

import "fmt"

// in notificaiotn we nned Email from user/repository
// so we will declare an interface where we will define the method signature of GetEmai
type UserFetcher interface {
	GetEmail(id string) (string, error)
}

type EmailDispatcher interface {
	Send(to, subject, body string) error
}

// then we have to embeds both of this interface in a single struct
type EmailSender struct {
	users UserFetcher
	mail  EmailDispatcher
}

func NewEmailSender(uf UserFetcher, ed EmailDispatcher) *EmailSender {
	return &EmailSender{
		users: uf,
		mail:  ed,
	}
}

// send a welcome email
func (w *EmailSender) SendWelcome(userId string) error {

	email, err := w.users.GetEmail(userId)
	if err != nil {
		return fmt.Errorf("email not found for this userId")
	}

	sub := "Welcome to the team!"
	body := "Hello your account is read!"

	if err := w.mail.Send(email, sub, body); err != nil {
		return fmt.Errorf("Cannot send the email at the moment")
	}

	return nil
}
