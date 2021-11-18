package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/anschwa/giftopotamus/giftex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

func main() {
	// Read CSV and generate gift exchange results
	db, err := giftex.ReadCSVFromFile("example.csv")
	if err != nil {
		panic(err)
	}

	opts := &giftex.GiftExchangeOptions{MaxPrevious: 2}
	ge, err := giftex.NewGiftExchange(db.Participants, opts)
	if err != nil {
		panic(err)
	}

	// Print results and ask for confirmation before sending mail
	fmt.Println(ge)
	if !confirm() {
		fmt.Println("Aborting")
		os.Exit(1)
	}

	// Save results as CSV
	f, err := os.Create("results.csv")
	if err != nil {
		panic(err)
	}

	if err := db.WriteCSV(f, ge.Assignment); err != nil {
		panic(err)
	}

	// Send emails
	mailer := &fakeMailer{}
	svc := giftex.NewEmailService(sender, subject, textTmpl, htmlTmpl, textBulkTmpl, htmlBulkTmpl, mailer)
	failed, err := svc.SendEmails(db.Participants, ge.Assignment)

	if err != nil {
		fmt.Println("Error! Some emails failed to send")
		for fail := range failed {
			fmt.Println(fail)
		}

		fmt.Println(err)
		os.Exit(1)
	}
}

const (
	sender  = "Giftopotamus <hello@example.com>"
	subject = "Gift Exchange"

	textTemplate = `Welcome to the gift exchange!

You have {{.AssignedName}}
`
	htmlTemplate = `Welcome to the gift exchange!<br/><br/>
You have {{.AssignedName}}
`

	textBulkTemplate = `Welcome to the gift exchange!
{{range .Entries}}
{{.SubjectName}} has {{.AssignedName}}
{{- end -}}
`
	htmlBulkTemplate = `Welcome to the gift exchange!<br/><br/>
{{range .Entries}}
{{.SubjectName}} has {{.AssignedName}}<br/>
{{- end -}}
`
)

var (
	textTmpl     = template.Must(template.New("text").Parse(textTemplate))
	textBulkTmpl = template.Must(template.New("textBulk").Parse(textBulkTemplate))

	htmlTmpl     = template.Must(template.New("html").Parse(htmlTemplate))
	htmlBulkTmpl = template.Must(template.New("htmlBulk").Parse(htmlBulkTemplate))
)

func confirm() bool {
	var yes string
	fmt.Print("Continue? (Yes/No) ")
	fmt.Scanln(&yes)

	return strings.ToLower(yes) == "yes"
}

type fakeMailer struct{}

func (m *fakeMailer) Send(e giftex.Email) error {
	fmt.Println("To:", e.To)
	fmt.Println("Text:", e.Text)
	fmt.Println("HTML:", e.HTML)
	fmt.Println()

	return nil
}

type sesMailer struct {
	svc *ses.SES
}

func (m *sesMailer) Send(e giftex.Email) error {
	const charSet = "UTF-8"

	input := &ses.SendEmailInput{
		Source: aws.String(e.From),
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(e.To),
			},
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(e.Subject),
			},
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(e.HTML),
				},
				Text: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(e.Text),
				},
			},
		},
	}

	_, err := m.svc.SendEmail(input)
	return err
}
