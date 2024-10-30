package pkg

import (
	"fmt"
	"sync"
	"time"
)

// image processing handler function
func ImageProcesHandler(inputPath, resultPath string, imgProcessInfo ImageProcessObj) {
	var readFiles = make(chan string)
	var wg sync.WaitGroup

	fmt.Println("----Image Process Started----")
	start := time.Now()

	wg.Add(1)
	go func() {
		defer wg.Done()
		getDirectoryFiles(inputPath, readFiles)
	}()
	for f := range readFiles {
		wg.Add(1)
		go func(fName string) {
			defer wg.Done()
			processImage(fName, resultPath, imgProcessInfo)
		}(f)
	}

	wg.Wait()

	fmt.Println(time.Since(start))
	fmt.Println("----Image Process Ended----")
}
