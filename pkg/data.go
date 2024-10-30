package pkg

const IMG_INPUTPATH string = "./img"
const IMG_OUTPUTPATH string = "./processed_img/"

type ImageProcessObj struct {
	IsScale       bool
	IsGrayscale   bool
	IsCompress    bool
	CompressLevel int    // 1-100 to compress image
	ScaleFactor   int    // by how much to scale image
	ScaleSpeed    string // 1,2,3 (1 fastest but lowest quality, 3 slowest but best quality)
	ScaleType     string // upscale or downscale
}
