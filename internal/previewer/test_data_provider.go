package previewer

import (
	"image"
	"log"
	"os"
)

const (
	sourceImageExampleName  = "_gopher_original_1024x504.jpg"
	resizedImageExampleName = "_gopher_resized_800x300.jpg"
	//nolint:lll
	previewImageURL = "raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/_gopher_original_1024x504.jpg"
)

func getExampleImage(fileName string) image.Image {
	filePath := "../../image_examples/" + fileName

	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
		os.Exit(1)
	}
	defer imgFile.Close()

	img, _, _ := image.Decode(imgFile)

	return img
}
