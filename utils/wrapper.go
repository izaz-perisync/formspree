package wrapper

import (
	"bytes"
	"text/template"

	"github.com/perisynctechnologies/formSpree/mail"
)

// ISetting is an interface that defines methods for wrapping mail templates.
type ISetting interface {
	WrapMailer(filename string, data any, recipients string, subject string) error
}

// MailTemplate represents the structure of a mail template, including its name and OTP (One-Time Password).
type MailTemplate struct {
	Name string // Name of the email user template
	Otp  int    // One-Time Password

	Password string
}

// Setting represents a configuration for mail template wrapping and mailing.
type Setting struct {
	path string           // The path to the directory containing mail templates.
	sm   *mailmon.Service // The mail service used for sending emails.
}

// New creates a new Setting instance with the provided path and mail service.
func New(path string, sm *mailmon.Service) ISetting {
	return &Setting{
		path: path,
		sm:   sm,
	}
}

// WrapMailer wraps a mail template with the provided data and sends it to the specified recipients.
//
// The filename parameter specifies the name of the mail template file.
// The data parameter represents the data to be used in the template.
// The recipients parameter specifies the email addresses of the recipients.
// The subject parameter specifies the subject of the email.
// If successful, WrapMailer sends the wrapped mail template to the recipients.
// Otherwise, it returns an error.

func (s *Setting) WrapMailer(filename string, data any, recipients, subject string) error {
	// Parse the specified mail template file.

	template, err := template.New(filename).Funcs(template.FuncMap{"increment": increment}).ParseFiles(s.path + "/" + filename)
	if err != nil {
		return err
	}

	// Execute the parsed template with the provided data.
	var buffer bytes.Buffer

	if err := template.Execute(&buffer, data); err != nil {
		return err
	}

	// Send the generated content as an email.
	err = s.sm.SendMail(subject, buffer.String(), []string{recipients})
	if err != nil {
		return err
	}

	return nil
}

func increment(i int) int {
	return i + 1
}
