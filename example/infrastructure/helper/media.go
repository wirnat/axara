package helper

import (
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
)

const width = 150

func GenerateThumbnail(filename string, img multipart.FileHeader) (*os.File, error) {
	file, err := img.Open()
	if err != nil {
		return nil, err
	}

	imgFile, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	m := resize.Thumbnail(width, width, imgFile, resize.NearestNeighbor)

	out, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	// write new image to file
	err = jpeg.Encode(out, m, nil)
	if err != nil {
		return nil, err
	}

	return out, nil
}
