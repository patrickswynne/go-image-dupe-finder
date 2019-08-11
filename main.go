package go_image_dupe_finder

import (
	"fmt"
	"image"
	"math"
)

func main(){
	//start at root directory
	//gather all image paths
	//compare each image to other known images
	//create result directory
	//display duplicates, pointing to actual locations of two dupes in result directory
}
/*
Visit every pixel, break it down into its parts: R, G, B, & A (in go: image.At(x,y).RGBA())
Subtract the RGBA vals from their corresponding pixel vals in the reference image.
Square the differences, add them up.
Take the square root of the total sum.
	*/
func FastCompareImages(img1, img2 *image.RGBA) (int64, error) {
	if img1.Bounds() != img2.Bounds() {
		return 0, fmt.Errorf("image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
	}

	accumError := int64(0)

	for i := 0; i < len(img1.Pix); i++ {
		accumError += int64(sqDiffUInt8(img1.Pix[i], img2.Pix[i]))
	}

	return int64(math.Sqrt(float64(accumError))), nil
}

func sqDiffUInt8(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}