package files

import (
	"image"
	"os"
	"regexp"
)

func ReadDir(path string) ([]string, error) {
	files, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var fileNames []string

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}

func ReadDirWithoutExtensions(path string) ([]string, error) {
	files, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var fileNames []string

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		rawName := file.Name()

		re := regexp.MustCompile(`\.\w+$`)
		name := re.ReplaceAllString(rawName, "")

		if name == "" {
			continue
		}

		fileNames = append(fileNames, name)
	}

	return fileNames, nil
}

func ReadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return image, nil
}
