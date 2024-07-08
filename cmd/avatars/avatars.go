package avatars

import (
	"image"
	"image/color"
	"image/draw"
	"strconv"

	"github.com/brian-nunez/adorable-avatars-go/cmd/files"
)

type AvaliableFaceOptions struct {
	Eyes   []string `json:"eyes"`
	Noses  []string `json:"noses"`
	Mouths []string `json:"mouths"`
}

var avaliableFaceOptions AvaliableFaceOptions

func CreateFaceOptions() (AvaliableFaceOptions, error) {
	eyes, err := files.ReadDirWithoutExtensions("assets/eyes")
	if err != nil {
		return AvaliableFaceOptions{}, err
	}

	noses, err := files.ReadDirWithoutExtensions("assets/nose")
	if err != nil {
		return AvaliableFaceOptions{}, err
	}

	mouths, err := files.ReadDirWithoutExtensions("assets/mouth")
	if err != nil {
		return AvaliableFaceOptions{}, err
	}

	faces := AvaliableFaceOptions{
		Eyes:   eyes,
		Noses:  noses,
		Mouths: mouths,
	}

	return faces, nil
}

func LoadFaceOptions() (AvaliableFaceOptions, error) {
	faces, err := CreateFaceOptions()
	if err != nil {
		return AvaliableFaceOptions{}, err
	}
	return faces, nil
}

func GetFaceOptions() (AvaliableFaceOptions, error) {
	if avaliableFaceOptions.Eyes == nil {
		avaliableFaceOptionsInitial, err := LoadFaceOptions()
		if err != nil {
			return AvaliableFaceOptions{}, err
		}

		avaliableFaceOptions = avaliableFaceOptionsInitial
	}

	return avaliableFaceOptions, nil
}

func GetFaceOptionsWithoutCache() (AvaliableFaceOptions, error) {
	freshAvaliableFaceOptions, err := LoadFaceOptions()
	if err != nil {
		return AvaliableFaceOptions{}, err
	}

	avaliableFaceOptions = freshAvaliableFaceOptions

	return avaliableFaceOptions, nil
}

type AvatarFace struct {
	Eyes  *image.Image
	Nose  *image.Image
	Mouth *image.Image
}

func GetFaceImages(eyes string, nose string, mouth string) (AvatarFace, error) {
	eyeImage, err := files.ReadImage("assets/eyes/" + eyes + ".png")
	if err != nil {
		return AvatarFace{}, err
	}

	noseImage, err := files.ReadImage("assets/nose/" + nose + ".png")
	if err != nil {
		return AvatarFace{}, err
	}

	mouthImage, err := files.ReadImage("assets/mouth/" + mouth + ".png")
	if err != nil {
		return AvatarFace{}, err
	}

	return AvatarFace{
		Eyes:  &eyeImage,
		Nose:  &noseImage,
		Mouth: &mouthImage,
	}, nil
}

func CreateFaceImage(backgroundColor color.RGBA, face AvatarFace) *image.RGBA {
	bgImage := image.NewRGBA(image.Rect(0, 0, 400, 400))
	draw.Draw(bgImage, bgImage.Bounds(), &image.Uniform{backgroundColor}, image.Point{0, 400}, draw.Src)

	point := image.Point{0, 0}

	draw.Draw(bgImage, bgImage.Bounds(), *face.Eyes, point, draw.Over)
	draw.Draw(bgImage, bgImage.Bounds(), *face.Nose, point, draw.Over)
	draw.Draw(bgImage, bgImage.Bounds(), *face.Mouth, point, draw.Over)

	return bgImage
}

func HexToRGB(hex string) color.RGBA {
	values, _ := strconv.ParseUint(string(hex), 16, 32)
	rgbColor := color.RGBA{
		R: uint8(values >> 16),
		G: uint8(values >> 8),
		B: uint8(values),
		A: 255,
	}

	return rgbColor
}
