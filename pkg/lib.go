package pkg

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
)

// returns : list of file names inside directory provided
func getDirectoryFiles(inputPath string, readFiles chan string) {
	// var filesList []string
	err := filepath.Walk(inputPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if it's a file (not a directory)
		if !info.IsDir() {
			readFiles <- path
			// filesList = append(filesList, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	// return filesList
	close(readFiles)
}

// processimage : processing function
func processImage(f, resultPath string, imgProcessInfo ImageProcessObj) {
	filePathArr := strings.Split(f, "\\")

	// original file
	input, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	// new resized file
	var outputFileName string
	outputFileNameArr := strings.Split(filePathArr[1], ".")

	var processFileName string
	if imgProcessInfo.IsScale {
		processFileName += "_" + imgProcessInfo.ScaleType
	}

	if imgProcessInfo.IsCompress {
		processFileName += "_compress"
	}

	if imgProcessInfo.IsGrayscale {
		processFileName += "_grayscale"
	}

	outputFileName = outputFileNameArr[0] + processFileName + "." + outputFileNameArr[1]
	output, _ := os.Create(resultPath + outputFileName)
	defer output.Close()

	img, _ := jpeg.Decode(input)
	var dst image.Image

	if imgProcessInfo.IsScale {
		dst = scaleImage(img, imgProcessInfo.ScaleSpeed, imgProcessInfo.ScaleType, imgProcessInfo.ScaleFactor)
	}

	if imgProcessInfo.IsGrayscale {
		if dst != nil {
			dst = applyGrayscale(dst)
		} else {
			img = applyGrayscale(img)
		}
	}

	var options *jpeg.Options
	if imgProcessInfo.IsCompress {
		if dst != nil {
			options = compressJPEG(dst, imgProcessInfo.CompressLevel)
		} else {
			options = compressJPEG(img, imgProcessInfo.CompressLevel)
		}
	}

	// Encode to `output`:
	if dst != nil {
		jpeg.Encode(output, dst, options)
	} else {
		jpeg.Encode(output, img, options)
	}
}

// applies grascale and returns image
func applyGrayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayscale := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, _ := originalColor.RGBA()
			gray := uint8((r*299 + g*587 + b*114) / 1000) // Luminance formula
			grayscale.Set(x, y, color.Gray{Y: gray})
		}
	}

	return grayscale
}

// compress image according to the input
func compressJPEG(img image.Image, quality int) *jpeg.Options {
	// output, err := os.Create(outputFile)
	// if err != nil {
	// 	return err
	// }
	// defer output.Close()

	options := &jpeg.Options{Quality: quality}
	return options
}

// scale image : upscale / downscale
func scaleImage(img image.Image, scaleSpeed, scaleType string, scaleFactor int) *image.RGBA {
	var dst *image.RGBA

	// create bounds
	if scaleType == "upscale" {
		dst = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X*scaleFactor, img.Bounds().Max.Y*scaleFactor))
	} else if scaleType == "downscale" {
		dst = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X/scaleFactor, img.Bounds().Max.Y/scaleFactor))
	} else {
		return nil
	}

	// scale image:
	switch scaleSpeed {
	case "1":
		draw.NearestNeighbor.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

	case "2":
		draw.BiLinear.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

	case "3":
		draw.CatmullRom.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

	default:
		return nil
	}

	return dst
}
