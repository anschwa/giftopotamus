package giftex

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"sort"
)

var (
	ErrEmailFailedDelivery = errors.New("Error sending mail")
)

type Mailer interface {
	Send(e Email) error
}

type Email struct {
	To, From   string
	Subject    string
	Text, HTML string
}

type TmplData struct {
	SubjectName, AssignedName string
}

type BulkTmplData struct {
	Entries []TmplData
}

type EmailService struct {
	sender, subject            string
	textTmpl, htmlTmpl         *template.Template
	bulkTextTmpl, bulkHTMLTmpl *template.Template

	mailer Mailer
}

func NewEmailService(sender, subject string, textTmpl, htmlTmpl, bulkTextTmpl, bulkHTMLTmpl *template.Template, mailer Mailer) *EmailService {
	return &EmailService{
		sender:       sender,
		subject:      subject,
		textTmpl:     textTmpl,
		htmlTmpl:     htmlTmpl,
		bulkTextTmpl: bulkTextTmpl,
		bulkHTMLTmpl: bulkHTMLTmpl,
		mailer:       mailer,
	}
}

func NewEmail(to, from, subject string, textTmpl, htmlTmpl *template.Template, data interface{}) (email Email, err error) {
	var textBuf bytes.Buffer
	if err := textTmpl.Execute(&textBuf, data); err != nil {
		return email, fmt.Errorf("Error rendering text: %w", err)
	}

	var htmlBuf bytes.Buffer
	if err := htmlTmpl.Execute(&htmlBuf, data); err != nil {
		return email, fmt.Errorf("Error rendering html: %w", err)
	}

	email = Email{
		To:      to,
		From:    from,
		Subject: subject,
		Text:    textBuf.String(),
		HTML:    htmlBuf.String(),
	}

	return email, nil
}

type FailedEmail struct {
	Email Email
	Err   error
}

// SendEmails will build and send emails using the provided Mailer and return a list of any failed deliveries
func (svc *EmailService) SendEmails(participants ParticipantMap, results Assignment) ([]FailedEmail, error) {
	emails := make([]Email, 0, len(results))

	// Find participants using the same email address so we can send
	// one email with all their assignments in it
	groupByEmail := make(map[string][]Participant)
	for pid := range results {
		p := participants[pid]
		groupByEmail[p.Email] = append(groupByEmail[p.Email], p)
	}

	for addr, entries := range groupByEmail {
		// Build regular emails
		if len(entries) == 1 {
			subject := entries[0]
			assigned := participants[results[subject.ID]]

			data := TmplData{
				SubjectName:  subject.Name,
				AssignedName: assigned.Name,
			}

			mail, err := NewEmail(subject.Email, svc.sender, svc.subject, svc.textTmpl, svc.htmlTmpl, data)
			if err != nil {
				return nil, fmt.Errorf("Error building email: %w", err)
			}

			emails = append(emails, mail)
			continue
		}

		// Build bulk emails
		var data BulkTmplData
		for _, p := range entries {
			data.Entries = append(data.Entries, TmplData{
				SubjectName:  p.Name,
				AssignedName: participants[results[p.ID]].Name,
			})
		}

		// Sort entries by name before building email
		sort.SliceStable(data.Entries, func(i, j int) bool {
			a, b := data.Entries[i], data.Entries[j]
			return a.SubjectName < b.SubjectName
		})

		mail, err := NewEmail(addr, svc.sender, svc.subject, svc.bulkTextTmpl, svc.bulkHTMLTmpl, data)
		if err != nil {
			return nil, fmt.Errorf("Error building email: %w", err)
		}

		emails = append(emails, mail)
	}

	// Send out all the emails!
	var failed []FailedEmail
	for _, mail := range emails {
		if err := svc.mailer.Send(mail); err != nil {
			failed = append(failed, FailedEmail{Email: mail, Err: err})
		}
	}

	if len(failed) > 0 {
		return failed, ErrEmailFailedDelivery
	}

	return nil, nil
}
