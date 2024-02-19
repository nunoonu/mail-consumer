package repositories

import (
	"context"
	"crypto/tls"
	"github.com/nunoonu/mail-consumer/internal/core/ports"
	gomail "gopkg.in/mail.v2"
	"io"
	"log/slog"
)

type MailParams struct {
	Host     string
	Port     int
	From     string
	Password string
	To       string
	Subject  string
	Body     string
}

func NewMailParams() *MailParams {
	return &MailParams{
		Host:     "smtp.gmail.com",
		Port:     587,
		From:     "ph.chunnapiyar@gmail.com",
		Password: "qwol yvcr arcc mxfx",
		To:       "nu__panupong@hotmail.com",
		Subject:  "File uploading service",
		Body:     "This is an attachment.",
	}
}

type mailRepository struct {
	params *MailParams
}

func NewMailRepository(mp *MailParams) ports.MailRepository {
	return &mailRepository{params: mp}
}

type Mail struct {
	FileName string
	File     []byte
}

func (m mailRepository) Send(_ context.Context, fileName string, file []byte) error {
	p := m.params
	gm := gomail.NewMessage()

	gm.SetHeader("From", p.From)
	gm.SetHeader("To", p.To)
	gm.SetHeader("Subject", p.Subject)
	gm.SetBody("text/plain", p.Body)
	gm.Attach(fileName, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(file)
		return err
	}))

	d := gomail.NewDialer(p.Host, p.Port, p.From, p.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(gm); err != nil {
		slog.Error("Could not send email", slog.String("Err", err.Error()))
		return err
	}
	slog.Info("Email sent")
	return nil
}
