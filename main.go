package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/image-processor/pkg"
)

func main() {

	inputPath := pkg.IMG_INPUTPATH
	resultPath := pkg.IMG_OUTPUTPATH

	args := os.Args
	// scale : upscale/downscale + factor
	// grayscale
	// compress : + level
	// process=scale:upscale:2|scale:downscale:2|grayscale|compress:80

	if len(args) < 2 {
		fmt.Println("Incorrect argument/s, correct(separated by | for multiple process) : process=scale:upscale:1/2/3(quality, 3 best):2|scale:downscale:2|grayscale|compress:80")
		return
	}

	var imgProcessInfo pkg.ImageProcessObj

	processArr := strings.Split(args[1], "|")
	for _, singleProcss := range processArr {
		processInfo := strings.Split(singleProcss, ":")
		if processInfo[0] == "" {
			fmt.Println("Incorrect argument/s, correct(separated by | for multiple process) : process=scale:upscale:1/2/3(quality, 3 best):2|scale:downscale:2|grayscale|compress:80")
			return
		}

		switch processInfo[0] {
		case "scale":
			if len(processInfo) < 4 {
				fmt.Println("Incorrect argument/s, correct(separated by | for multiple process) : process=scale:upscale:1/2/3(quality, 3 best):2")
				return
			}
			scalefactor, _ := strconv.Atoi(processInfo[3])
			imgProcessInfo.IsScale = true
			imgProcessInfo.ScaleType = processInfo[1]
			imgProcessInfo.ScaleSpeed = processInfo[2]
			imgProcessInfo.ScaleFactor = scalefactor
		case "grayscale":
			imgProcessInfo.IsGrayscale = true

		case "compress":
			if len(processInfo) < 2 {
				fmt.Println("Incorrect argument/s, correct(separated by | for multiple process) : process=compress:80")
				return
			}
			compressLevel, _ := strconv.Atoi(processInfo[1])
			imgProcessInfo.IsCompress = true
			imgProcessInfo.CompressLevel = compressLevel
		}
	}
	pkg.ImageProcesHandler(inputPath, resultPath, imgProcessInfo)
}
