package email

type SMTP struct {
	Server   string
	Port     string
	User     string
	Password string
	UseTLS   bool
}
type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

type Message interface {
	SetSMTP(smtp SMTP)
	GetSMTP() SMTP
	SetMessage(m Message)
	GetMessage() Message
	Send() error
}
