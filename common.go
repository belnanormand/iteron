package iteron

import (
	"image"
	"os"

	_ "image/png" //We need this to decode PNG's

	"github.com/faiface/pixel"
)

func loadPicture(path string) (pixel.Picture, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil

}

//Size General size struct
type Size struct {
	Width  float64
	Height float64
}

//Position Position of an object
type Position struct {
	X float64
	Y float64
}
