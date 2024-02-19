package domain

type SendFileRequest struct {
	File     []byte
	FileName string
}

type SendFileResponse struct {
}
