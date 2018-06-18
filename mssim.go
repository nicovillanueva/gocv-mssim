package ssim

import (
	cv "gocv.io/x/gocv"
	"image"
)

var (
	C1Sc cv.Scalar
	C2Sc cv.Scalar
)

// MSSIM constants
const (
	C1 = 6.5025
	C2 = 58.5225
)

// https://docs.opencv.org/master/d5/dc4/tutorial_video_input_psnr_ssim.html
func MSSIM(f1, f2 *cv.Mat) cv.Scalar {
	p := image.Point{11, 11}

	F1 := cv.NewMat()
	defer F1.Close()
	F2 := cv.NewMat()
	defer F2.Close()
	F1_F2 := cv.NewMat()
	defer F1_F2.Close()

	f1.ConvertTo(&F1, cv.MatTypeCV32F)
	f2.ConvertTo(&F2, cv.MatTypeCV32F)

	cv.Pow(F1, 2, &F1)
	cv.Pow(F2, 2, &F2)
	cv.Multiply(F1, F2, &F1_F2)

	// ---

	mu1 := cv.NewMat()
	defer mu1.Close()
	mu2 := cv.NewMat()
	defer mu2.Close()

	cv.GaussianBlur(F1, &mu1, p, 1.5, 0.0, cv.BorderDefault)
	cv.GaussianBlur(F2, &mu2, p, 1.5, 0.0, cv.BorderDefault)

	cv.Pow(mu1, 2, &mu1)
	cv.Pow(mu2, 2, &mu2)
	mu1_mu2 := cv.NewMat()
	defer mu1_mu2.Close()
	cv.Multiply(mu1, mu2, &mu1_mu2)

	sig1 := cv.NewMat()
	defer sig1.Close()
	sig2 := cv.NewMat()
	defer sig2.Close()
	sig1_2 := cv.NewMat()
	defer sig1_2.Close()

	cv.GaussianBlur(F1, &sig1, p, 1.5, 0.0, cv.BorderDefault)
	cv.Subtract(sig1, mu1, &sig1)
	cv.GaussianBlur(F2, &sig2, p, 1.5, 0.0, cv.BorderDefault)
	cv.Subtract(sig2, mu2, &sig2)
	cv.GaussianBlur(F1_F2, &sig1_2, p, 1.5, 0.0, cv.BorderDefault)
	cv.Subtract(sig1_2, mu1_mu2, &sig1_2)

	t1 := cv.NewMat()
	defer t1.Close()
	t2 := cv.NewMat()
	defer t2.Close()
	t3 := cv.NewMat()
	defer t3.Close()
	mat2 := cv.NewMatFromScalar(cv.NewScalar(2, 2, 2, 2), cv.MatTypeCV64F)
	defer mat2.Close()
	C1Mat := cv.NewMatFromScalar(C1Sc, cv.MatTypeCV64F)
	defer C1Mat.Close()
	C2Mat := cv.NewMatFromScalar(C2Sc, cv.MatTypeCV64F)
	defer C2Mat.Close()

	cv.Multiply(mu1_mu2, mat2, &t1)
	cv.Multiply(sig1_2, mat2, &t2)
	cv.Add(t1, C1Mat, &t1)
	cv.Add(t2, C2Mat, &t2)
	cv.Multiply(t1, t2, &t3)

	cv.Add(mu1_mu2, mu2, &mu1_mu2)
	cv.Add(mu1_mu2, C1Mat, &mu1_mu2)

	cv.Add(sig1, sig2, &sig1)
	cv.Add(sig1, C2Mat, &sig1)

	cv.Multiply(sig1, sig2, &sig1)

	ssim_map := cv.NewMat()
	defer ssim_map.Close()
	cv.Divide(t3, sig1, &ssim_map)

	return ssim_map.Mean()
}
