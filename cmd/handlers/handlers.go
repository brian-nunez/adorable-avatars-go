package handlers

import (
	"image/color"
	"image/png"
	"math/rand"

	"github.com/brian-nunez/adorable-avatars-go/cmd/avatars"
	"github.com/gin-gonic/gin"
)

func ListGETHandler(c *gin.Context) {
	avaliableFaceOptions, err := avatars.GetFaceOptions()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not load face options",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": avaliableFaceOptions,
	})
}

func RandomGETHandler(c *gin.Context) {
	faceOptions, err := avatars.GetFaceOptions()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not load face options",
		})
		return
	}

	eyesNum := rand.Intn(len(faceOptions.Eyes))
	noseNum := rand.Intn(len(faceOptions.Noses))
	mouthNum := rand.Intn(len(faceOptions.Mouths))

	eyes := faceOptions.Eyes[eyesNum]
	nose := faceOptions.Noses[noseNum]
	mouth := faceOptions.Mouths[mouthNum]

	images, err := avatars.GetFaceImages(
		eyes,
		nose,
		mouth,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not load face images",
		})
		return
	}

	finalImage := avatars.CreateFaceImage(
		color.RGBA{
			uint8(rand.Intn(255)),
			uint8(rand.Intn(255)),
			uint8(rand.Intn(255)),
			255,
		},
		images,
	)

	c.Header("Content-Type", "image/png")
	c.Writer.WriteHeader(200)
	err = png.Encode(c.Writer, finalImage)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not encode image",
		})
		return
	}
	c.Done()
}
