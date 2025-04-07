package uploader

import (
	"context"
	"io"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type CloudinaryUploader struct {
	cloudinaryUrl string
}

func (c *CloudinaryUploader) Upload(file io.Reader) (string, error) {
	cld, err := cloudinary.NewFromURL(c.cloudinaryUrl)
	if err != nil {
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	log.Println("Upload successful. URL:", uploadResult.SecureURL)
	return uploadResult.SecureURL, nil
}
