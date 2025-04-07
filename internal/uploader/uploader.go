package uploader

import (
	"io"

	"github.com/shivajee98/aamishrit/internal/config"
)

type FileUploader interface {
	Upload(file io.Reader) (string, error)
}

// uploader/cloudinary.go
func NewCloudinaryUploader(cfg *config.Config) *CloudinaryUploader {
	return &CloudinaryUploader{
		cloudinaryUrl: cfg.CLOUDINARY_URL,
	}
}
