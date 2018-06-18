package ssim

import (
	cv "gocv.io/x/gocv"
	"math"
	"errors"
)

func MSE(reference, distorted *cv.Mat) (float64, error) {
	if reference.Cols() != distorted.Cols() {
		return 0.0, errors.New("images are not of same size (different width)")
	}
	if reference.Rows() != distorted.Rows() {
		return 0.0, errors.New("images are not of same size (different height)")
	}
	s1 := cv.NewMat()
	cv.AbsDiff(*reference, *distorted, &s1)
	s1.ConvertTo(&s1, cv.MatTypeCV32F)
	cv.Pow(s1, 2, &s1)
	sc := s1.Sum()
	defer s1.Close()
	
	sse := sc.Val1 + sc.Val2 + sc.Val3 + sc.Val4  // sum of squares
	
	if sse < 1e-10 {
		// TODO: is actually perfectly similar
		return 0, nil
	}
	mse := sse / float64(reference.Channels()*(reference.Cols() * reference.Rows()))
	
	return mse, nil
}

func PSNR(f1, f2 *cv.Mat) (float64, error) {
	mse, err := MSE(f1, f2)
	if err != nil {
		return 0.0, err
	}
	return 10.0 * math.Log10(math.Pow(255, 2)/mse), nil
}


// https://docs.opencv.org/master/d5/dc4/tutorial_video_input_psnr_ssim.html
