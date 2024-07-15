package models

type ProcessingRequest struct {
	Resize    string  `form:"resize"`
	Crop      string  `form:"crop"`
	Rotate    bool    `form:"rotate"`
	Blur      float64 `form:"blur"`
	Grayscale bool    `form:"grayscale"`
	Sharpen   bool    `form:"sharpen"`
}
