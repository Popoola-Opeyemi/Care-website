package utils

import (
	"fmt"
)

// ImageHandler ...
func ImageHandler(img []byte, env *Enviroment) (image *ImageData, err error) {

	image, err = SaveImageData(env, img)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
