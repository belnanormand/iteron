package iteron

import (
	"image"
	"math"
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

//DegreesToRadians Converts degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return (degrees / 360) * 2 * math.Pi
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
