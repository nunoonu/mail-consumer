package ports

import "context"

type MailRepository interface {
	Send(ctx context.Context, fileName string, file []byte) error
}
