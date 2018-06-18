# GoCV Mean Structural Similarity

SSIM implementation in Golang, using GoCV.

Largely based off https://docs.opencv.org/master/d5/dc4/tutorial_video_input_psnr_ssim.html and the original paper, "Image Quality Assessment: From Error Visibility to Structural Similarity"

Not even in ALPHA stage.

## Example
```go
package main

import (
	"fmt"
	s "github.com/nicovillanueva/gocv-ssim"
	cv "gocv.io/x/gocv"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: compare file1 file2")
	}
	i1 := cv.IMRead(os.Args[1], cv.IMReadUnchanged)
	i2 := cv.IMRead(os.Args[2], cv.IMReadUnchanged)

	p, _ := s.PSNR(&i1, &i2)
	m := s.MSSIM(&i1, &i2)
	mse, _ := s.MSE(&i1, &i2)

	fmt.Printf("MSE : %f\n", mse)
	fmt.Printf("PSNR: %f\n", p)
	fmt.Printf("SSIM: %f\n", m)
}

```
