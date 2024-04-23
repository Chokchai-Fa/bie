package email

import (
	"fastmath/models"
	"fastmath/utils/logs"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type EmailServer struct {
	From       string
	SMTPServer string
	SMTPPort   int
	Email      string
	Password   string
}

func NewEmailServer(serverEmail string, serverPassword string, smtpServer string, smtpPort int) (EmailServer, error) {
	return EmailServer{serverEmail, smtpServer, smtpPort, serverEmail, serverPassword}, nil
}

func (h *EmailServer) SendEmailActivateUser(toEmail string, activationToken uuid.UUID) error {

	message := gomail.NewMessage()
	message.SetHeader("From", h.From)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", "FastMath Account Activation")
	message.SetBody("text/html", fmt.Sprintf("Click <a href=\"http://localhost:8080/fastmath/user/activate?email=%s&token=%s\">here</a> to activate your account.",
		toEmail, activationToken))

	// Build the dialer and send the email.
	dialer := gomail.NewDialer(h.SMTPServer, h.SMTPPort, h.Email, h.Password)

	err := dialer.DialAndSend(message)
	if err != nil {
		errMsg := fmt.Sprintf("error on send email activate user| %s", toEmail)
		logs.Error(errMsg)
	}

	return err
}

func (h *EmailServer) SendEmailForgetPassword(toEmail string, activationToken uuid.UUID) error {

	message := gomail.NewMessage()
	message.SetHeader("From", h.From)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", "FastMath Account Forget Password")

	// url from mail will open FE which FE should get token from url and send it when redirect password
	message.SetBody("text/html", fmt.Sprintf("Click <a href=\"http://localhost:8080/fastmath/fe/email=%s&token=%s\">here</a> to re-create password.",
		toEmail, activationToken))

	// Build the dialer and send the email.
	dialer := gomail.NewDialer(h.SMTPServer, h.SMTPPort, h.Email, h.Password)

	err := dialer.DialAndSend(message)
	if err != nil {
		errMsg := fmt.Sprintf("error on send email forget password| %s", toEmail)
		logs.Error(errMsg)
	}

	return err
}

func (h *EmailServer) SendEmailPayment(toEmail string, paymentDetail models.PaymentDetail) error {
	message := gomail.NewMessage()
	message.SetHeader("From", h.From)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", "FastMath SubScription Detail")

	message.SetBody("text/html", fmt.Sprintf("Thank you For Subscription <br /> Payment Method: %s <br /> Amount: %s", paymentDetail.PaymentType, strconv.FormatInt(paymentDetail.Amount, 10)))

	// Build the dialer and send the email.
	dialer := gomail.NewDialer(h.SMTPServer, h.SMTPPort, h.Email, h.Password)

	err := dialer.DialAndSend(message)
	if err != nil {
		errMsg := fmt.Sprintf("error on send email payment sucessfully| %s", toEmail)
		logs.Error(errMsg)
	}

	return err
}
