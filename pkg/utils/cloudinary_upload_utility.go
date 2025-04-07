package utils

import (
	"context"
	"io"
	"sync"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/shivajee98/aamishrit/internal/config"
)

var (
	cloudinaryInstance *cloudinary.Cloudinary
	initOnce           sync.Once
	initErr            error
)

func getCloudinaryInstance() (*cloudinary.Cloudinary, error) {
	initOnce.Do(func() {
		cfg := config.LoadEnv()
		cloudinaryInstance, initErr = cloudinary.NewFromURL(cfg.CLOUDINARY_URL)
	})
	return cloudinaryInstance, initErr
}

// 🔥 THIS is what you're calling from outside
func UploadImage(reader io.Reader) (string, error) {
	cld, err := getCloudinaryInstance()
	if err != nil {
		return "", err
	}

	resp, err := cld.Upload.Upload(context.Background(), reader, uploader.UploadParams{
		Folder: "exhibitor-images",
	})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
