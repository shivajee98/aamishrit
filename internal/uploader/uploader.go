package uploader

import "io"

type FileUploader interface {
	Upload(file io.Reader) (string, error)
}
