package utils

import (
	"context"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/shivajee98/aamishrit/internal/config"
)

var cloudinaryInstance *cloudinary.Cloudinary

func init() {

	cfg := config.LoadEnv()

	CloudinarySecretKey := cfg.ClerkSecretKey

	var err error

	cloudinaryInstance, err = cloudinary.NewFromURL(CloudinarySecretKey)

	CheckError("Error initialising Cloudinary", err)

}

func UploadImageFromStream(reader io.Reader) (string, error) {
	resp, err := cloudinaryInstance.Upload.Upload(context.Background(), reader, uploader.UploadParams{
		Folder: "exhibitor-images",
	})

	CheckError("Error uploading image", err)

	return resp.SecureURL, nil
}
