package uploader

import (
	"io"

	"github.com/shivajee98/aamishrit/pkg/utils"
)

type CloudinaryUploader struct{}

func (c *CloudinaryUploader) Upload(file io.Reader) (string, error) {
	return utils.UploadImage(file)
}
