package utils

import (
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

var cloudinaryInstance *cloudinary.Cloudinary

func init() {

	err := godotenv.Load()

	CheckError("Error Loading Env file", err)

	cloudinaryInstance, err = cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))

	CheckError("Error initialising Cloudinary", err)

}

func UploadImageFromStream(reader io.Reader) (string, error) {
	resp, err := cloudinaryInstance.Upload.Upload(context.Background(), reader, uploader.UploadParams{
		Folder: "exhibitor-images",
	})

	CheckError("Error uploading image", err)

	return resp.SecureURL, nil
}
